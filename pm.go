package go_wypok

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/url"
)

const (
	PM_URL_PATTERN                  = "https://a.wykop.pl/pm/%s/appkey/%s/userkey/%s"
	PRIVATE_MESSAGE_POSITIVE_ANSWER = "[true]"
)

type Conversation struct {
	Author          string        `json:author`
	LastUpdate      WypokShitDate `json:last_update`
	AuthorAvatar    string        `json:author_avatar`
	AuthorAvatarBig string        `json:author_avatar_big`
	AuthorAvatarMed string        `json:author_avatar_med`
	AuthorAvatarLow string        `json:author_avatar_Lo`
	AuthorGroup     int           `json:author_group`
	AuthorSex       string        `json:author_sex`
	Status          string        `json:status`
}

type PrivateMessage struct {
	Author          string        `json:author`
	AuthorAvatar    string        `json:author_avatar`
	AuthorAvatarBig string        `json:author_avatar_big`
	AuthorAvatarMed string        `json:author_avatar_med`
	AuthorAvatarLow string        `json:author_avatar_Lo`
	AuthorGroup     int           `json:author_group`
	AuthorSex       string        `json:author_sex`
	Status          string        `json:status`
	Date            WypokShitDate `json:date`
	Content         string        `json:body`
	Direction       string        `json:direction`
	Embed           string        `json:embed`
	App             string        `json:app`
}

func (wh *WykopHandler) GetConversationsList() (conversationsList []Conversation, wykopError *WykopError) {
	urlAddress := getConversationsListUrl(wh)

	_, responseBody, _ := gorequest.New().Get(urlAddress).
									Set(contentType, mediaTypeFormType).
									Set(apiSignHeader, wh.hashRequest(urlAddress)).
									End()
	wykopError = wh.getObjectFromJson(responseBody, &conversationsList)

	return
}

func (wh *WykopHandler) SendPrivateMessageTo(to string, message string) (succeeded bool, wykopError *WykopError) {
	urlAddress := getSendMessageUrl(to, wh)

	body := url.Values{}
	body.Set("body", message)

	_, responseBody, _ := gorequest.New().Post(urlAddress).
									Set(contentType, mediaTypeFormType).
									Set(apiSignHeader, wh.hashRequest(urlAddress+body.Get("body"))).
									Send(body).
									End()
	succeeded = responseBody == PRIVATE_MESSAGE_POSITIVE_ANSWER

	if !succeeded {
		wykopError = wh.getObjectFromJson(responseBody, nil)
	}

	return
}

func (wh *WykopHandler) SendPrivateMessageWithEmbeddedUrlTo(to string, message string, embed string) (succeeded bool, wykopError *WykopError) {
	urlAddress := getSendMessageUrl(to, wh)

	body := url.Values{}
	body.Set("body", message)
	body.Set("embed", embed)

	_, responseBody, _ := gorequest.New().Post(urlAddress).
									Set(contentType, mediaTypeFormType).
									Set(apiSignHeader, wh.hashRequest(urlAddress+body.Get("body")+","+body.Get("embed"))).
									Send(body).
									End()
	succeeded = responseBody == PRIVATE_MESSAGE_POSITIVE_ANSWER

	if !succeeded {
		wykopError = wh.getObjectFromJson(responseBody, nil)
	}

	return
}

func (wh *WykopHandler) GetConversation(conversation string) (messages []PrivateMessage, wykopError *WykopError) {
	urlAddress := getConversationUrl(conversation, wh)

	_, responseBody, _ := gorequest.New().Get(urlAddress).
									Set(contentType, mediaTypeFormType).
									Set(apiSignHeader, wh.hashRequest(urlAddress)).
									End()
	wykopError = wh.getObjectFromJson(responseBody, &messages)

	return
}

func (wh *WykopHandler) DeleteConversation(conversation string) (succeeded bool, wykopError *WykopError) {
	urlAddress := getDeleteConversationUrl(conversation, wh)

	_, responseBody, _ := gorequest.New().Post(urlAddress).
									Set(contentType, mediaTypeFormType).
									Set(apiSignHeader, wh.hashRequest(urlAddress)).
									End()
	succeeded = responseBody == PRIVATE_MESSAGE_POSITIVE_ANSWER

	if !succeeded {
		wykopError = wh.getObjectFromJson(responseBody, nil)
	}

	return
}

func getConversationsListUrl(wh *WykopHandler) string {
	return fmt.Sprintf(PM_URL_PATTERN, "ConversationsList", wh.appKey, wh.authResponse.Userkey)
}

func getConversationUrl(conversation string, wh *WykopHandler) string {
	return fmt.Sprintf(PM_URL_PATTERN, "Conversation/" + conversation, wh.appKey, wh.authResponse.Userkey)
}

func getDeleteConversationUrl(conversation string, wh *WykopHandler) string {
	return fmt.Sprintf(PM_URL_PATTERN, "DeleteConversation/" + conversation, wh.appKey, wh.authResponse.Userkey)
}

func getSendMessageUrl(to string, wh *WykopHandler) string {
	return fmt.Sprintf(PM_URL_PATTERN, "SendMessage/" + to, wh.appKey, wh.authResponse.Userkey)
}
