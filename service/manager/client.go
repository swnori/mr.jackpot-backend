package manager

import (
	"time"

	"github.com/golang-jwt/jwt"
)



type AuthService interface {
	CreateToken(userid int, status string) (string, error)
	ParseToken(tokenString string) (id int, status string, err error)
	GetAccessTokenKey(token *jwt.Token) (interface{}, error)
	GetAccessExpireTime() time.Duration
}

type Client struct {
	Token TokenService
}

func (c *Client) CreateToken(userid int, status string) (string, error) {
	return c.Token.CreateToken(userid, status)
}
func (c *Client) ParseToken(tokenString string) (id int, status string, err error) {
	return c.Token.ParseToken(tokenString)
}
func (c *Client) GetAccessTokenKey(token *jwt.Token) (interface{}, error) {
	return c.Token.GetAccessTokenKey(token)
}

func (c *Client) GetAccessExpireTime() time.Duration {
	return c.Token.GetAccessExpireTime()
}