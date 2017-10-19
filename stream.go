package go_wypok

import (
	"fmt"
)

// GetStreamEntries returns entries from Mikroblog
func (wh *WykopHandler) GetStreamEntries(page uint) (entries []Entry, wypokError *WykopError) {
	responseBody := wh.sendPostRequestForBody(
		getStreamIndexUrl(wh.appKey, page),
	)

	wypokError = wh.getObjectFromJson(responseBody, &entries)
	return
}

// GetStreamLast6HoursHotEntries returns hot entries ("Gorące dyskusje") from Mikroblog which took place in last 6 hours
func (wh *WykopHandler) GetStreamLast6HoursHotEntries(page uint) (entries []Entry, wypokError *WykopError) {
	return wh.getStreamHotEntries(page, 6)
}

// GetStreamLast12HoursHotEntries returns hot entries ("Gorące dyskusje") from Mikroblog which took place in last 12 hours
func (wh *WykopHandler) GetStreamLast12HoursHotEntries(page uint) (entries []Entry, wypokError *WykopError) {
	return wh.getStreamHotEntries(page, 12)
}

// GetStreamLast24HoursHotEntries returns hot entries ("Gorące dyskusje") from Mikroblog which took place in last 24 hours
func (wh *WykopHandler) GetStreamLast24HoursHotEntries(page uint) (entries []Entry, wypokError *WykopError) {
	return wh.getStreamHotEntries(page, 24)
}

func (wh *WykopHandler) getStreamHotEntries(page, period uint) (entries []Entry, wypokError *WykopError) {
	responseBody := wh.sendPostRequestForBody(
		getStreamHotUrl(wh.appKey, page, period),
	)

	wypokError = wh.getObjectFromJson(responseBody, &entries)
	return
}

func getStreamIndexUrl(appKey string, page uint) string {
	return fmt.Sprintf(STREAM_INDEX, appKey, page)
}

func getStreamHotUrl(appKey string, page, period uint) string {
	return fmt.Sprintf(STREAM_HOT, appKey, page, period)
}
