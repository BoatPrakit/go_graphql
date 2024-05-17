package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/graphql-go/graphql"
)

type AuthResolver struct {
	auth *Storage
}

func NewResolver(authStorage *Storage) *AuthResolver {
	return &AuthResolver{
		auth: authStorage,
	}
}

func (r *AuthResolver) Login(params graphql.ResolveParams) (interface{}, error) {
	email := params.Args["email"].(string)
	password := params.Args["password"].(string)
	isSuccess, err := r.auth.Verify(email, password)
	if err != nil {
		return nil, err
	}

	if !isSuccess {
		return nil, errors.New("login failed")
	}

	claims := JwtToken{
		UserId: "1",
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	res := JwtTokenResponse{
		Token: accessToken,
	}

	return res, nil
}

func (r *AuthResolver) Register(params graphql.ResolveParams) (interface{}, error) {
	userInput := params.Args["user"].(map[string]interface{})

	input := RegisterInput{
		Name:     userInput["name"].(string),
		Email:    userInput["email"].(string),
		Password: userInput["password"].(string),
	}
	_, err := r.auth.RegisterUser(input)
	if err != nil {
		return nil, err
	}
	return true, nil
}
