package jwt

import (
	"blog-system/config"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a3f8b2c9d1e7f6g4h5i0j2k8l9m1n7o3p6q4r8s2t5u9v1w7x3y0z4")

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(userID uint, username string) (string, error) {
	cfg := config.GetConfig()
	expireHours := cfg.JWT.ExpireHours
	if expireHours == 0 {
		expireHours = 24
	}

	expirationTime := time.Now().Add(time.Duration(expireHours) * time.Hour)
	claims := &Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    cfg.JWT.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.JWT.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JWT.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, errors.New("无效的令牌")
	}

	return claims, nil
}
