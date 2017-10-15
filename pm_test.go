package go_wypok

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWykopHandler_GetConversationsList(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.LoginToWypok()

	conversationsList, err := wh.GetConversationsList()

	assert.Nil(t, err)
	assert.NotNil(t, conversationsList)
}

func TestWykopHandler_SendPrivateMessageTo(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.LoginToWypok()

	content := "Testy integracyjne wrappera WykopAPI w GoLangu. Test wiadomości bez dodatkowych elementów."
	to := "m__b"
	succeeded, err := wh.SendPrivateMessageTo(to, content)
	assert.True(t, succeeded)
	assert.Nil(t, err)

	to = "interfacec"
	succeeded, err = wh.SendPrivateMessageTo(to, content)
	assert.False(t, succeeded)
	assert.NotNil(t, err)
	assert.Equal(t, err.ErrorObject.Code, 13)
}

func TestWykopHandler_SendPrivateMessageWithEmbeddedUrlTo(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.LoginToWypok()

	content := "Testy integracyjne wrappera WykopAPI w GoLangu. Test wiadomości z podlinkowanym obrazkiem."
	embed := "http://www.unixstickers.com/image/data/stickers/golang/golang.sh.png"
	to := "m__b"
	succeeded, err := wh.SendPrivateMessageWithEmbeddedUrlTo(to, content, embed)
	assert.True(t, succeeded)
	assert.Nil(t, err)

	to = "interfacec"
	succeeded, err = wh.SendPrivateMessageWithEmbeddedUrlTo(to, content, embed)
	assert.False(t, succeeded)
	assert.NotNil(t, err)
	assert.Equal(t, err.ErrorObject.Code, 13)
}

func TestWykopHandler_GetConversation(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.LoginToWypok()
	conversationTitle := "m__b"

	conversation, err := wh.GetConversation(conversationTitle)

	assert.NotNil(t, conversation)
	assert.Nil(t, err)
}

func TestWykopHandler_DeleteConversation(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	wh.LoginToWypok()

	conversationTitle := "m__b"

	succeeded, err := wh.DeleteConversation(conversationTitle)

	assert.True(t, succeeded)
	assert.Nil(t, err)
}

func TestWykopHandler_VaildatePMUrlGeneration(t *testing.T) {
	username := "m__b"
	wh := WykopHandler{}
	wh.authResponse.Userkey = "123"
	wh.appKey = "abc"

	sendMessageUrl := getSendMessageUrl(username, &wh)
	conversationsListUrl := getConversationsListUrl(&wh)
	conversationUrl := getConversationUrl(username, &wh)
	deleteConversationUrl := getDeleteConversationUrl(username, &wh)

	assert.Equal(t, sendMessageUrl, "https://a.wykop.pl/pm/SendMessage/m__b/appkey/abc/userkey/123")
	assert.Equal(t, conversationsListUrl, "https://a.wykop.pl/pm/ConversationsList/appkey/abc/userkey/123")
	assert.Equal(t, conversationUrl, "https://a.wykop.pl/pm/Conversation/m__b/appkey/abc/userkey/123")
	assert.Equal(t, deleteConversationUrl, "https://a.wykop.pl/pm/DeleteConversation/m__b/appkey/abc/userkey/123")
}
