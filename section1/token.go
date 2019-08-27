package section1

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

func Create() string {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		log.Fatal(err.Error())
	}

	return token
}

func CreateWithClaim() string {
	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"user_id": 22,
	})
	token, err := sign.SignedString([]byte("iriantosecret"))
	if err != nil {
		log.Fatal(err.Error())
	}

	return token
}

func CreateWithClaimAndSecret() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 25,
		"email": "iriantogo@gmail.com",
	})

	tokenString, err := token.SignedString([]byte("AllYourBase"))
	fmt.Println(tokenString)
	return tokenString, err
}

func CreateWithClaimAndSecretAndExpired() (string, error) {
	mySigningKey := []byte("iriantosecret")

	expireToken := time.Now().Add(time.Second * 20).Unix()

	type MyCustomClaim struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// create claim
	claims := MyCustomClaim{
		Foo: "hello foo",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer: "irianto",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

func ParseToken(tokenStr string)  {
	mySigningKey := []byte("iriantosecret")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return mySigningKey, nil
	})

	if token.Valid {
		fmt.Println("token is valid")

	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors & jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors & (jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token here:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
}