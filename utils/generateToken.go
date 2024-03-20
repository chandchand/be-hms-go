package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// const secretKey = "aptx4869"

func GenerateToken(userID, firstName, username, role string) (string, error) {
	claims := jwt.MapClaims{
		"userID":     userID,
		"first_name": firstName,
		"username":   username,
		"role":       role,
		"exp":        time.Now().Add(time.Hour * 1).Unix(), // Tambahkan waktu kadaluwarsa sesuai kebutuhan Anda.
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
