package go_wypok

const (
	LOGIN_URL = "https://a.wykop.pl/user/login"

	PROFILE_INDEX          = "https://a.wykop.pl/profile/index/"
	PROFILE_ENTRIES        = "https://a.wykop.pl/profile/entries/"
	PROFILE_COMMENTS       = "https://a.wykop.pl/profile/comments/"
	PROFILE_ENTRY_COMMENTS = "https://a.wykop.pl/profile/entriesComments/"

	ENTRY_INDEX          = "https://a.wykop.pl/entries/index/"
	ENTRY_ADD            = "https://a.wykop.pl/entries/add"
	ENTRY_DELETE         = "https://a.wykop.pl/entries/delete/"
	ENTRY_VOTE           = "https://a.wykop.pl/entries/vote/"
	ENTRY_COMMENT_DELETE = "https://a.wykop.pl/entries/deleteComment/"

	ENTRIES_FROM_TAG = "https://a.wykop.pl/tag/entries/"

	MAIN_PAGE     = "https://a.wykop.pl/links/promoted/"
	UPCOMING_PAGE = "https://a.wykop.pl/links/upcoming/"
)

func getMainPageUrl() string {
	return MAIN_PAGE
}

func getUpcomingPageUrl() string {
	return UPCOMING_PAGE
}

func getLoginUrl(appkey string) string {
	return LOGIN_URL + "/appkey/" + appkey
}

func getProfileUrl(username string) string {
	return PROFILE_INDEX + username
}

func getProfileEntriesUrl(username string) string {
	return PROFILE_ENTRIES + username
}

func getProfileCommentsUrl(username string) string {
	return PROFILE_COMMENTS + username
}

func getProfileEntriesCommentsUrl(username string) string {
	return PROFILE_ENTRY_COMMENTS + username
}

func getEntryUrl(entry string) string {
	return ENTRY_INDEX + entry
}

func getDeleteEntryUrl(entry string) string {
	return ENTRY_DELETE + entry
}

func getDeleteCommentUrl(entryId string, commentId string) string {
	return ENTRY_COMMENT_DELETE + entryId + "/" + commentId
}

func getAddEntryUrl() string {
	return ENTRY_ADD
}

func getEntryVoteUrl(objectType string, entryId string, commentId string) string {
	return ENTRY_VOTE + objectType + "/" + entryId + "/" + commentId
}

func getTagEntries(tag string) string {
	return ENTRIES_FROM_TAG + tag
}
