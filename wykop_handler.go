package go_wypok

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/url"
	"strconv"
	"strings"
)

type WykopHandler struct {
	appKey        string
	authResponse  AuthenticationResponse
	connectionKey string
	secret        string
}

var Handler *WykopHandler

func Init(handler *WykopHandler) {
	Handler = handler

	Handler.LoginToWypok()
}

func (wh *WykopHandler) LoginToWypok() {

	responseBody := wh.sendPostRequestForBody(getLoginUrl(wh.appKey))

	wh.authResponse = AuthenticationResponse{}

	wypokError := wh.getObjectFromJson(responseBody, &wh.authResponse)
	if wypokError != nil {
		panic(wypokError.ErrorObject.Message)
	}
}

func (wh *WykopHandler) GetMainPageLinks(page int) (mainPageLinks []Link, wypokError *WykopError) {
	urlAddress := getMainPageUrl() + "/appkey/" + wh.appKey

	if wh.authResponse.Userkey != "" {
		urlAddress = urlAddress + "/userkey/" + wh.authResponse.Userkey
	}

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	mainPageLinks = []Link{}
	wypokError = wh.getObjectFromJson(responseBody, &mainPageLinks)

	return
}

func (wh *WykopHandler) GetUpcomingLinks(page int) (mainPageLinks []Link, wypokError *WykopError) {
	urlAddress := getUpcomingPageUrl() + "/appkey/" + wh.appKey

	if wh.authResponse.Userkey != "" {
		urlAddress = urlAddress + "/userkey/" + wh.authResponse.Userkey
	}

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	mainPageLinks = []Link{}
	wypokError = wh.getObjectFromJson(responseBody, &mainPageLinks)

	return
}

func (wh *WykopHandler) UpvoteEntry(entry Entry) (voteResponse VoteResponse, wypokError *WykopError) {
	urlAddress := getEntryVoteUrl("entry", strconv.Itoa(entry.Id), "") + "/appkey/" + wh.appKey + "/userkey/" + wh.authResponse.Userkey

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	voteResponse = VoteResponse{}
	wypokError = wh.getObjectFromJson(responseBody, &voteResponse)

	return
}

func (wh *WykopHandler) GetEntriesFromTag(tag string, page int) (tagEntries TagsEntries, wypokError *WykopError) {
	urlAddress := getTagEntries(tag) + "/appkey/" + wh.appKey + "/page/" + strconv.Itoa(page)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	entries := TagsEntries{}
	wypokError = wh.getObjectFromJson(responseBody, &entries)

	return
}

func (wh *WykopHandler) sendPostRequestForBody(address string) string {
	body := url.Values{}
	body.Add("accountkey", wh.connectionKey)

	_, bodyResp, _ := gorequest.New().Post(address).
		Set("Content-Type", "application/x-www-form-urlencoded").
		Set("apisign", wh.hashRequest(address+wh.connectionKey)).
		Send(body).
		End()

	return bodyResp
}

func (wh *WykopHandler) preparePostRequest(address string) *gorequest.SuperAgent {
	body := url.Values{}
	body.Add("accountkey", wh.connectionKey)

	return gorequest.New().Post(address).
		Set("Content-Type", "application/x-www-form-urlencoded").
		Set("apisign", wh.hashRequest(address+wh.connectionKey)).
		Send(body)
}

func (wh *WykopHandler) sendGetRequestForBody(address string) string {
	body := url.Values{}
	body.Add("accountkey", wh.connectionKey)

	_, bodyResp, _ := gorequest.New().Get(address).
		Set("Content-Type", "application/x-www-form-urlencoded").
		Set("apisign", wh.hashRequest(address+wh.connectionKey)).
		Send(body).
		End()

	return bodyResp
}

func (wh *WykopHandler) GetProfileEntries(username string, page int) (entries []Entry, wypokError *WykopError) {
	urlAddress := getProfileEntriesUrl(username) + "/appkey/" + wh.appKey + "/userkey/" + wh.authResponse.Userkey + "/page/" + strconv.Itoa(page)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).
		End()

	entries = []Entry{}
	wypokError = wh.getObjectFromJson(responseBody, &entries)

	return
}

func (wh *WykopHandler) GetEntriesComments(username string, page int) (comments []EntryComment, wypokError *WykopError) {
	urlAddress := getEntriesCommentsUrl(username) + "/appkey/" + wh.appKey + "/userkey/" + wh.authResponse.Userkey + "/page/" + strconv.Itoa(page)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	comments = []EntryComment{}
	wypokError = wh.getObjectFromJson(responseBody, &comments)

	return
}

func (wh *WykopHandler) DeleteEntryComment(entryId string, commentId string) (wypokError *WykopError) {
	urlAddress := getDeleteCommentUrl(entryId, commentId) + "/appkey/" + wh.appKey + "/userkey/" + wh.authResponse.Userkey

	responseBody := wh.sendPostRequestForBody(urlAddress)

	commentResponse := CommentResponse{}
	wypokError = wh.getObjectFromJson(responseBody, &commentResponse)

	return
}

func (wh *WykopHandler) DeleteEntry(id string) (entryResponse EntryResponse, wypokError *WykopError) {
	urlAddress := getDeleteEntryUrl(id) + "/appkey/" + wh.appKey + "/userkey/" + wh.authResponse.Userkey

	responseBody := wh.sendPostRequestForBody(urlAddress)

	entryResponse = EntryResponse{}
	wypokError = wh.getObjectFromJson(responseBody, &entryResponse)

	return
}

func (wh *WykopHandler) GetEntry(id int) (entry Entry, wypokError *WykopError) {
	responseBody := wh.sendPostRequestForBody(getEntryUrl(strconv.Itoa(id)))

	entry = Entry{}
	wypokError = wh.getObjectFromJson(responseBody, &entry)

	return
}

func (wh *WykopHandler) PostEntry(content *string) (entryResponse EntryResponse, wypokError *WykopError) {
	body := url.Values{}
	body.Set("body", *content)

	urlAddress := getAddEntryUrl() + "/appkey/" + wh.appKey + "/userkey/" + wh.authResponse.Userkey

	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set("Content-Type", "application/x-www-form-urlencoded").
		Set("apisign", wh.hashRequest(urlAddress+body.Get("body"))).
		Send(body).
		End()

	entryResponse = EntryResponse{}
	wypokError = wh.getObjectFromJson(responseBody, &entryResponse)

	return
}

func (wh *WykopHandler) GetProfile(username string) (profile Profile, wypokError *WykopError) {
	urlAddress := getProfileUrl(username) + "/appkey/" + wh.appKey

	_, responseBody, _ := gorequest.New().Get(urlAddress).
		Set("Content-Type", "application/x-www-form-urlencoded").
		Set("apisign", wh.hashRequest(urlAddress)).
		End()

	profile = Profile{}
	wypokError = wh.getObjectFromJson(responseBody, &profile)

	return
}

func (wh *WykopHandler) getObjectFromJson(bodyReader string, target interface{}) *WykopError {
	parsingError := json.NewDecoder(strings.NewReader(bodyReader)).Decode(target)
	if parsingError != nil {
		fmt.Println(parsingError.Error())
		var err *WykopError = &WykopError{}
		parsingError = json.NewDecoder(strings.NewReader(bodyReader)).Decode(err)
		if parsingError != nil {
			panic(parsingError.Error())
		}
		return err
	}
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
