package services

import (
	"time"

	//"github.com/golang-jwt/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jedzeins/jlpt_api/userService/src/models"
)

var (
	secret1 = []byte("3456yujhgfde4%^&YG")
	secret2 = []byte("<MnbgyUIJKmnhy&*IK")
)

func GenerateSignedTokens(ownerId int, sessionId int) (string, string, *models.ApiError) {

	expTime := time.Now().Add(time.Hour * 24 * 7).Unix()

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionId": sessionId,
		"nbf":       expTime,
	})

	refreshTokenString, err := refreshToken.SignedString(secret1)

	if err != nil {
		return "", "", &models.ApiError{ErrorMessage: err.Error()}
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ownerId":   ownerId,
		"sessionId": sessionId,
		"nbf":       expTime,
	})

	accessTokenString, err := accessToken.SignedString(secret2)

	if err != nil {
		return "", "", &models.ApiError{ErrorMessage: err.Error()}
	}

	return refreshTokenString, accessTokenString, nil

}
