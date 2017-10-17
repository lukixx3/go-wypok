package go_wypok

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildingUrls(t *testing.T) {
	expectedGetEntryUrl := "https://a.wykop.pl/entries/index/999/appkey/APPKEY"
	expectedAddEntryUrl := "https://a.wykop.pl/entries/add/appkey/APPKEY/userkey/USERKEY"
	expectedEditEntryUrl := "https://a.wykop.pl/entries/edit/999/appkey/APPKEY/userkey/USERKEY"
	expectedDelEntryUrl := "https://a.wykop.pl/entries/delete/999/appkey/APPKEY/userkey/USERKEY"

	expectedAddEntryCommentUrl := "https://a.wykop.pl/entries/addComment/999/appkey/APPKEY/userkey/USERKEY"
	expectedEditEntryCommentUrl := "https://a.wykop.pl/entries/editComment/999/666/appkey/APPKEY/userkey/USERKEY"
	expectedDelEntryCommentUrl := "https://a.wykop.pl/entries/deleteComment/999/666/appkey/APPKEY/userkey/USERKEY"

	expectedEntryVoteUrl := "https://a.wykop.pl/entries/vote/entry/999/appkey/APPKEY/userkey/USERKEY"
	expectedEntryUnvoteUrl := "https://a.wykop.pl/entries/unvote/entry/999/appkey/APPKEY/userkey/USERKEY"

	expectedEntryCommentVoteUrl := "https://a.wykop.pl/entries/vote/comment/999/666/appkey/APPKEY/userkey/USERKEY"
	expectedEntryCommentUnoteUrl := "https://a.wykop.pl/entries/unvote/comment/999/666/appkey/APPKEY/userkey/USERKEY"

	expectedEntryFavoriteUrl := "https://a.wykop.pl/entries/favorite/999/appkey/APPKEY/userkey/USERKEY"

	appKey := "APPKEY"
	entryId := uint(999)
	commentId := uint(666)
	userKey := "USERKEY"

	assert.Equal(t, expectedGetEntryUrl, getEntryUrl(entryId, appKey))
	assert.Equal(t, expectedAddEntryUrl, getAddEntryUrl(appKey, userKey))
	assert.Equal(t, expectedEditEntryUrl, getEditEntryUrl(entryId, appKey, userKey))
	assert.Equal(t, expectedDelEntryUrl, getDeleteEntryUrl(entryId, appKey, userKey))
	assert.Equal(t, expectedAddEntryCommentUrl, getEntryAddCommentUrl(entryId, appKey, userKey))
	assert.Equal(t, expectedEditEntryCommentUrl, getEditEntryCommentUrl(entryId, commentId, appKey, userKey))
	assert.Equal(t, expectedDelEntryCommentUrl, getDeleteCommentUrl(entryId, commentId, appKey, userKey))
	assert.Equal(t, expectedEntryVoteUrl, getEntryVoteUrl(entry, entryId, commentId, appKey, userKey))
	assert.Equal(t, expectedEntryUnvoteUrl, getEntryUnvoteUrl(entry, entryId, commentId, appKey, userKey))
	assert.Equal(t, expectedEntryCommentVoteUrl, getEntryVoteUrl(comment, entryId, commentId, appKey, userKey))
	assert.Equal(t, expectedEntryCommentUnoteUrl, getEntryUnvoteUrl(comment, entryId, commentId, appKey, userKey))
	assert.Equal(t, expectedEntryFavoriteUrl, getEntryFavoriteUrl(entryId, appKey, userKey))
}

func TestWykopHandlerGetEntry(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	entry, wypokError := wh.GetEntry(0)
	assert.NotNil(t, wypokError)
	assert.Equal(t, 2, wypokError.ErrorObject.Code)
	assert.Equal(t, "Niepoprawne parametry", wypokError.ErrorObject.Message)

	assert.Equal(t, uint(0), entry.Id)
	assert.Empty(t, entry.Author)
}

func TestWykopHandlerGetCorrectEntry(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	entry, wypokError := wh.GetEntry(23391703)
	assert.Nil(t, wypokError)

	assert.Equal(t, uint(23391703), entry.Id)
	assert.NotEmpty(t, entry.Author)
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

	assert.Equal(t, entry.Id, newEntry.Id)

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

	newEntry, newEntryError := wh.GetEntry(response.Id)
	assert.Nil(t, newEntryError)
	assert.NotEmpty(t, newEntry.Embed.Url)

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

	newEntryResponse, wypokError := wh.PostEntry("testfsada sdas d")
	assert.Nil(t, wypokError)
	assert.NotEmpty(t, newEntryResponse.Id)

	comment := "Michau bialek ssie rogalek, a ten wpis zaraz zniknie ( ͡° ͜ʖ ͡°)"
	embed := "http://www.unixstickers.com/image/data/stickers/golang/golang.sh.png"

	firstCommentResponse, firstCommentError := wh.AddEntryComment(newEntryResponse.Id, comment)
	assert.Nil(t, firstCommentError)
	assert.NotEmpty(t, firstCommentResponse.Id)

	commentWithEmbedResp, commentWithEmbeddError := wh.AddEntryCommentWithEmbeddedContent(newEntryResponse.Id, comment, embed)
	assert.Nil(t, commentWithEmbeddError)
	assert.NotEmpty(t, commentWithEmbedResp.Id)

	newEntry, newEntryError := wh.GetEntry(newEntryResponse.Id)
	assert.Nil(t, newEntryError)
	assert.Len(t, newEntry.Comments, 2)
	assert.NotEmpty(t, newEntry.Comments[1].Embed.Url)

	delEntryResponse, delEntryError := wh.DeleteEntryComment(newEntryResponse.Id, firstCommentResponse.Id)
	assert.Nil(t, delEntryError)
	assert.Equal(t, firstCommentResponse.Id, delEntryResponse.Id)

	deleteEntryResponse, deleteEntryError := wh.DeleteEntry(newEntryResponse.Id)
	assert.Nil(t, deleteEntryError)
	assert.Equal(t, deleteEntryResponse.Id, newEntryResponse.Id)
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

	voteResponse, voteError := wh.UpvoteEntry(27336289)
	assert.Nil(t, voteError)
	assert.True(t, voteResponse.Vote > 0)
}

func TestWykopHandlerUnvoteEntry(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	voteResponse, voteError := wh.UnvoteEntry(27336289)
	assert.Nil(t, voteError)
	assert.True(t, voteResponse.Vote == 0, "This might fail and this is ok.")
}

func TestWykopHandlerFavoriteEntry(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	// add to favorite
	favResponse, favResponseError := wh.FavoriteEntry(27336289)
	assert.Nil(t, favResponseError)
	assert.True(t, favResponse.UserFavorite)

	// re-add to favorite = unfavorite
	unfavResponse, unfavResponseError := wh.FavoriteEntry(27336289)
	assert.Nil(t, unfavResponseError)
	assert.False(t, unfavResponse.UserFavorite)
}
