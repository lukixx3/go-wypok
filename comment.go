package go_wypok

import (
	"fmt"
	"net/url"
	"github.com/parnurzeal/gorequest"
)

const (
	COMMENTS_BASE_URL = "https://a.wykop.pl/comments/%s/appkey/%s/userkey/%s"
	FAKE_VOTE_VALUE = -9999
)

type Vote struct {
	Vote      int `json:"vote"`
	VotePlus  int `json:"vote_plus"`
	VoteMinus int `json:"vote_minus"`
}

/*
 *
 * parentCommentId - for a comment directly to link (so no parent comment) set to 0
 *
 */
func (wh *WykopHandler) AddComment(linkId uint, parentCommentId uint, content string, embedUrl string) (CommentResponse, *WykopError) {
	urlAddress := getAddCommentUrl(linkId, parentCommentId, wh)

	body := url.Values{}
	body.Set("body", content)
	urlToHash := urlAddress + content

	if embedUrl != "" {
		body.Set("embed", embedUrl)
		urlToHash += ","+embedUrl
	}

	_, responseBody, _ := gorequest.New().Post(urlAddress).
									Set(contentType, mediaTypeFormType).
									Set(apiSignHeader, wh.hashRequest(urlToHash)).
									Send(body).
									End()

	return wh.getResponseBodyAsCommentResponse(responseBody)
}

func (wh *WykopHandler) PlusComment(linkId uint, commentId uint) (vote Vote, wykopError *WykopError) {
	urlAddress := getPlusLinkCommentUrl(linkId, commentId, wh)

	vote = Vote{FAKE_VOTE_VALUE, FAKE_VOTE_VALUE, FAKE_VOTE_VALUE}

	_, responseBody, _ := gorequest.New().Post(urlAddress).
									Set(contentType, mediaTypeFormType).
									Set(apiSignHeader, wh.hashRequest(urlAddress)).
									End()

	wykopError = wh.getObjectFromJson(responseBody, &vote)

	return
}

func (wh *WykopHandler) MinusComment(linkId uint, commentId uint) (vote Vote, wykopError *WykopError) {
	urlAddress := getMinusLinkCommentUrl(linkId, commentId, wh)

	vote = Vote{FAKE_VOTE_VALUE, FAKE_VOTE_VALUE, FAKE_VOTE_VALUE}

	_, responseBody, _ := gorequest.New().Post(urlAddress).
									Set(contentType, mediaTypeFormType).
									Set(apiSignHeader, wh.hashRequest(urlAddress)).
									End()

	wykopError = wh.getObjectFromJson(responseBody, &vote)

	return
}

func (wh *WykopHandler) EditComment(commentId uint, newContent string) (commentResponse CommentResponse, wykopError *WykopError) {
	urlAddress := getEditLinkCommentUrl(commentId, wh)

	body := url.Values{}
	body.Set("body", newContent)
	urlToHash := urlAddress + newContent

	_, responseBody, _ := gorequest.New().Post(urlAddress).
									Set(contentType, mediaTypeFormType).
									Set(apiSignHeader, wh.hashRequest(urlToHash)).
									Send(body).
									End()

	return wh.getResponseBodyAsCommentResponse(responseBody)
}

func (wh *WykopHandler) DeleteComment(commentId uint) (commentResponse CommentResponse, wykopError *WykopError) {
	urlAddress := getDeleteLinkCommentUrl(commentId, wh)

	_, responseBody, _ := gorequest.New().Post(urlAddress).
									Set(contentType, mediaTypeFormType).
									Set(apiSignHeader, wh.hashRequest(urlAddress)).
									End()

	return wh.getResponseBodyAsCommentResponse(responseBody)
}

func getAddCommentUrl(linkId uint, parentCommentId uint, wh *WykopHandler) string {
	methodWithParams := "add/" + fmt.Sprint(linkId)
	if parentCommentId > 0 {
		methodWithParams += "/" + fmt.Sprint(parentCommentId)
	}
	return fmt.Sprintf(COMMENTS_BASE_URL, methodWithParams, wh.appKey, wh.authResponse.Userkey)
}

func getPlusLinkCommentUrl(linkId uint, commentId uint, wh *WykopHandler) string {
	return fmt.Sprintf(COMMENTS_BASE_URL, "plus/"+fmt.Sprint(linkId)+"/"+fmt.Sprint(commentId), wh.appKey, wh.authResponse.Userkey)
}

func getMinusLinkCommentUrl(linkId uint, commentId uint, wh *WykopHandler) string {
	return fmt.Sprintf(COMMENTS_BASE_URL, "minus/"+fmt.Sprint(linkId)+"/"+fmt.Sprint(commentId), wh.appKey, wh.authResponse.Userkey)
}

func getEditLinkCommentUrl(commentId uint, wh *WykopHandler) string {
	return fmt.Sprintf(COMMENTS_BASE_URL, "edit/" + fmt.Sprint(commentId), wh.appKey, wh.authResponse.Userkey)
}

func getDeleteLinkCommentUrl(commentId uint, wh *WykopHandler) string {
	return fmt.Sprintf(COMMENTS_BASE_URL, "delete/" + fmt.Sprint(commentId), wh.appKey, wh.authResponse.Userkey)
}