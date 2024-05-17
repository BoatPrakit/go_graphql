package auth

import "github.com/golang-jwt/jwt/v5"

type JwtToken struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type JwtTokenResponse struct {
	Token string `json:"token"`
}

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
