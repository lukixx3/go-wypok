package go_wypok

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

// Rank returns list of Profile struct

func (wh *WykopHandler) GetRank() (profiles []Profile, wypokError *WykopError) {
	urlAddress := getRankUrl(wh)

	_, responseBody, _ := gorequest.New().Get(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress)).
		End()

	wypokError = wh.getObjectFromJson(responseBody, &profiles)
	return
}

func getRankUrl(wh *WykopHandler) string {
	return fmt.Sprintf(RANK_INDEX, wh.appKey)
}
