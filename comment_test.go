package go_wypok

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWykopHandler_AddComment(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	response, err := wh.AddComment(1760442, 0, "Test integracyjny dodawania komentarza", "")

	assert.Nil(t, err)
	assert.NotEqual(t, response.Id, 0)

	response, err = wh.AddComment(1760442, 0, "Test integracyjny dodawania komentarza z obrazkię", "http://www.unixstickers.com/image/data/stickers/golang/golang.sh.png")

	assert.Nil(t, err)
	assert.NotEqual(t, response.Id, 0)
}

func TestWykopHandler_PlusComment(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	vote, err := wh.PlusComment(1760442, 18549048)

	assert.Nil(t, err)
	assert.NotEqual(t, vote.Vote, FAKE_VOTE_VALUE)
	assert.NotEqual(t, vote.VotePlus, FAKE_VOTE_VALUE)
	assert.NotEqual(t, vote.VoteMinus, FAKE_VOTE_VALUE)
	// Those below are waiting for better times at WykopApi, cause seems that you cannot revoke your "plus" action
	//votePlus := vote.VotePlus
	//vote, err = wh.PlusComment(1760442, 18549048)
	//
	//assert.NotNil(t, vote)
	//assert.Nil(t, err)
	//assert.Equal(t, vote.VotePlus, votePlus-1)
}

func TestWykopHandler_MinusComment(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	vote, err := wh.MinusComment(1760442, 18549048)

	assert.Nil(t, err)
	assert.NotEqual(t, vote.Vote, FAKE_VOTE_VALUE)
	assert.NotEqual(t, vote.VotePlus, FAKE_VOTE_VALUE)
	assert.NotEqual(t, vote.VoteMinus, FAKE_VOTE_VALUE)

	// Those below are waiting for better times at WykopApi, cause seems that you cannot revoke your "plus" action
	//voteMinus := vote.VoteMinus
	//vote, err = wh.MinusComment(1760442, 18549048)
	//
	//assert.NotNil(t, vote)
	//assert.Nil(t, err)
	//assert.Equal(t, vote.VoteMinus, voteMinus-1)
}

func TestWykopHandler_EditComment(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	newComment, err := wh.AddComment(1760442, 0, "Test integracyjny edycji komentarza", "")

	assert.NotEqual(t, newComment.Id, 0)
	assert.Nil(t, err)

	editedComment, err := wh.EditComment(newComment.Id, "Wyedytowany komentarz")

	assert.NotEqual(t, editedComment.Id, 0)
	assert.Nil(t, err)
}

func TestWykopHandler_DeleteComment(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	wh.LoginToWypok()

	newComment, err := wh.AddComment(1760442, 0, "Test integracyjny usuwania komentarza", "")

	assert.NotEqual(t, newComment.Id, 0)
	assert.Nil(t, err)

	deletedComment, err := wh.DeleteComment(newComment.Id)

	assert.NotEqual(t, deletedComment.Id, 0)
	assert.Nil(t, err)
}

func TestWykopHandler_LinkCommentsUrlBuilder(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh := WykopHandler{}
	wh.authResponse.Userkey = "123"
	wh.appKey = "abc"

	addCommentUrl := getAddCommentUrl(1, 2, &wh)       // With parent comment
	addParentCommentUrl := getAddCommentUrl(1, 0, &wh) // Without parent comment
	plusCommentUrl := getPlusLinkCommentUrl(1, 2, &wh)
	minusCommentUrl := getMinusLinkCommentUrl(1, 2, &wh)
	deleteCommentUrl := getDeleteLinkCommentUrl(2, &wh)
	editCommentUrl := getEditLinkCommentUrl(2, &wh)

	assert.Equal(t, addCommentUrl, "https://a.wykop.pl/comments/add/1/2/appkey/abc/userkey/123")
	assert.Equal(t, addParentCommentUrl, "https://a.wykop.pl/comments/add/1/appkey/abc/userkey/123")
	assert.Equal(t, plusCommentUrl, "https://a.wykop.pl/comments/plus/1/2/appkey/abc/userkey/123")
	assert.Equal(t, minusCommentUrl, "https://a.wykop.pl/comments/minus/1/2/appkey/abc/userkey/123")
	assert.Equal(t, deleteCommentUrl, "https://a.wykop.pl/comments/delete/2/appkey/abc/userkey/123")
	assert.Equal(t, editCommentUrl, "https://a.wykop.pl/comments/edit/2/appkey/abc/userkey/123")
}
