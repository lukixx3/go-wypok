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

func TestWykopHandlerPostEntryWithEmbeddedContent(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	content := "Test"
	embed := "http://www.unixstickers.com/image/data/stickers/golang/golang.sh.png"

	response, wykopError := wh.PostEntryWithEmbeddedContent(content, embed)
	assert.Nil(t, wykopError)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Id)

	deleteResponse, deleteResponseError := wh.DeleteEntry(response.Id)
	assert.Nil(t, deleteResponseError, "Expected no error deleting entry")
	assert.NotNil(t, deleteResponse)

	assert.Equal(t, response.Id, deleteResponse.Id)
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

func TestWykopHandlerEditEntryComment(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	newEntry, wypokError := wh.PostEntry("test editing entry comment")
	assert.Nil(t, wypokError)
	assert.NotEmpty(t, newEntry.Id)

	newCommentContent := "comment content"

	newComment, addCommentError := wh.AddEntryComment(newEntry.Id, "Michau biauek ssie rogalek, a ten wpis zaraz zniknie ( ͡° ͜ʖ ͡°)")
	assert.Nil(t, addCommentError)
	assert.NotEmpty(t, newComment.Id)

	editCommentResponse, editCommentError := wh.EditEntryComment(newEntry.Id, newComment.Id, newCommentContent)
	assert.Nil(t, editCommentError)
	assert.NotEmpty(t, editCommentResponse.Id)

	entry, entryError := wh.GetEntry(newEntry.Id)
	assert.Nil(t, entryError)
	assert.NotEmpty(t, entry.Id)
	assert.Equal(t, newCommentContent, entry.Comments[0].Body)

	deleteEntryResponse, deleteEntryError := wh.DeleteEntry(newEntry.Id)
	assert.Nil(t, deleteEntryError)
	assert.Equal(t, deleteEntryResponse.Id, newEntry.Id)
}

func TestWykopHandlerUpvoteEntry(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	voteResponse, voteError := wh.UpvoteEntry("27336289")
	assert.Nil(t, voteError)
	assert.True(t, voteResponse.Vote > 0)
}

func TestWykopHandlerUnvoteEntry(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	voteResponse, voteError := wh.UnvoteEntry("27336289")
	assert.Nil(t, voteError)
	assert.True(t, voteResponse.Vote == 0, "This might fail and this is ok.")
}

func TestWykopHandlerFavoriteEntry(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	// add to favorite
	favResponse, favResponseError := wh.FavoriteEntry("27336289")
	assert.Nil(t, favResponseError)
	assert.True(t, favResponse.UserFavorite)

	// re-add to favorite = unfavorite
	unfavResponse, unfavResponseError := wh.FavoriteEntry("27336289")
	assert.Nil(t, unfavResponseError)
	assert.False(t, unfavResponse.UserFavorite)
}

