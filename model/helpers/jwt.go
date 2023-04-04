package helpers

import "github.com/golang-jwt/jwt/v5"

func CreateJWTToken(userID int, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	at := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
