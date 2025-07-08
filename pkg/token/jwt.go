package token

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userUUID string) (string, error) {
    secret := []byte(os.Getenv("JWT_SECRET"))

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "uuid": userUUID,
        "exp":  time.Now().Add(72 * time.Hour).Unix(),
    })

    return token.SignedString(secret)
}

func ValidateJWT(tokenStr string) (map[string]interface{}, error) {
    secret := []byte(os.Getenv("JWT_SECRET"))

    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return secret, nil
    })

    if err != nil || !token.Valid {
        return nil, errors.New("invalid or expired token")
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        return claims, nil
    }
    return nil, errors.New("invalid token claims")
}
