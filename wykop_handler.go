package go_wypok

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"net/url"
	"strconv"
)

const (
	contentType        = "Content-Type"
	mediaTypeFormType  = "application/x-www-form-urlencoded"
	apiSignHeader      = "apisign"
	userKeyPathElement = "/userkey/"
	appKeyPathElement  = "/appkey/"
	accountKeyHeader   = "accountkey"
)

type WykopHandler struct {
	appKey        string
	authResponse  AuthenticationResponse
	connectionKey string
	secret        string
}

func (wh *WykopHandler) PostEntryWithImage(content string, absolutePath string) (entryResponse EntryResponse, wypokError *WykopError) {
	body := url.Values{}
	body.Set("body", content)

	urlAddress := getAddEntryUrl() + appKeyPathElement + wh.appKey + userKeyPathElement + wh.authResponse.Userkey

	reqBody := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress+body.Get("body")))
	b, _ := ioutil.ReadFile(absolutePath)

	_, responseBody, _ := reqBody.Send(body).SendFile(b, "", "file").End()

	entryResponse = EntryResponse{}
	wypokError = wh.getObjectFromJson(responseBody, &entryResponse)

	return
}

func (wh *WykopHandler) LoginToWypok() *WykopError {

	responseBody := wh.sendPostRequestForBody(getLoginUrl(wh.appKey))

	wh.authResponse = AuthenticationResponse{}

	return wh.getObjectFromJson(responseBody, &wh.authResponse)
}

func (wh *WykopHandler) GetMainPageLinks(page int) (mainPageLinks []Link, wypokError *WykopError) {
	urlAddress := getMainPageUrl() + appKeyPathElement + wh.appKey

	if wh.authResponse.Userkey != "" {
		urlAddress = urlAddress + userKeyPathElement + wh.authResponse.Userkey
	}

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &mainPageLinks)

	return
}

func (wh *WykopHandler) GetUpcomingLinks(page int) (mainPageLinks []Link, wypokError *WykopError) {
	urlAddress := getUpcomingPageUrl() + appKeyPathElement + wh.appKey

	if wh.authResponse.Userkey != "" {
		urlAddress = urlAddress + userKeyPathElement + wh.authResponse.Userkey
	}

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &mainPageLinks)

	return
}

func (wh *WykopHandler) UpvoteEntry(entry Entry) (voteResponse VoteResponse, wypokError *WykopError) {
	urlAddress := getEntryVoteUrl("entry", strconv.Itoa(entry.Id), "") + appKeyPathElement + wh.appKey + userKeyPathElement + wh.authResponse.Userkey

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &voteResponse)

	return
}

func (wh *WykopHandler) GetEntriesFromTag(tag string, page int) (tagEntries TagsEntries, wypokError *WykopError) {
	urlAddress := getTagEntries(tag) + appKeyPathElement + wh.appKey + "/page/" + strconv.Itoa(page)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &tagEntries)

	return
}

func (wh *WykopHandler) GetProfileEntries(username string, page int) (entries []Entry, wypokError *WykopError) {
	urlAddress := getProfileEntriesUrl(username) + appKeyPathElement + wh.appKey + userKeyPathElement + wh.authResponse.Userkey + "/page/" + strconv.Itoa(page)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &entries)

	return
}

func (wh *WykopHandler) GetProfileComments(username string, page int) (entries []LinkComment, wypokError *WykopError) {
	urlAddress := getProfileCommentsUrl(username) + appKeyPathElement + wh.appKey + userKeyPathElement + wh.authResponse.Userkey + "/page/" + strconv.Itoa(page)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &entries)

	return
}

func (wh *WykopHandler) GetProfileEntriesComments(username string, page int) (entryComments []EntryComment, wypokError *WykopError) {
	urlAddress := getProfileEntriesCommentsUrl(username) + appKeyPathElement + wh.appKey + userKeyPathElement + wh.authResponse.Userkey + "/page/" + strconv.Itoa(page)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &entryComments)

	return
}

func (wh *WykopHandler) DeleteEntryComment(entryId int, commentId int) (commentResponse CommentResponse, wypokError *WykopError) {
	urlAddress := getDeleteCommentUrl(strconv.Itoa(entryId), strconv.Itoa(commentId)) + appKeyPathElement + wh.appKey + userKeyPathElement + wh.authResponse.Userkey

	responseBody := wh.sendPostRequestForBody(urlAddress)

	wypokError = wh.getObjectFromJson(responseBody, &commentResponse)

	return
}

