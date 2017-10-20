package go_wypok

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var wh WykopHandler

func setupTestCase(t *testing.T) func(t *testing.T) {
	wh = WykopHandler{}
	wykopSecret := os.Getenv("WYKOPSECRET")
	wykopAppKey := os.Getenv("WYKOPAPPKEY")
	wykopConnKey := os.Getenv("WYKOPCONNECTIONKEY")

	wh.SetAppKey(wykopAppKey)
	wh.SetSecret(wykopSecret)
	wh.SetConnectionKey(wykopConnKey)

	return func(t *testing.T) {
		wh = WykopHandler{}
	}
}

func TestLogin(t *testing.T) {
	localWh := WykopHandler{}
	assert.Equal(t, "", localWh.authResponse.Userkey, "UserKey should be empty before login")
	wypokError := localWh.LoginToWypok()
	assert.Equal(t, "", localWh.authResponse.Userkey, "UserKeygit  should be empty when login failed")
	assert.Equal(t, "Niepoprawny klucz API", wypokError.ErrorObject.Message, "Expected login to fail and get error message")
	assert.Equal(t, 1, wypokError.ErrorObject.Code, "Expected login to fail and get error message")
}

