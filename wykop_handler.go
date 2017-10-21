package go_wypok

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
	"net/url"
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

type AuthenticationResponse struct {
	Login        string
	Email        string
	ViolationUrl string `json:"violation_url"`
	Userkey      string
}

func (wh *WykopHandler) LoginToWypok() *WykopError {

	responseBody := wh.sendPostRequestForBody(getLoginUrl(wh))

	wh.authResponse = AuthenticationResponse{}

	return wh.getObjectFromJson(responseBody, &wh.authResponse)
}

func (wh *WykopHandler) GetEntriesFromTag(tag string, page uint) (tagEntries TagsEntries, wypokError *WykopError) {
	urlAddress := getTagEntries(tag, wh, page)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &tagEntries)

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

func (wh *WykopHandler) sendGetRequest(address string) string {
	body := url.Values{}
	body.Add(accountKeyHeader, wh.connectionKey)

	_, bodyResp, _ := gorequest.New().Post(address).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(address+wh.connectionKey)).
		Send(body).
		End()
	return bodyResp
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
		// this might happen when wypok is being ddosed/updated or Kiner is making a party in the server room
		// this might happen when a.wykop.pl returned html, or empty response, happens.
		wypokError = new(WykopError)
		wypokError.ErrorObject.Message = "Coś poszło nie tak, wykop api nie zwróciło ani błędu ani obiektu"
		wypokError.ErrorObject.Code = -1
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
