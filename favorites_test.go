package go_wypok

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWykopHandlerGetFavoritesLists(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.LoginToWypok()

	lists, wypokError := wh.GetFavoritesLists()
	assert.Nil(t, wypokError, "Expected that error will be nil")
	assert.NotEmpty(t, lists, "Expected that favorites list won't be empty")
}

func TestWykopHandlerGetFavoritesListLinks(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.LoginToWypok()

	lists, _ := wh.GetFavoritesLists()
	assert.NotEmpty(t, lists, "Expected that favorites list won't be empty")

	links, wypokError := wh.GetFavoritesListLinks(lists[0].Id)
	assert.Nil(t, wypokError, "Expected that error will be nil")
	assert.NotEmptyf(t, links, "Expected that links from favorites list '%s' won't be empty", lists[0].Name)
}

func TestWykopHandlerGetFavoritesComments(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.LoginToWypok()

	comments, wypokError := wh.GetFavoritesComments()
	assert.Nil(t, wypokError, "Expected that error will be nil")
	assert.NotEmpty(t, comments, "Expected that favorites list won't be empty")
}

func TestWykopHandlerGetFavoritesEntries(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.LoginToWypok()

	lists, wypokError := wh.GetFavoritesEntries()
	assert.Nil(t, wypokError, "Expected that error will be nil")
	assert.NotEmpty(t, lists, "Expected that favorites list won't be empty")
}

func TestBuildingFavoritesURLs(t *testing.T) {
	appKey := "APPKEY"
	favoritesId := 999
	userKey := "USERKEY"

	assert.Equal(t,
		"https://a.wykop.pl/favorites/index/999/appkey/APPKEY/userkey/USERKEY",
		getFavoritesListLinksURL(favoritesId, appKey, userKey),
	)
	assert.Equal(t,
		"https://a.wykop.pl/favorites/lists/appkey/APPKEY/userkey/USERKEY",
		getFavoritesListsURL(appKey, userKey),
	)
	assert.Equal(t,
		"https://a.wykop.pl/favorites/comments/appkey/APPKEY/userkey/USERKEY",
		getFavoritesCommentsURL(appKey, userKey),
	)
	assert.Equal(t,
		"https://a.wykop.pl/favorites/entries/appkey/APPKEY/userkey/USERKEY",
		getFavoritesEntriesURL(appKey, userKey),
	)
}
