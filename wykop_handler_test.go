package go_wypok

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
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

func TestGettingProfileEntriesComments(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	profileEntryComments, wypokError := wh.GetProfileEntriesComments("m__b", 1)
	assert.Nil(t, wypokError)
	assert.NotNil(t, profileEntryComments)
	assert.True(t, len(profileEntryComments) > 0, "Looks like user m__b doesnt have any comments on links")

	for _, entryComment := range profileEntryComments {
		assert.NotEqual(t, "", entryComment.Body, "Expected body of an entryComment to be populated")
	}
}

func TestGettingProfileLinksComments(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	profileComments, wypokError := wh.GetProfileComments("m__b", 1)
	assert.Nil(t, wypokError)
	assert.NotNil(t, profileComments)
	assert.True(t, len(profileComments) > 0, "Looks like user m__b doesnt have any comments on links")

	for _, comment := range profileComments {
		assert.NotEqual(t, "", comment.Body, "Expected body of a comment to be populated")
	}
}

func TestWykopHandler_PostEntryWithEmbeddedContent(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	content := "Test"
	embed := "http://www.unixstickers.com/image/data/stickers/golang/golang.sh.png"

	response, wykopError := wh.PostEntryWithEmbeddedContent(&content, &embed)
	assert.Nil(t, wykopError)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Id)

	entryId, _ := strconv.Atoi(response.Id)
	deleteResponse, deleteResponseError := wh.DeleteEntry(entryId)
	assert.Nil(t, deleteResponseError, "Expected no error deleting entry")
	assert.NotNil(t, deleteResponse)

	deletedEntryId, _ := strconv.Atoi(deleteResponse.Id)
	assert.Equal(t, entryId, deletedEntryId)
}

//func TestUploadingEntryWithImage(t *testing.T) {
//	teardownTestCase := setupTestCase(t)
//	defer teardownTestCase(t)
//	wh.LoginToWypok()
//
//	entryBody := "test"
//
//	entryResponse, wypokError := wh.PostEntryWithImage(entryBody, "/home/agilob/Pictures/penguin_wings.jpg")
//	assert.Nil(t, wypokError)
//	assert.NotNil(t, entryResponse.Id)
//
//	entryId, _ := strconv.Atoi(entryResponse.Id)
//	entry, errorGettingEntry := wh.GetEntry(entryId)
//	assert.Nil(t, errorGettingEntry, "Expected no error getting entry that was created before")
//	assert.Equal(t, entryBody, entry.Body, "Message body is not what was submitted")
//
//	// assert here that entry.Embed is populated
//
//	deleteResponse, deleteResponseError := wh.DeleteEntry(entryId)
//	assert.Nil(t, deleteResponseError, "Expected no error deleting entry")
//	assert.Equal(t, entryResponse.Id, deleteResponse.Id)
//}
