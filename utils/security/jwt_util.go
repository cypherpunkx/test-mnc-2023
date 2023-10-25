package security

import (
	"gorm-practice/config"
	"gorm-practice/models"
	"gorm-practice/utils/exception"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(customer *models.Customer) (string, error) {
	now := time.Now().UTC()
	end := now.Add(config.Cfg.TokenConfig.AccessTokenLifeTime)

	claims := &TokenMyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Cfg.TokenConfig.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: customer.UserName,
	}

	token := jwt.NewWithClaims(config.Cfg.TokenConfig.JWTSigningMethod, claims)
	ss, err := token.SignedString(config.Cfg.TokenConfig.JWTSignatureKey)

	if err != nil {
		return "", exception.ErrFailedCreateToken
	}

	return ss, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != config.Cfg.TokenConfig.JWTSigningMethod {
			return nil, exception.ErrInvalidTokenSigningMethod
		}
		return config.Cfg.TokenConfig.JWTSignatureKey, nil
	})

	if err != nil {
		return nil, exception.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid || claims["iss"] != config.Cfg.TokenConfig.ApplicationName {
		return nil, err
	}

	return claims, nil
}
