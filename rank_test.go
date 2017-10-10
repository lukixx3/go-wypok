package go_wypok

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWykopHandlerGetRank(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()
	profiles, wypokError := wh.GetRank()

	assert.Nil(t, wypokError)
	assert.True(t, len(profiles) > 0)
}

func TestWykopHandlerGetRankWithSortingType(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()
	profiles, wypokError := wh.GetRankBySortingType(comment_count)

	assert.Nil(t, wypokError)
	assert.True(t, len(profiles) > 0)
}