func (wh *WykopHandler) DeleteEntry(id int) (entryResponse EntryResponse, wypokError *WykopError) {
	urlAddress := getDeleteEntryUrl(strconv.Itoa(id)) + appKeyPathElement + wh.appKey + userKeyPathElement + wh.authResponse.Userkey

	responseBody := wh.sendPostRequestForBody(urlAddress)

	wypokError = wh.getObjectFromJson(responseBody, &entryResponse)

	return
}

func (wh *WykopHandler) GetEntry(id int) (entry Entry, wypokError *WykopError) {
	responseBody := wh.sendPostRequestForBody(getEntryUrl(strconv.Itoa(id)) + appKeyPathElement + wh.appKey)

	wypokError = wh.getObjectFromJson(responseBody, &entry)
	return
}

func (wh *WykopHandler) PostEntry(content string) (entryResponse EntryResponse, wypokError *WykopError) {
	body := url.Values{}
	body.Set("body", content)

	urlAddress := getAddEntryUrl() + appKeyPathElement + wh.appKey + userKeyPathElement + wh.authResponse.Userkey

	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress+body.Get("body"))).
		Send(body).
		End()

	entryResponse = EntryResponse{}
	wypokError = wh.getObjectFromJson(responseBody, &entryResponse)

	return
}

func (wh *WykopHandler) PostEntryWithEmbeddedContent(content *string, embeddedUrl *string) (entryResponse EntryResponse, wykopError *WykopError) {
	body := url.Values{}
	body.Set("body", *content)
	body.Set("embed", *embeddedUrl)

	urlAddress := getAddEntryUrl() + appKeyPathElement + wh.appKey + userKeyPathElement + wh.authResponse.Userkey

	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress+body.Get("body")+","+body.Get("embed"))).
		Send(body).
		End()

	wykopError = wh.getObjectFromJson(responseBody, &entryResponse)

	return
}

func (wh *WykopHandler) GetProfile(username string) (profile Profile, wypokError *WykopError) {
	urlAddress := getProfileUrl(username) + appKeyPathElement + wh.appKey

	_, responseBody, _ := gorequest.New().Get(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress)).
		End()

	wypokError = wh.getObjectFromJson(responseBody, &profile)

	return
}

func (wh *WykopHandler) sendPostRequestForBody(address string) string {
	body := url.Values{}
	body.Add(accountKeyHeader, wh.connectionKey)

	_, bodyResp, _ := gorequest.New().Post(address).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(address+wh.connectionKey)).
		Send(body).
		End()

	return bodyResp
}

func (wh *WykopHandler) preparePostRequest(address string) *gorequest.SuperAgent {
	body := url.Values{}
	body.Add(accountKeyHeader, wh.connectionKey)

	return gorequest.New().Post(address).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(address+wh.connectionKey)).
		Send(body)
}

func (wh *WykopHandler) sendGetRequestForBody(address string) string {
	body := url.Values{}
	body.Add(accountKeyHeader, wh.connectionKey)

	_, bodyResp, _ := gorequest.New().Get(address).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(address+wh.connectionKey)).
		Send(body).
		End()

	return bodyResp
}

func (wh *WykopHandler) getObjectFromJson(bodyResponse string, target interface{}) (wypokError *WykopError) {
	b := []byte(bodyResponse)
	if err := json.Unmarshal(b, &wypokError); err != nil {
		// failed to unmarshall response to WypokError, this is actually good
		// wypok-api returned non-error response
	}
	if wypokError.ErrorObject.Message != "" {
		return wypokError
	}
	json.Unmarshal(b, target)
	return nil
}

func (wh *WykopHandler) hashRequest(address string) string {
	toHash := wh.secret + address
	mString := []byte(toHash)
	hash := md5.Sum([]byte(mString))
	return hex.EncodeToString(hash[:])
}

func (w *WykopHandler) SetAppKey(appKey string) {
	w.appKey = appKey
}

func (w *WykopHandler) SetSecret(secret string) {
	w.secret = secret
}

func (w *WykopHandler) SetConnectionKey(connectionKey string) {
	w.connectionKey = connectionKey
}
