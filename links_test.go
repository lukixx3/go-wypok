package go_wypok

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGettingMainPageWithFakeUserKey(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.authResponse.Userkey = "dsadasdas"

	links, wypokError := wh.GetMainPageLinks(1)

	assert.NotNil(t, wypokError)
	assert.Equal(t, 0, len(links))

	assert.Equal(t, "Niepoprawny klucz użytkownika", wypokError.ErrorObject.Message, "Expected login to fail and get error message")
	assert.Equal(t, 11, wypokError.ErrorObject.Code, "Expected login to fail and get error message")
}

func TestGettingMainPage(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	assert.Equal(t, "", wh.authResponse.Userkey, "UserKey should be empty before login")
	// send unauthorized request so we expect to get nothing as userVote
	links, wypokError := wh.GetMainPageLinks(1)

	assert.Equal(t, 50, len(links))
	for _, v := range links {
		assert.Equal(t, "false", string(v.UserVote))
	}
	assert.Nil(t, wypokError)

	// now login to wypok and send request as user
	wh.LoginToWypok()
	assert.NotEqual(t, "", wh.authResponse.Userkey, "Expected userKey to be populated after login")

	links, wypokError = wh.GetMainPageLinks(1)

	assert.Equal(t, 50, len(links))
	for _, v := range links {
		// userVote field might be dig, bury or false, but never empty string
		assert.NotEqual(t, "", string(v.UserVote))
	}
	assert.Nil(t, wypokError)
}

func TestGettingUpcomingPage(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	assert.Equal(t, "", wh.authResponse.Userkey, "UserKey should be empty before login")
	links, wypokError := wh.GetUpcomingLinks(1)

	assert.Equal(t, 50, len(links))
	for _, v := range links {
		assert.Equal(t, "false", string(v.UserVote))
	}
	assert.Nil(t, wypokError)

	// now login to wypok and send request as user, some fields might be dig, bury or false, but never empty string
	wh.LoginToWypok()
	assert.NotEqual(t, "", wh.authResponse.Userkey, "Expected userKey to be populated after login")

	links, wypokError = wh.GetUpcomingLinks(1)
	assert.Equal(t, 50, len(links))
	for _, v := range links {
		// userVote field might be dig, bury or false, but never empty string
		assert.NotEqual(t, "", string(v.UserVote))
	}
	assert.Nil(t, wypokError)

}
