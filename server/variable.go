package server

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type response struct {
	Status  string
	Message interface{}
	Code    int
}

var key []byte

type credential struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type responseParam struct {
	W      http.ResponseWriter
	Body   []byte
	Status int
}
type user struct {
	Id       int
	Username string
	Email    string
	// LastLogin datatypes.Date
	// CreatedAt datatypes.Date
	// UpdatedAt datatypes.Date
	// DeletedAt datatypes.Date
}
