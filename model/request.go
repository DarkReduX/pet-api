package model

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
