package go_wypok

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"net/url"
)

type Entry struct {
	Id                int
	Author            string
	AuthorAvatar      string        `json:"author_avatar"`
	AuthorAvatarBig   string        `json:"author_big"`
	AuthorAvatarMed   string        `json:"author_med"`
	AuthorAvatarLo    string        `json:"author_lo"`
	AuthorGroup       int           `json:"author_group"`
	AuthorSex         string        `json:"author_sex"`
	Date              WypokShitDate `json:"date"`
	Body              string
	Source            string
	Url               string
	Receiver          string
	ReceiverAvatar    string `json:"receiver_avatar"`
	ReceiverAvatarBig string `json:"receiver_avatar_big"`
	ReceiverAvatarMed string `json:"receiver_avatar_med"`
	ReceiverAvatarLo  string `json:"receiver_avatar_lo"`
	ReceiverGroup     string `json:"receiver_group"`
	ReceiverSex       string `json:"receiver_sex"`
	Comments          []EntryComment
	Blocked           bool
	VoteCount         int `json:"vote_count"`
	UserVote          int `json:"user_vote"`
	Voters            []Voter
	UserFavorite      bool   `json:"user_favorite"`
	EntryType         string `json:"type"`
	Embed             Embed
	Deleted           bool
	ViolationUrl      string `json:"violation_url"`
	CanComment        bool   `json:"can_comment"`
	App               string
	CommentCount      int `json:"comment_count"`
}

type EntryComment struct {
	ID              int           `json:"id"`
	Author          string        `json:"author"`
	AuthorAvatar    string        `json:"author_avatar"`
	AuthorAvatarBig string        `json:"author_avatar_big"`
	AuthorAvatarMed string        `json:"author_avatar_med"`
	AuthorAvatarLo  string        `json:"author_avatar_lo"`
	AuthorGroup     int           `json:"author_group"`
	AuthorSex       string        `json:"author_sex"`
	Date            WypokShitDate `json:"date"`
	Body            string        `json:"body"`
	Source          string        `json:"source"`
	EntryID         int           `json:"entry_id"`
	Blocked         bool          `json:"blocked"`
	Deleted         bool          `json:"deleted"`
	VoteCount       int           `json:"vote_count"`
	UserVote        int           `json:"user_vote"`
	Voters          []Voter
	Embed           Embed  `json:"embed"`
	Type            string `json:"type"`
	App             string `json:"app"`
	ViolationURL    string `json:"violation_url"`
	Entry           Entry
}

type UpvoteType int

const (
	entry = 1 + iota
	comment
)

var upvoteTypes = [...]string{
	"entry",
	"comment",
}

// return sorting type string based on int value of const
func (m UpvoteType) String() string {
	return upvoteTypes[m-1]
}

func (wh *WykopHandler) GetEntry(id string) (entry Entry, wypokError *WykopError) {
	urlAddress := getEntryUrl(id, wh.appKey)

	responseBody := wh.sendPostRequestForBody(urlAddress)

	wypokError = wh.getObjectFromJson(responseBody, &entry)
	return
}

func (wh *WykopHandler) PostEntry(content string) (entryResponse EntryResponse, wypokError *WykopError) {
	urlAddress := getAddEntryUrl(wh.appKey, wh.authResponse.Userkey)

	body := url.Values{}
	body.Set("body", content)
	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress+body.Get("body"))).
		Send(body).
		End()

	wypokError = wh.getObjectFromJson(responseBody, &entryResponse)

	return
}

func (wh *WykopHandler) PostEntryWithEmbeddedContent(content string, embeddedUrl string) (entryResponse EntryResponse, wykopError *WykopError) {
	urlAddress := getAddEntryUrl(wh.appKey, wh.authResponse.Userkey)

	body := url.Values{}
	body.Set("body", content)
	body.Set("embed", embeddedUrl)
	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress+body.Get("body")+","+body.Get("embed"))).
		Send(body).
		End()

	wykopError = wh.getObjectFromJson(responseBody, &entryResponse)

	return
}

func (wh *WykopHandler) PostEntryWithImage(content string, absolutePath string) (entryResponse EntryResponse, wypokError *WykopError) {
	urlAddress := getAddEntryUrl(wh.appKey, wh.authResponse.Userkey)

	body := url.Values{}
	body.Set("body", content)
	reqBody := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress+body.Get("body")))
	b, _ := ioutil.ReadFile(absolutePath)

	_, responseBody, _ := reqBody.Send(body).SendFile(b, "", "file").End()

	wypokError = wh.getObjectFromJson(responseBody, &entryResponse)

	return
}

func (wh *WykopHandler) EditEntry(entryId string, content string) (entryResponse EntryResponse, wypokError *WykopError) {
	urlAddress := getEditEntryUrl(entryId, wh.appKey, wh.authResponse.Userkey)

	body := url.Values{}
	body.Set("body", content)
	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress+body.Get("body"))).
		Send(body).
		End()

	wypokError = wh.getObjectFromJson(responseBody, &entryResponse)

	return
}

func (wh *WykopHandler) AddEntryComment(entryId string, comment string) (commentResponse CommentResponse, wypokError *WykopError) {
	urlAddress := getEntryAddCommentUrl(entryId, wh.appKey, wh.authResponse.Userkey)

	body := url.Values{}
	body.Set("body", comment)
	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress+body.Get("body"))).
		Send(body).
		End()

	wypokError = wh.getObjectFromJson(responseBody, &commentResponse)

	return
}

