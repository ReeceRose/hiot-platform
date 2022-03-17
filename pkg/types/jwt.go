package types

import "github.com/golang-jwt/jwt"

//AuthClaims is a wrapper for StandardClaims
type AuthClaims struct {
	jwt.StandardClaims
}
