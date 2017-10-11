package go_wypok

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestWykopHandlerGetEntry(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	entry, wypokError := wh.GetEntry("0")
	assert.NotNil(t, wypokError)
	assert.Equal(t, 2, wypokError.ErrorObject.Code)
	assert.Equal(t, "Niepoprawne parametry", wypokError.ErrorObject.Message)

	assert.Equal(t, 0, entry.Id)
	assert.Empty(t, entry.Author)
}

func TestWykopHandlerPostEntry(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	newEntry, wypokError := wh.PostEntry("test")

	assert.Nil(t, wypokError)
	assert.NotEmpty(t, newEntry.Id)

	entry, getEntryError := wh.GetEntry(newEntry.Id)
	assert.Nil(t, getEntryError)
	assert.NotEmpty(t, entry.Author)

	assert.Equal(t, strconv.Itoa(entry.Id), newEntry.Id)

	deleteEntryResponse, deleteEntryError := wh.DeleteEntry(newEntry.Id)
	assert.Nil(t, deleteEntryError)

	assert.Equal(t, deleteEntryResponse.Id, newEntry.Id)
}

func TestWykopHandlerDeleteEntry(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	newEntry, wypokError := wh.PostEntry("test deleting entry")
	assert.Nil(t, wypokError)
	assert.NotEmpty(t, newEntry.Id)

	deleteEntryResponse, deleteEntryError := wh.DeleteEntry(newEntry.Id)
	assert.Nil(t, deleteEntryError)
	assert.Equal(t, deleteEntryResponse.Id, newEntry.Id)
}

func TestWykopHandlerAddEntryComment(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	newEntry, wypokError := wh.PostEntry("testfsada sdas d")
	assert.Nil(t, wypokError)
	assert.NotEmpty(t, newEntry.Id)

	addCommentResponse, addCommentError := wh.AddEntryComment(newEntry.Id, "Michau bialek ssie rogalek, a ten wpis zaraz zniknie ( ͡° ͜ʖ ͡°)")
	assert.Nil(t, addCommentError)
	assert.NotEmpty(t, addCommentResponse.Id)

	delEntryResponse, delEntryError := wh.DeleteEntryComment(newEntry.Id, addCommentResponse.Id)
	assert.Nil(t, delEntryError)
	assert.Equal(t, addCommentResponse.Id, delEntryResponse.Id)

	deleteEntryResponse, deleteEntryError := wh.DeleteEntry(newEntry.Id)
	assert.Nil(t, deleteEntryError)
	assert.Equal(t, deleteEntryResponse.Id, newEntry.Id)
}

func TestWykopHandlerEditEntry(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	newEntry, wypokError := wh.PostEntry("test editing entry")
	assert.Nil(t, wypokError)
	assert.NotEmpty(t, newEntry.Id)

	newEntryContent := "new entry content"

	editEntryResponse, editEntryError := wh.EditEntry(newEntry.Id, newEntryContent)
	assert.Nil(t, editEntryError)
	assert.NotEmpty(t, editEntryResponse.Id)

	editedEntry, editedEntryError := wh.GetEntry(newEntry.Id)
	assert.Nil(t, editedEntryError)
	assert.NotEmpty(t, editEntryResponse.Id)
	assert.Equal(t, newEntryContent, editedEntry.Body)

	deleteEntryResponse, deleteEntryError := wh.DeleteEntry(newEntry.Id)
	assert.Nil(t, deleteEntryError)
	assert.Equal(t, deleteEntryResponse.Id, newEntry.Id)
}

