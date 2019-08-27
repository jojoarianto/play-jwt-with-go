package section1

import (
	"testing"
)

func TestCreate(t *testing.T) {
	token := Create()
	t.Log(token)
}

func TestCreateWithClaim(t *testing.T) {
	token := CreateWithClaim()
	t.Log(token)
}

func TestCreateWithClaimAndSecret(t *testing.T) {
	tokenString, _ := CreateWithClaimAndSecret()
	t.Log(tokenString)
}

func TestCreateWithClaimAndSecretAndExpired(t *testing.T) {
	tokenString, _ := CreateWithClaimAndSecretAndExpired()
	t.Log(tokenString)
}

func TestParseToken(t *testing.T) {
	ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJoZWxsbyBmb28iLCJleHAiOjE1NjY5NDUzNDgsImlzcyI6ImlyaWFudG8ifQ.5vEicxZTG0oknJrkQLJNyqtW1nd9_qoGnbPZodR_mKQ")
}
