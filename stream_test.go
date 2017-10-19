package go_wypok

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWykopHandlerGetStreamLast6HoursHotEntries(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	entries, wypokError := wh.GetStreamLast6HoursHotEntries(0)
	checkGetStreamFunctionResult(t, entries, wypokError)
}

func TestWykopHandlerGetStreamLast12HoursHotEntries(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	entries, wypokError := wh.GetStreamLast12HoursHotEntries(0)
	checkGetStreamFunctionResult(t, entries, wypokError)
}

func TestWykopHandlerGetStreamLast24HoursHotEntries(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	entries, wypokError := wh.GetStreamLast24HoursHotEntries(0)
	checkGetStreamFunctionResult(t, entries, wypokError)
}

func TestWykopHandlerGetStreamEntries(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	entries, wypokError := wh.GetStreamEntries(0)
	checkGetStreamFunctionResult(t, entries, wypokError)
}

func checkGetStreamFunctionResult(t *testing.T, entries []Entry, wError *WykopError) {
	assert.Nil(t, wError, "Expected that error will be nil")
	assert.NotEmpty(t, entries, "Expected that entries list won't be empty. It's impossible!")
}

func TestBuildingStreamURLs(t *testing.T) {
	testCases := []struct{ actual, expected string }{
		struct{ actual, expected string }{getStreamIndexUrl("APPKEY", 0), "https://a.wykop.pl/stream/index/appkey/APPKEY/page/0"},
		struct{ actual, expected string }{getStreamHotUrl("APPKEY", 0, 1), "https://a.wykop.pl/stream/hot/appkey/APPKEY/page/0/period/1"},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, testCase.actual)
	}
}
