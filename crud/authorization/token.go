package authorization

import (
	"errors"
	"github.com/cmorales95/golang_api/crud/models"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

//GenerateToken .
func GenerateToken(data *models.Login) (string, error) {
	// claim
	claim := models.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "MyCompany",
		},
	}

	// preparing token to sign
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	// token is signed
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return signedToken, nil
}

//ValidateToken .
func ValidateToken(t string) (models.Claim, error) {
	// parse the token with the public certificate
	token, err := jwt.ParseWithClaims(t,&models.Claim{}, verifyFunction)
	if err != nil {
		return models.Claim{}, err
	}
	if !token.Valid {
		return models.Claim{}, errors.New("token is not valid")
	}

	// validate the data and cast to the cliams struct
	claim, ok := token.Claims.(*models.Claim)
	if !ok {
		return models.Claim{}, errors.New("error getting claim")
	}

	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	// verifyKey is the public certificate
	return verifyKey, nil
}
