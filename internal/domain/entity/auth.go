package entity

import "time"

type LoginParams struct {
	Username string `json:"username" binding:"required,min=5,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}
