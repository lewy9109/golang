package helper

import "github.com/golang-jwt/jwt/v5"

func CreateJWTToken(userID uint, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	Ptoken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return Ptoken, nil
}

func ValidateJWTToken(tokenSting string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenSting, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return token, err
}
