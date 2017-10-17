package go_wypok

import "fmt"

// FavoritesList contains properties of single favorites list
type FavoritesList struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Public bool   `json:"is_public"`
}

// GetFavoritesLists returns user favorites lists
// It requires privileges for profile.
func (wh *WykopHandler) GetFavoritesLists() (lists []FavoritesList, wypokError *WykopError) {
	responseBody := wh.sendPostRequestForBody(
		getFavoritesListsURL(wh.appKey, wh.authResponse.Userkey),
	)

	wypokError = wh.getObjectFromJson(responseBody, &lists)
	return
}

// GetFavoritesListLinks returns user favorites list links.
// It requires privileges for profile.
func (wh *WykopHandler) GetFavoritesListLinks(id uint) (links []Link, wypokError *WykopError) {
	responseBody := wh.sendPostRequestForBody(
		getFavoritesListLinksURL(id, wh.appKey, wh.authResponse.Userkey),
	)

	wypokError = wh.getObjectFromJson(responseBody, &links)
	return
}

// GetFavoritesComments returns user favorites links comments.
// It requires privileges for profile.
func (wh *WykopHandler) GetFavoritesComments() (comments []LinkComment, wypokError *WykopError) {
	responseBody := wh.sendPostRequestForBody(
		getFavoritesCommentsURL(wh.appKey, wh.authResponse.Userkey),
	)

	wypokError = wh.getObjectFromJson(responseBody, &comments)
	return
}

// GetFavoritesEntries returns user favorites entries.
// It requires privileges for profile.
func (wh *WykopHandler) GetFavoritesEntries() (entries []Entry, wypokError *WykopError) {
	responseBody := wh.sendPostRequestForBody(
		getFavoritesEntriesURL(wh.appKey, wh.authResponse.Userkey),
	)

	wypokError = wh.getObjectFromJson(responseBody, &entries)
	return
}

func getFavoritesListsURL(appKey, userkey string) string {
	return fmt.Sprintf(FAVORITES_LISTS, appKey, userkey)
}

func getFavoritesListLinksURL(listID uint, appKey, userkey string) string {
	return fmt.Sprintf(FAVORITES_INDEX, listID, appKey, userkey)
}

func getFavoritesCommentsURL(appKey, userkey string) string {
	return fmt.Sprintf(FAVORITES_COMMENTS, appKey, userkey)
}

func getFavoritesEntriesURL(appKey, userkey string) string {
	return fmt.Sprintf(FAVORITES_ENTRIES, appKey, userkey)
}
