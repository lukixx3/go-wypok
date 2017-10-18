package go_wypok

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProfile(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	profile, profileError := wh.GetProfile("m__b")
	assert.Nil(t, profileError)
	assert.NotEmpty(t, profile.Email)
	assert.NotEmpty(t, profile.Login)
}

func TestGetProfileAdded(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	addedLinks, wypokError := wh.GetProfileAdded("m__b", 1)
	assert.Nil(t, wypokError)
	assert.NotEmpty(t, addedLinks)
	for _, link := range addedLinks {
		assert.NotEmpty(t, link.Url)
		assert.NotEmpty(t, link.Title)
	}
}

func TestGetProfileFavorites(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	favLinks, linkError := wh.GetProfileFavorites(wh.authResponse.Login)
	assert.Nil(t, linkError)
	assert.NotEmpty(t, favLinks)
	for _, link := range favLinks {
		assert.NotEmpty(t, link.Url)
		assert.NotEmpty(t, link.Title)
	}
}

func TestGetProfilePublished(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	favLinks, linkError := wh.GetProfilePublished("m__b", 1)
	assert.Nil(t, linkError)
	assert.NotEmpty(t, favLinks)
	for _, link := range favLinks {
		assert.NotEmpty(t, link.Url)
		assert.NotEmpty(t, link.Title)
	}
}

func TestGetProfileCommented(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	favLinks, linkError := wh.GetProfileCommented("m__b", 1)
	assert.Nil(t, linkError)
	assert.NotEmpty(t, favLinks)
	for _, link := range favLinks {
		assert.NotEmpty(t, link.Url)
		assert.NotEmpty(t, link.Title)
	}
}

func TestGetProfileDigged(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	diggedLinks, linkError := wh.GetProfileDigged("m__b", 1)
	assert.Nil(t, linkError)
	assert.NotEmpty(t, diggedLinks)
	for _, link := range diggedLinks {
		assert.NotEmpty(t, link.Url)
		assert.NotEmpty(t, link.Title)
	}
}

func TestGetProfileBuried(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	buriedLinks, linkError := wh.GetProfileBuried(wh.authResponse.Login, 1)
	assert.Nil(t, linkError)
	assert.NotEmpty(t, buriedLinks)
	for _, link := range buriedLinks {
		assert.NotEmpty(t, link.Url)
		assert.NotEmpty(t, link.Title)
	}
}

func TestObserveProfile(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	_, linkError := wh.ObserveProfile(wh.authResponse.Login)
	assert.NotNil(t, linkError, "Expected error, but there was none")
	assert.Equal(t, 33, linkError.ErrorObject.Code, "Expected error code 33, cannot observe yourself")

	followed, followedError := wh.ProfileFollowed(wh.authResponse.Login, 1)
	assert.Nil(t, followedError)
	assert.NotEmpty(t, followed)

	success, observeError := wh.ObserveProfile("UncleBilly")
	assert.Nil(t, observeError)
	assert.True(t, success, "Expected observe to succeed")

	followed2, followedError2 := wh.ProfileFollowed(wh.authResponse.Login, 1)
	assert.Nil(t, followedError2)
	assert.NotEmpty(t, followed2)

	assert.True(t, len(followed2) > len(followed), "Expected list of followed to have more elements after observe")

	unobserveSsuccess, unobserveError := wh.UnobserveProfile("UncleBilly")
	assert.Nil(t, unobserveError)
	assert.True(t, unobserveSsuccess, "Expected unobserve to succeed")
}

func TestBlockingAndUnblockingProfile(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	blocked, blockingError := wh.BlockProfile(wh.authResponse.Login)
	assert.NotNil(t, blockingError, "Expected error, cannot block yourself")
	assert.False(t, blocked)

	blocked2, blockingError2 := wh.BlockProfile("UncleBilly")
	assert.Nil(t, blockingError2)
	assert.True(t, blocked2)

	unblocked, unblockingError := wh.UnblockProfile("UncleBilly")
	assert.Nil(t, unblockingError)
	assert.True(t, unblocked)
}

func TestProfileFollowers(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	followers, followedError := wh.ProfileFollowers(wh.authResponse.Login, 1)
	assert.Nil(t, followedError)
	assert.NotEmpty(t, followers)

	for _, follower := range followers {
		assert.NotEmpty(t, follower.Login, "Follower login was empty, this is wrong")
	}
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
