package go_wypok

const (
	LOGIN_URL = "https://a.wykop.pl/user/login"

	PROFILE_INDEX          = "https://a.wykop.pl/profile/index/"
	PROFILE_ENTRIES        = "https://a.wykop.pl/profile/entries/"
	PROFILE_COMMENTS       = "https://a.wykop.pl/profile/comments/"
	PROFILE_ENTRY_COMMENTS = "https://a.wykop.pl/profile/entriesComments/"

	ENTRY_INDEX          = "https://a.wykop.pl/entries/index/%s/appkey/%s"
	ENTRY_ADD            = "https://a.wykop.pl/entries/add/appkey/%s/userkey/%s"
	ENTRY_EDIT           = "https://a.wykop.pl/entries/edit/%s/appkey/%s/userkey/%s"
	ENTRY_DELETE         = "https://a.wykop.pl/entries/delete/%s/appkey/%s/userkey/%s"
	ENTRY_ADD_COMMENT    = "https://a.wykop.pl/entries/addComment/%s/appkey/%s/userkey/%s"
	ENTRY_COMMENT_EDIT   = "https://a.wykop.pl/entries/editComment/%s/%s/appkey/%s/userkey/%s"
	ENTRY_COMMENT_DELETE = "https://a.wykop.pl/entries/deleteComment/%s/%s/appkey/%s/userkey/%s"
	ENTRY_VOTE           = "https://a.wykop.pl/entries/vote/entry/%s/appkey/%s/userkey/%s"
	ENTRY_COMMENT_VOTE   = "https://a.wykop.pl/entries/vote/comment/%s/%s/appkey/%s/userkey/%s"
	ENTRY_UNVOTE         = "https://a.wykop.pl/entries/unvote/entry/%s/appkey/%s/userkey/%s"
	ENTRY_COMMENT_UNVOTE = "https://a.wykop.pl/entries/unvote/comment/%s/%s/appkey/%s/userkey/%s"
	ENTRY_FAVORITE       = "https://a.wykop.pl/entries/favorite/%s/appkey/%s/userkey/%s"

	RANK_INDEX = "https://a.wykop.pl/rank/index/appkey/%s/order/%s"

	MAIN_PAGE     = "https://a.wykop.pl/links/promoted/"
	UPCOMING_PAGE = "https://a.wykop.pl/links/upcoming/"

	ENTRIES_FROM_TAG = "https://a.wykop.pl/tag/entries/"

	FAVORITES_INDEX = "https://a.wykop.pl/favorites/index/%s/appkey/%s/userkey/%s"
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
