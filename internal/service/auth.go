package service

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"petProject/internal/config"
	"petProject/internal/repository"
	"petProject/model"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	cfg *config.JWT

	userRep *repository.User
	jwtRep  repository.JwtKeeper
}

func NewAuth(cfg *config.JWT, userRep *repository.User, jwtRep repository.JwtKeeper) *Auth {
	return &Auth{cfg: cfg, userRep: userRep, jwtRep: jwtRep}
}

func (s *Auth) SignUp(
	ctx context.Context,
	user *model.User,
) (createdUser *model.User, tokens *model.JSONWebTokens, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}

	user.Password = string(passwordHash)

	createdUser, err = s.userRep.Create(ctx, user)
	if err != nil {
		return nil, nil, err
	}

	tokens, err = s.createTokens(ctx, createdUser.ID)
	if err != nil {
		return nil, nil, err
	}

	// don't return password
	createdUser.Password = ""

	return createdUser, tokens, err
}

func (s *Auth) SignIn(
	ctx context.Context,
	req *model.SignInRequest,
) (createdUser *model.User, tokens *model.JSONWebTokens, err error) {
	storedUser, err := s.userRep.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(req.Password))
	if err != nil {
		return nil, nil, err
	}

	tokens, err = s.createTokens(ctx, storedUser.ID)

	// don't return password
	storedUser.Password = ""

	return storedUser, tokens, err
}

func (s *Auth) createTokens(ctx context.Context, userUUID uuid.UUID) (*model.JSONWebTokens, error) {
	at, err := s.createToken(userUUID)
	if err != nil {
		return nil, err
	}

	rt, err := s.createToken(userUUID)
	if err != nil {
		return nil, err
	}

	rtHash := fmt.Sprintf("%x", sha256.Sum256([]byte(rt)))

	err = s.jwtRep.CreateToken(ctx, userUUID, rtHash)

	return &model.JSONWebTokens{AccessToken: at, RefreshToken: rt}, err
}

func (s *Auth) createToken(userUUID uuid.UUID) (string, error) {
	at, err := jwt.NewWithClaims(jwt.SigningMethodHS512, model.JWTClaims{
		GUID:     uuid.New(),
		UserUUID: userUUID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(s.cfg.ExpireMinutesAT)).Unix(),
		},
	}).SignedString([]byte(s.cfg.SigningKeyAT))

	return at, err
}

func (s *Auth) RefreshToken(ctx context.Context, userUUID uuid.UUID, rt string) (*model.JSONWebTokens, error) {
	rtHashBin := sha256.Sum256([]byte(rt))
	rtHashStr := fmt.Sprintf("%x", rtHashBin)

	storedHash, err := s.jwtRep.GetTokenHashByUserID(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	if rtHashStr != storedHash {
		return nil, errors.New("incorrect refresh token")
	}

	at, err := s.createToken(userUUID)
	if err != nil {
		return nil, err
	}

	rt, err = s.createToken(userUUID)
	if err != nil {
		return nil, err
	}

	rtHash := fmt.Sprintf("%x", sha256.Sum256([]byte(rt)))
	err = s.jwtRep.CreateToken(ctx, userUUID, rtHash)

	return &model.JSONWebTokens{AccessToken: at, RefreshToken: rt}, err
}