func (wh *WykopHandler) AddEntryCommentWithEmbeddedContent(entryId string, comment string, embeddedUrl string) (commentResponse CommentResponse, wypokError *WykopError) {
	urlAddress := getEntryAddCommentUrl(entryId, wh.appKey, wh.authResponse.Userkey)

	body := url.Values{}
	body.Set("body", comment)
	body.Set("embed", embeddedUrl)
	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress+body.Get("body")+","+body.Get("embed"))).
		Send(body).
		End()

	wypokError = wh.getObjectFromJson(responseBody, &commentResponse)

	return
}

func (wh *WykopHandler) EditEntryComment(entryId string, commentId string, comment string) (commentResponse CommentResponse, wypokError *WykopError) {
	urlAddress := getEditEntryCommentUrl(entryId, commentId, wh.appKey, wh.authResponse.Userkey)

	body := url.Values{}
	body.Set("body", comment)
	_, responseBody, _ := gorequest.New().Post(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress+body.Get("body"))).
		Send(body).
		End()

	wypokError = wh.getObjectFromJson(responseBody, &commentResponse)

	return
}

func (wh *WykopHandler) DeleteEntryComment(entryId string, commentId string) (commentResponse CommentResponse, wypokError *WykopError) {
	urlAddress := getDeleteCommentUrl(entryId, commentId, wh.appKey, wh.authResponse.Userkey)

	responseBody := wh.sendPostRequestForBody(urlAddress)

	wypokError = wh.getObjectFromJson(responseBody, &commentResponse)

	return
}

func (wh *WykopHandler) DeleteEntry(id string) (entryResponse EntryResponse, wypokError *WykopError) {
	urlAddress := getDeleteEntryUrl(id, wh.appKey, wh.authResponse.Userkey)

	responseBody := wh.sendPostRequestForBody(urlAddress)

	wypokError = wh.getObjectFromJson(responseBody, &entryResponse)

	return
}

func (wh *WykopHandler) UpvoteEntry(entryId string) (voteResponse VoteResponse, wypokError *WykopError) {
	urlAddress := getEntryVoteUrl(entry, entryId, "", wh.appKey, wh.authResponse.Userkey)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &voteResponse)

	return
}

func (wh *WykopHandler) UpvoteEntryComment(entryId string, commentId string) (voteResponse VoteResponse, wypokError *WykopError) {
	urlAddress := getEntryVoteUrl(entry, entryId, commentId, wh.appKey, wh.authResponse.Userkey)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &voteResponse)

	return
}

func (wh *WykopHandler) UnvoteEntryComment(entryId string, commentId string) (voteResponse VoteResponse, wypokError *WykopError) {
	urlAddress := getEntryUnvoteUrl(entry, entryId, commentId, wh.appKey, wh.authResponse.Userkey)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &voteResponse)

	return
}

func (wh *WykopHandler) UnvoteEntry(entryId string) (voteResponse VoteResponse, wypokError *WykopError) {
	urlAddress := getEntryUnvoteUrl(entry, entryId, "", wh.appKey, wh.authResponse.Userkey)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &voteResponse)

	return
}

func (wh *WykopHandler) FavoriteEntry(entryId string) (favoriteResponse FavoriteResponse, wypokError *WykopError) {
	urlAddress := getEntryFavoriteUrl(entryId, wh.appKey, wh.authResponse.Userkey)

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &favoriteResponse)

	return
}

func getEntryUrl(entry string, appKey string) string {
	return fmt.Sprintf(ENTRY_INDEX, entry, appKey)
}

func getAddEntryUrl(appKey string, userKey string) string {
	return fmt.Sprintf(ENTRY_ADD, appKey, userKey)
}

func getEditEntryUrl(entryId string, appKey string, userKey string) string {
	return fmt.Sprintf(ENTRY_EDIT, entryId, appKey, userKey)
}

func getDeleteEntryUrl(entryId string, appKey string, userKey string) string {
	return fmt.Sprintf(ENTRY_DELETE, entryId, appKey, userKey)
}

func getEntryAddCommentUrl(entryId string, appKey string, userKey string) string {
	return fmt.Sprintf(ENTRY_ADD_COMMENT, entryId, appKey, userKey)
}

func getEditEntryCommentUrl(entryId string, commentId string, appKey string, userKey string) string {
	return fmt.Sprintf(ENTRY_COMMENT_EDIT, entryId, commentId, appKey, userKey)
}

func getDeleteCommentUrl(entryId string, commentId string, appKey string, userKey string) string {
	return fmt.Sprintf(ENTRY_COMMENT_DELETE, entryId, commentId, appKey, userKey)
}

func getEntryVoteUrl(voteType UpvoteType, entryId string, commentId string, appKey string, userKey string) string {
	if voteType == entry {
		return fmt.Sprintf(ENTRY_VOTE, entryId, appKey, userKey)
	} else {
		return fmt.Sprintf(ENTRY_COMMENT_VOTE, entryId, commentId, appKey, userKey)
	}
}

func getEntryUnvoteUrl(voteType UpvoteType, entryId string, commentId string, appKey string, userKey string) string {
	if voteType == entry {
		return fmt.Sprintf(ENTRY_UNVOTE, entryId, appKey, userKey)
	} else {
		return fmt.Sprintf(ENTRY_COMMENT_UNVOTE, entryId, commentId, appKey, userKey)
	}
}

func getEntryFavoriteUrl(entryId string, appKey string, userKey string) string {
	return fmt.Sprintf(ENTRY_FAVORITE, entryId, appKey, userKey)
}
