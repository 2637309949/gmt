package jwt

import "github.com/dgrijalva/jwt-go"

// Claims defined TODO
type Claims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

// GetUser defined TODO
func (c *Claims) GetUser() string {
	return c.User
}
