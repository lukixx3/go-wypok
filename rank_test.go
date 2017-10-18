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
	assert.NotEmpty(t, profiles, "List of profiles in the rank was empty")
	for k, profile := range profiles {
		if k == len(profiles)-1 {
			continue
		}
		assert.True(t, profile.Rank <= profiles[k+1].Rank)
	}
}

func TestWykopHandlerGetRankWithSortingType(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()
	profiles, wypokError := wh.GetRankBySortingType(comment_count)

	assert.Nil(t, wypokError)
	assert.NotEmpty(t, profiles, "List of profiles in the rank was empty")

	for k, profile := range profiles {
		if k == len(profiles)-1 {
			continue
		}
		assert.True(t, profile.Comments >= profiles[k+1].Comments)
	}
}
