package data

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id   int
	Name string
}

type UseClaims struct {
	jwt.StandardClaims
	User User `json:"user"`
}
