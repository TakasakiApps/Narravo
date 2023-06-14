package utils

import (
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ohanakogo/exceptiongo"
	"time"
)

type JWTClaims[T any] struct {
	obj T
	jwt.RegisteredClaims
}

// JWTSign generates a JWT signed with the given key that expires after the specified number of seconds.
// The payload of the token is obj, which can be of any type.
func JWTSign[T any](obj T, key string, expireSeconds int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTClaims[T]{
		obj: obj,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expireSeconds))),
		},
	})
	signedString, err := token.SignedString([]byte(key))
	exceptiongo.QuickThrow[types.JWTSignFailedException](err)
	return signedString
}

// mJWTParse parses and verifies the given JWT token string using the provided key.
// It returns a pointer to a JWTClaims object containing the parsed claims if the token is valid.
// If the token is invalid or parsing fails, it throws a JWTParseFailedException.
func mJWTParse[T any](tokenString string, key string) *JWTClaims[T] {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims[T]{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if jwtClaims, ok := token.Claims.(*JWTClaims[T]); ok && token.Valid {
		return jwtClaims
	}
	exceptiongo.QuickThrow[types.JWTParseFailedException](err)
	return nil
}

// JWTParseObj parses the given JWT token string using the provided key and returns the decoded payload as an object of type T.
func JWTParseObj[T any](tokenString string, key string) T {
	return mJWTParse[T](tokenString, key).obj
}

// JWTParse parses the given JWT token string using the provided key and invokes the specified callback function with
// two arguments - a boolean value indicating whether the token has expired and the decoded payload as an object of type T.
func JWTParse[T any](tokenString string, key string, callback func(isExpired bool, obj T)) {
	claims := mJWTParse[T](tokenString, key)
	isExpired := claims.ExpiresAt.Before(time.Now())
	callback(isExpired, claims.obj)
}
