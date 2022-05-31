package auth

import (
	"strings"
	"study-service/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

type UserClaim struct {
	Token       string   `json:"-"`
	Id          int64    `json:"id"`
	Nbf         int64    `json:"-"`
	Exp         int64    `json:"-"`
	Nickname    string   `json:"-"`
	Aud         string   `json:"-"`
	Iss         string   `json:"-"`
	Name        string   `json:"name"`
	Datetime    string   `json:"datetime"`
}

func CreateToken(id int64, nickname string, name string, roles ...string) (string, error) {
	log.Traceln("")

	expireTime := time.Now().Add(time.Hour * 24)

	claims := jwt.MapClaims{
		"id":       id,
		"nickname": nickname,
		"name":     name,
		"roles":    strings.Join(roles, ","),
		"iss":      "sharing_platform",
		"aud":      "sharing_platform",
		"nbf":      time.Now().Add(-time.Minute * 5).Unix(),
		"exp":      expireTime.Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(config.Config.JwtSecret))
}

func CreateTokenWithExpire(id int64, nickname string, orgid int64, name string, durationDays int, roles ...string) (string, error) {
	log.Traceln("")

	expireTime := time.Now().Add(time.Hour * 24 * time.Duration(durationDays))

	claims := jwt.MapClaims{
		"id":       id,
		"nickname": nickname,
		"orgid":    orgid,
		"name":     name,
		"roles":    strings.Join(roles, ","),
		"iss":      "sharing_platform",
		"aud":      "sharing_platform",
		"nbf":      time.Now().Add(-time.Minute * 5).Unix(),
		"exp":      expireTime.Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(config.Config.JwtSecret))
}


