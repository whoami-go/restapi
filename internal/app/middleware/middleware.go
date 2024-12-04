package middleware

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

var (
	SecretKey      []byte      = []byte("UltraRestApikey4000")
	emptyValidFunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	}
)

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: emptyValidFunc,
	SigningMethod:       jwt.SigningMethodES256,
})
