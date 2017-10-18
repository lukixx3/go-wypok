package go_wypok

import (
	"fmt"
)

type Profile struct {
	Id              uint
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
	urlAddress := getProfileUrl(username, wh)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &profile)

	return
}

func (wh *WykopHandler) GetProfileAdded(username string, page uint) (profile []Link, wypokError *WykopError) {
	urlAddress := getProfileAddedUrl(username, wh, page)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &profile)

	return
}

func (wh *WykopHandler) GetProfilePublished(username string, page uint) (profile []Link, wypokError *WykopError) {
	urlAddress := getProfilePublishedUrl(username, wh, page)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &profile)

	return
}

func (wh *WykopHandler) GetProfileCommented(username string, page uint) (profile []Link, wypokError *WykopError) {
	urlAddress := getProfileCommentedUrl(username, wh, page)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &profile)

	return
}

func (wh *WykopHandler) GetProfileComments(username string, page uint) (entries []LinkComment, wypokError *WykopError) {
	urlAddress := getProfileCommentsUrl(username, wh, page)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &entries)

	return
}

func (wh *WykopHandler) GetProfileDigged(username string, page uint) (entries []Link, wypokError *WykopError) {
	urlAddress := getProfileDiggedUrl(username, wh, page)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &entries)

	return
}

func (wh *WykopHandler) GetProfileBuried(username string, page uint) (entries []Link, wypokError *WykopError) {
	urlAddress := getProfileBuriedUrl(username, wh, page)
	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &entries)

	return
}

func (wh *WykopHandler) GetProfileFavorites(username string) (links []Link, wypokError *WykopError) {
	urlAddress := getProfileFavoritesUrl(username, wh)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &links)

	return
}

func (wh *WykopHandler) GetProfileEntries(username string, page uint) (entries []Entry, wypokError *WykopError) {
	urlAddress := getProfileEntriesUrl(username, wh, page)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &entries)

	return
}

func (wh *WykopHandler) GetProfileEntriesComments(username string, page uint) (entryComments []EntryComment, wypokError *WykopError) {
	urlAddress := getProfileEntriesCommentsUrl(username, wh, page)

	responseBody := wh.sendGetRequest(urlAddress)
	wypokError = wh.getObjectFromJson(responseBody, &entryComments)

	return
}

func getProfileUrl(username string, wh *WykopHandler) string {
	return fmt.Sprintf(PROFILE_INDEX, username, wh.appKey)
}

func getProfileAddedUrl(username string, wh *WykopHandler, page uint) string {
	return fmt.Sprintf(PROFILE_ADDED, username, wh.appKey, page)
}

func getProfilePublishedUrl(username string, wh *WykopHandler, page uint) string {
	return fmt.Sprintf(PROFILE_PUBLISHED, username, wh.appKey, page)
}

func getProfileCommentedUrl(username string, wh *WykopHandler, page uint) string {
	return fmt.Sprintf(PROFILE_COMMENTED, username, wh.appKey, page)
}

func getProfileCommentsUrl(username string, wh *WykopHandler, page uint) string {
	return fmt.Sprintf(PROFILE_COMMENTS, username, wh.appKey, wh.authResponse.Userkey, page)
}

func getProfileDiggedUrl(username string, wh *WykopHandler, page uint) string {
	return fmt.Sprintf(PROFILE_DIGGED, username, wh.appKey, page)
}

func getProfileBuriedUrl(username string, wh *WykopHandler, page uint) string {
	return fmt.Sprintf(PROFILE_BURIED, username, wh.appKey, wh.authResponse.Userkey, page)
}

func getProfileFavoritesUrl(username string, wh *WykopHandler) string {
	return fmt.Sprintf(PROFILE_FAVORITES, username, wh.appKey, wh.authResponse.Userkey)
}

func getProfileEntriesUrl(username string, wh *WykopHandler, page uint) string {
	return fmt.Sprintf(PROFILE_ENTRIES, username, wh.appKey, wh.authResponse.Userkey, page)
}

func getProfileEntriesCommentsUrl(username string, wh *WykopHandler, page uint) string {
	return fmt.Sprintf(PROFILE_ENTRY_COMMENTS, username, wh.appKey, wh.authResponse.Userkey, page)
}
