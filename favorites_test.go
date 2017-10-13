package go_wypok

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWykopHandlerGetFavoritesLists(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.LoginToWypok()

	lists, wypokError := wh.GetFavoritesLists()
	assert.Nil(t, wypokError, "Expected that error will be nil")
	assert.NotEmpty(t, lists, "Expected that farovites list won't be empty")
}

func TestWykopHandlerGetFavoritesListLinks(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.LoginToWypok()

	lists, _ := wh.GetFavoritesLists()
	if len(lists) > 0 {
		links, wypokError := wh.GetFavoritesListLinks(strconv.Itoa(lists[0].ID))
		assert.Nil(t, wypokError, "Expected that error will be nil")
		assert.NotEmptyf(t, links, "Expected that links from favorites list '%s' won't be empty", lists[0].Name)
	} else {
		t.Error("GetFavoritesLists probably failed, lists are empty, checking not existing list")
	}
}

func TestWykopHandlerGetFavoritesComments(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.LoginToWypok()

	comments, wypokError := wh.GetFavoritesComments()
	assert.Nil(t, wypokError, "Expected that error will be nil")
	assert.NotEmpty(t, comments, "Expected that farovites list won't be empty")
}

func TestWykopHandlerGetFavoritesEntries(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.LoginToWypok()

	lists, wypokError := wh.GetFavoritesEntries()
	assert.Nil(t, wypokError, "Expected that error will be nil")
	assert.NotEmpty(t, lists, "Expected that farovites list won't be empty")
}
