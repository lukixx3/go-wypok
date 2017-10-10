package go_wypok

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type RankSorting int

const (
	rank = 1 + iota
	comment_count
	link_count
	hp_link_count
	followers_count
)

var sortingTypes = [...]string{
	"rank",
	"comment_count",
	"link_count",
	"hp_link_count",
	"followers_count",
}

// return sorting type string based on int value of const
func (m RankSorting) String() string {
	return sortingTypes[m-1]
}

// Rank returns list of Profile struct

func (wh *WykopHandler) GetRank() (profiles []Profile, wypokError *WykopError) {
	return wh.GetRankBySortingType(rank)
}

func (wh *WykopHandler) GetRankBySortingType(sorting RankSorting) (profiles []Profile, wypokError *WykopError) {
	urlAddress := getRankUrl(wh, sorting)

	_, responseBody, _ := gorequest.New().Get(urlAddress).
		Set(contentType, mediaTypeFormType).
		Set(apiSignHeader, wh.hashRequest(urlAddress)).
		End()

	wypokError = wh.getObjectFromJson(responseBody, &profiles)
	return
}

func getRankUrl(wh *WykopHandler, sorting RankSorting) string {
	return fmt.Sprintf(RANK_INDEX, wh.appKey, sorting.String())
}
