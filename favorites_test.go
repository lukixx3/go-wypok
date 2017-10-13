package go_wypok

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWykopHandlerGetFavorites(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	favoritesID := os.Getenv("FAVORITES_ID")
	wh.LoginToWypok()

	if favoritesID != "" {
		links, wypokError := wh.GetFavorites(favoritesID)
		assert.Nil(t, wypokError, "Expected that error will be nil")
		assert.NotEmpty(t, links, "Expected that farovites list won't be empty")
	} else {
		t.Log("OS environment variable has not been set, testing empty favorites list")

		links, wypokError := wh.GetFavorites("0")
		assert.Nil(t, wypokError, "Expected that error will be nil")
		assert.Empty(t, links, "Expected that farovites list won't be empty")
	}
}
