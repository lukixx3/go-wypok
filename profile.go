package go_wypok

import (
	"fmt"
)

type Profile struct {
	Id              int
	Login           string
	Email           string
	PublicEmail     string `json:"public_email"`
	Name            string
	Www             string
	Jabber          string
	Gg              int
	City            string
	About           string
	AuthorGroup     int `json:"author_group"`
	LinksAdded      int `json:"links_added"`
	LinksPublished  int `json:"links_published"`
	Comments        int
	Rank            int
	Followers       int
	Following       int
	Entries         int
	EntriesComments int `json:"entries_comments"`
	Diggs           int
	Buries          int
	RelatedLinks    int `json:"related_links"`
	Groups          int
	Sex             string
	Avatar          string
	AvatarLo        string `json:"avatar_lo"`
	AvatarMed       string `json:"avatar_med"`
	AvatarBig       string `json:"avatar_big"`
	IsObserved      bool   `json:"is_observed"`
}

func (wh *WykopHandler) GetProfile(username string) (profile Profile, wypokError *WykopError) {
	urlAddress := getProfileUrl(username, wh.appKey)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &profile)

	return
}

func (wh *WykopHandler) GetProfileAdded(username string, page uint) (profile []Link, wypokError *WykopError) {
	urlAddress := getProfileAddedUrl(username, wh.appKey, page)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &profile)

	return
}

func (wh *WykopHandler) GetProfilePublished(username string, page uint) (profile []Link, wypokError *WykopError) {
	urlAddress := getProfilePublishedUrl(username, wh.appKey, page)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &profile)

	return
}

func (wh *WykopHandler) GetProfileCommented(username string, page uint) (profile []Link, wypokError *WykopError) {
	urlAddress := getProfileCommentedUrl(username, wh.appKey, page)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &profile)

	return
}

func (wh *WykopHandler) GetProfileComments(username string, page uint) (entries []LinkComment, wypokError *WykopError) {
	urlAddress := getProfileCommentsUrl(username, wh.appKey, wh.authResponse.Userkey, page)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &entries)

	return
}

func (wh *WykopHandler) GetProfileDigged(username string, page uint) (entries []Link, wypokError *WykopError) {
	urlAddress := getProfileDiggedUrl(username, wh.appKey, page)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &entries)

	return
}

func (wh *WykopHandler) GetProfileBuried(username string, page uint) (entries []Link, wypokError *WykopError) {
	urlAddress := getProfileBuriedUrl(username, wh.appKey, wh.authResponse.Userkey, page)
	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &entries)

	return
}

func (wh *WykopHandler) GetProfileFavorites(username string) (links []Link, wypokError *WykopError) {
	urlAddress := getProfileFavoritesUrl(username, wh.appKey, wh.authResponse.Userkey)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &links)

	return
}

func (wh *WykopHandler) GetProfileEntries(username string, page uint) (entries []Entry, wypokError *WykopError) {
	urlAddress := getProfileEntriesUrl(username, wh.appKey, wh.authResponse.Userkey, page)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &entries)

	return
}

func (wh *WykopHandler) GetProfileEntriesComments(username string, page uint) (entryComments []EntryComment, wypokError *WykopError) {
	urlAddress := getProfileEntriesCommentsUrl(username, wh.appKey, wh.authResponse.Userkey, page)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &entryComments)

	return
}

func getProfileUrl(username, appkey string) string {
	return fmt.Sprintf(PROFILE_INDEX, username, appkey)
}

func getProfileAddedUrl(username, appkey string, page uint) string {
	return fmt.Sprintf(PROFILE_ADDED, username, appkey, page)
}

func getProfilePublishedUrl(username, appkey string, page uint) string {
	return fmt.Sprintf(PROFILE_PUBLISHED, username, appkey, page)
}

func getProfileCommentedUrl(username, appkey string, page uint) string {
	return fmt.Sprintf(PROFILE_COMMENTED, username, appkey, page)
}

func getProfileCommentsUrl(username, appkey, userkey string, page uint) string {
	return fmt.Sprintf(PROFILE_COMMENTS, username, appkey, userkey, page)
}

func getProfileDiggedUrl(username, appkey string, page uint) string {
	return fmt.Sprintf(PROFILE_DIGGED, username, appkey, page)
}

func getProfileBuriedUrl(username, appkey, userkey string, page uint) string {
	return fmt.Sprintf(PROFILE_BURIED, username, appkey, userkey, page)
}

func getProfileFavoritesUrl(username, appkey, userkey string) string {
	return fmt.Sprintf(PROFILE_FAVORITES, username, appkey, userkey)
}

func getProfileEntriesUrl(username, appkey, userkey string, page uint) string {
	return fmt.Sprintf(PROFILE_ENTRIES, username, appkey, userkey, page)
}

func getProfileEntriesCommentsUrl(username, appkey, userkey string, page uint) string {
	return fmt.Sprintf(PROFILE_ENTRY_COMMENTS, username, appkey, userkey, page)
}
