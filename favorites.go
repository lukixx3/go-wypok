package go_wypok

import "fmt"

// GetFavorites returns user favorites links.
// It requires privileges for profile.
func (wh *WykopHandler) GetFavorites(id string) (links []Link, wypokError *WykopError) {
	responseBody := wh.sendPostRequestForBody(
		getFavoritesURL(id, wh.appKey, wh.authResponse.Userkey),
	)

	wypokError = wh.getObjectFromJson(responseBody, &links)
	return
}

func getFavoritesURL(listID, appKey, userkey string) string {
	return fmt.Sprintf(FAVORITES_INDEX, listID, appKey, userkey)
}
