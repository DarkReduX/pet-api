package model

type AuthResponse struct {
	Tokens *JSONWebTokens `json:"tokens"`
	User   *User          `json:"user"`
}
