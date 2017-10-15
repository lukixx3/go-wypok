package go_wypok

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProfile(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	profile, profileError := wh.GetProfile("interface")
	assert.Nil(t, profileError)
	assert.NotEmpty(t, profile.Email)
	assert.NotEmpty(t, profile.Login)
}

func TestGetProfileFavorites(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()
	wh.GetProfileFavorites("interface")
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
