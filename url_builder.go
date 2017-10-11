package go_wypok

const (
	LOGIN_URL = "https://a.wykop.pl/user/login"

	PROFILE_INDEX          = "https://a.wykop.pl/profile/index/"
	PROFILE_ENTRIES        = "https://a.wykop.pl/profile/entries/"
	PROFILE_COMMENTS       = "https://a.wykop.pl/profile/comments/"
	PROFILE_ENTRY_COMMENTS = "https://a.wykop.pl/profile/entriesComments/"

	ENTRY_INDEX          = "https://a.wykop.pl/entries/index/"
	ENTRY_ADD            = "https://a.wykop.pl/entries/add/"
	ENTRY_EDIT           = "https://a.wykop.pl/entries/edit/"
	ENTRY_DELETE         = "https://a.wykop.pl/entries/delete/"
	ENTRY_ADD_COMMENT    = "https://a.wykop.pl/entries/addComment/"
	ENTRY_COMMENT_EDIT   = "https://a.wykop.pl/entries/editComment/"
	ENTRY_COMMENT_DELETE = "https://a.wykop.pl/entries/deleteComment/"
	ENTRY_VOTE           = "https://a.wykop.pl/entries/vote/"
	ENTRY_UNVOTE         = "https://a.wykop.pl/entries/unvote/"
	ENTRY_FAVORITE       = "https://a.wykop.pl/entries/favorite/"

	RANK_INDEX = "https://a.wykop.pl/rank/index/appkey/%s/order/%s"

	MAIN_PAGE     = "https://a.wykop.pl/links/promoted/"
	UPCOMING_PAGE = "https://a.wykop.pl/links/upcoming/"

	ENTRIES_FROM_TAG = "https://a.wykop.pl/tag/entries/"
)

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

func getTagEntries(tag string) string {
	return ENTRIES_FROM_TAG + tag
}
