package manager

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"mr.jackpot-backend/model"
	"mr.jackpot-backend/utility/util"
)

type TokenService interface {
	CreateToken(userid int, status string) (string, error)
	ParseToken(tokenString string) (id int, status string, err error)
	GetAccessTokenKey(token *jwt.Token) (interface{}, error)
}

type Token struct {
	_ACCESS_SECRET string
	_AccessTokenExpiredAt time.Duration
}

var DefaultToken = &Token{
	_ACCESS_SECRET: util.GetRandomString(24),
	_AccessTokenExpiredAt: time.Hour * 24,
}

func (m *Token) CreateToken(userid int, status string) (string, error) {
	var atd model.TokenDetails

	atd.Expires = time.Now().Add(m.GetAccessExpireTime()).Unix()
	atd.Uuid = uuid.NewV4().String()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = atd.Uuid
	atClaims["user_id"] = userid
	atClaims["user_status"] = status
	atClaims["exp"] = atd.Expires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(m.GetAccessSecret()))
	return token, err
}

func (m *Token) ParseToken(tokenString string) (id int, status string, err error) {
	tokenParsed, err := jwt.Parse(tokenString, m.GetAccessTokenKey)
	if err != nil {
		return
	}

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !(ok && tokenParsed.Valid) {
		err = errors.New("")
		return
	}

	status, ok = claims["user_status"].(string)
	if !ok {
		err = errors.New("")
		return
	}

	userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		err = errors.New("")
		return
	}

	id = int(userId)
	err = nil
	return
}

func (m *Token) GetAccessTokenKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(m.GetAccessSecret()), nil
}

func (m *Token) GetAccessSecret() string {
	return m._ACCESS_SECRET
}

func (m *Token) GetAccessExpireTime() time.Duration {
	return m._AccessTokenExpiredAt
}

