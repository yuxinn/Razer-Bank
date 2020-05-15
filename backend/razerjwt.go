package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a%*84*2L7MTz77aO2wlcTdECpPQ7msk1D5$MN1c@&8Gsl0$QeOyeyylzg96AnZNsEWn8i6!U3ZE8Q$r1N7AbqgsVaj1AAhN6969")

// var jwtKey = techchillaconfig.GetJWTKey()

type JWTAccountClaims struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	NRIC      string `json:"nric"`
	Mambukey  string `json:"mambuKey"`
	jwt.StandardClaims
}

func GenerateJWT(email string, firstname string, lastname string, nric string, mambukey string) string {
	expirationTime := time.Now().Add(180 * time.Minute)
	claims := JWTAccountClaims{
		email,
		firstname,
		lastname,
		nric,
		mambukey,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "RazerBank",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	ss, err := token.SignedString(jwtKey)
	if err != nil {
		//fmt.Printf("%v %v", ss, err)
		return ""
	}

	return ss
}

func GetEmailFromJWT(tknStr string) string {
	claims := &JWTAccountClaims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	_ = tkn
	_ = err
	return claims.Email
}

func VerifyJWT(tknStr string) bool {
	claims := &JWTAccountClaims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false
		}
		return false
	}
	if !tkn.Valid {
		return false
	}
	return true
}
