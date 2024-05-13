package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type JwtKeeper interface {
	CreateToken(ctx context.Context, userID uuid.UUID, rtHash string) error
	UpdateToken(ctx context.Context, userID uuid.UUID, rtHash string) error
	GetTokenHashByUserID(ctx context.Context, userID uuid.UUID) (rtHash string, err error)
}

type JwtPostgres struct {
	pool *pgxpool.Pool
}

func NewJwtPostgres(pool *pgxpool.Pool) *JwtPostgres {
	return &JwtPostgres{pool: pool}
}

func (r *JwtPostgres) CreateToken(ctx context.Context, userID uuid.UUID, rtHash string) error {
	_, err := r.pool.Exec(ctx, `INSERT INTO jwt (user_uuid, refresh_token_hash) VALUES ($1, $2) 
                                                ON CONFLICT (user_uuid) DO UPDATE SET refresh_token_hash = $2`,
		userID, rtHash)

	return err
}

func (r *JwtPostgres) UpdateToken(ctx context.Context, userID uuid.UUID, rtHash string) error {
	_, err := r.pool.Exec(ctx, "UPDATE jwt SET refresh_token_hash = $2 WHERE user_uuid = $1", userID, rtHash)

	return err
}

func (r *JwtPostgres) GetTokenHashByUserID(ctx context.Context, userID uuid.UUID) (rtHash string, err error) {
	err = r.pool.QueryRow(ctx, "SELECT refresh_token_hash FROM jwt WHERE user_uuid = $1", userID).Scan(&rtHash)

	return rtHash, err
}
