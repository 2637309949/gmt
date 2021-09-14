package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

// Decode defined TODO
func Decode(jwtSecret, raw string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(raw, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	}
	return nil, err
}

// Encode defined TODO
func Encode(jwtSecret, issuer, user string, expireTime int64) (string, error) {
	claims := Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(jwtSecret))
}
