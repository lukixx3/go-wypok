package go_wypok

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
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

func (wh *WykopHandler) UpvoteEntries(entries TagsEntries) {
	for i := range entries.Items {
		lastEntry := entries.Items[len(entries.Items)-1-i]

		if lastEntry.Vote_count >= 10 {
			go wh.UpvoteEntry(lastEntry)
		}
	}
}

func (wh *WykopHandler) UpvoteEntry(entry Entry) VoteResponse {
	urlAddress := getEntryVoteUrl("entry", strconv.Itoa(entry.Id), "") + "/appkey/" + wh.appKey + "/userkey/" + wh.authResponse.Userkey

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	voteResponse := VoteResponse{}

	wypokError := wh.getObjectFromJson(responseBody, &voteResponse)
	if wypokError != nil {
		panic(wypokError.ErrorObject.Message)
	}

	return voteResponse
}

func (wh *WykopHandler) GetEntriesFromTag(tag string, page int) TagsEntries {
	urlAddress := getTagEntries(tag) + "/appkey/" + wh.appKey + "/page/" + strconv.Itoa(page)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	entries := TagsEntries{}

	wypokError := wh.getObjectFromJson(responseBody, &entries)
	if wypokError != nil {
		panic(wypokError.ErrorObject.Message)
	}

	return entries
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

func (wh *WykopHandler) GetProfileEntries(username string, page int) []Entry {
	urlAddress := getProfileEntriesUrl(username) + "/appkey/" + wh.appKey + "/userkey/" + wh.authResponse.Userkey + "/page/" + strconv.Itoa(page)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).
		End()

	entries := []Entry{}

	wypokError := wh.getObjectFromJson(responseBody, &entries)
	if wypokError != nil {
		panic(wypokError.ErrorObject.Message)
	}

	return entries
}

func (wh *WykopHandler) GetEntriesComments(username string, page int) []EntryComment {
	urlAddress := getEntriesCommentsUrl(username) + "/appkey/" + wh.appKey + "/userkey/" + wh.authResponse.Userkey + "/page/" + strconv.Itoa(page)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	comments := []EntryComment{}

	wypokError := wh.getObjectFromJson(responseBody, &comments)
	if wypokError != nil {
		panic(wypokError.ErrorObject.Message)
	}

	return comments
}

func (wh *WykopHandler) DeleteEntryComment(entryId string, commentId string) {
	urlAddress := getDeleteCommentUrl(entryId, commentId) + "/appkey/" + wh.appKey + "/userkey/" + wh.authResponse.Userkey

	responseBody := wh.sendPostRequestForBody(urlAddress)

	commentResponse := CommentResponse{}
	wypokError := wh.getObjectFromJson(responseBody, &commentResponse)
	if wypokError != nil {
		panic(wypokError.ErrorObject.Message)
	}
}

func (wh *WykopHandler) DeleteEntry(id string) EntryResponse {
	urlAddress := getDeleteEntryUrl(id) + "/appkey/" + wh.appKey + "/userkey/" + wh.authResponse.Userkey

	responseBody := wh.sendPostRequestForBody(urlAddress)

	entry := EntryResponse{}

	wypokError := wh.getObjectFromJson(responseBody, &entry)
	if wypokError != nil {
		panic(wypokError.ErrorObject.Message)
	}

	return entry
}

func (wh *WykopHandler) GetEntry(id int) Entry {
	responseBody := wh.sendPostRequestForBody(getEntryUrl(strconv.Itoa(id)))

	entry := Entry{}

	wypokError := wh.getObjectFromJson(responseBody, &entry)
	if wypokError != nil {
		panic(wypokError.ErrorObject.Message)
	}

	return entry
}

func (wh *WykopHandler) PostEntry(content *string) EntryResponse {
	body := url.Values{}
	body.Set("body", *content)

	urlAddress := getAddEntryUrl() + "/appkey/" + wh.appKey + "/userkey/" + wh.authResponse.Userkey

	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set("Content-Type", "application/x-www-form-urlencoded").
		Set("apisign", wh.hashRequest(urlAddress+body.Get("body"))).
		Send(body).
		End()

	entry := EntryResponse{}

	wypokError := wh.getObjectFromJson(responseBody, &entry)
	if wypokError != nil {
		panic(wypokError.ErrorObject.Message)
	}

	return entry
}

func (wh *WykopHandler) GetProfile(username string) Profile {
	urlAddress := getProfileUrl(username) + "/appkey/" + wh.appKey

	_, responseBody, _ := gorequest.New().Get(urlAddress).
		Set("Content-Type", "application/x-www-form-urlencoded").
		Set("apisign", wh.hashRequest(urlAddress)).
		End()

	profile := Profile{}

	wypokError := wh.getObjectFromJson(responseBody, &profile)
	if wypokError != nil {
		panic(wypokError.ErrorObject.Message)
	}

	return profile
}

func (wh *WykopHandler) getObjectFromJson(bodyReader string, target interface{}) *WykopError {
	parsinError := json.NewDecoder(strings.NewReader(bodyReader)).Decode(target)
	if parsinError != nil {
		var err *WykopError = &WykopError{}
		parsinError = json.NewDecoder(strings.NewReader(bodyReader)).Decode(err)
		if parsinError != nil {
			panic(parsinError.Error())
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
