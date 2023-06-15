package utils

import (
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ohanakogo/exceptiongo"
	"time"
)

type JWTClaims[T any] struct {
	Obj T
	jwt.RegisteredClaims
}

// JWTSign generates a JWT with a given object, key and expiration time.
// The token is signed using HS256 algorithm.
// The token is returned as a string.
func JWTSign[T any](obj T, key string, expireSeconds int) string {
	// Create a new token with the given signing method and claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTClaims[T]{
		Obj: obj,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expireSeconds))),
		},
	})

	// Sign the token with the given key and return the signed token as a string.
	signedString, err := token.SignedString([]byte(key))
	exceptiongo.QuickThrow[types.JWTSignFailedException](err)
	return signedString
}

// mJWTParse parses a JWT token with the given string and key.
// If the token is valid, it returns its claims and a boolean indicating validity.
func mJWTParse[T any](tokenString string, key string) (*JWTClaims[T], bool) {
	// Parse the token with the given claims and function to validate the signature.
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims[T]{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	// Check if the JWT token and its claims are valid, and throw an exception if not.
	if token == nil || token.Claims == nil {
		exceptiongo.QuickThrow[types.JWTParseFailedException](err)
	}

	// If the claims can be extracted from the token, return them along with the validity of the token.
	if jwtClaims, ok := token.Claims.(*JWTClaims[T]); ok {
		return jwtClaims, token.Valid
	}

	// If there was an error while parsing the token or extracting the claims, throw an exception.
	exceptiongo.QuickThrow[types.JWTParseFailedException](err)
	return nil, false
}

// JWTParseObj parses a JWT token with the given string and key.
// It returns only the object contained in the token's claims.
func JWTParseObj[T any](tokenString string, key string) T {
	claims, _ := mJWTParse[T](tokenString, key)
	return claims.Obj
}

// JWTParse parses a JWT token with the given string and key.
// It takes as argument a callback function that is called with the validity of the token and its object,
// if the callback is not nil.
func JWTParse[T any](tokenString string, key string, callback func(isValid bool, obj T)) {
	claims, isValid := mJWTParse[T](tokenString, key)
	if callback != nil {
		callback(isValid, claims.Obj)
	}
}
