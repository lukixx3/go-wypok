package go_wypok

import "fmt"

const (
	LOGIN_URL = "https://a.wykop.pl/user/login/appkey/%s"

	PROFILE_INDEX          = "https://a.wykop.pl/profile/index/%s/appkey/%s"
	PROFILE_ADDED          = "https://a.wykop.pl/profile/added/%s/appkey/%s/page/%d"
	PROFILE_PUBLISHED      = "https://a.wykop.pl/profile/published/%s/appkey/%s/page/%d"
	PROFILE_COMMENTED      = "https://a.wykop.pl/profile/commented/%s/appkey/%s/page/%d"
	PROFILE_COMMENTS       = "https://a.wykop.pl/profile/comments/%s/appkey/%s/userkey/%s/page/%d"
	PROFILE_DIGGED         = "https://a.wykop.pl/profile/digged/%s/appkey/%s/page/%d"
	PROFILE_BURIED         = "https://a.wykop.pl/profile/buried/%s/appkey/%s/userkey/%s/page/%d"
	PROFILE_FAVORITES      = "https://a.wykop.pl/profile/favorites/%s/appkey/%s/userkey/%s"
	PROFILE_ENTRIES        = "https://a.wykop.pl/profile/entries/%s/appkey/%s/userkey/%s/page/%d"
	PROFILE_ENTRY_COMMENTS = "https://a.wykop.pl/profile/entriesComments/%s/appkey/%s/userkey/%s/page/%d"

	ENTRY_INDEX          = "https://a.wykop.pl/entries/index/%d/appkey/%s"
	ENTRY_ADD            = "https://a.wykop.pl/entries/add/appkey/%s/userkey/%s"
	ENTRY_EDIT           = "https://a.wykop.pl/entries/edit/%d/appkey/%s/userkey/%s"
	ENTRY_DELETE         = "https://a.wykop.pl/entries/delete/%d/appkey/%s/userkey/%s"
	ENTRY_ADD_COMMENT    = "https://a.wykop.pl/entries/addComment/%d/appkey/%s/userkey/%s"
	ENTRY_COMMENT_EDIT   = "https://a.wykop.pl/entries/editComment/%d/%d/appkey/%s/userkey/%s"
	ENTRY_COMMENT_DELETE = "https://a.wykop.pl/entries/deleteComment/%d/%d/appkey/%s/userkey/%s"
	ENTRY_VOTE           = "https://a.wykop.pl/entries/vote/entry/%d/appkey/%s/userkey/%s"
	ENTRY_COMMENT_VOTE   = "https://a.wykop.pl/entries/vote/comment/%d/%d/appkey/%s/userkey/%s"
	ENTRY_UNVOTE         = "https://a.wykop.pl/entries/unvote/entry/%d/appkey/%s/userkey/%s"
	ENTRY_COMMENT_UNVOTE = "https://a.wykop.pl/entries/unvote/comment/%d/%d/appkey/%s/userkey/%s"
	ENTRY_FAVORITE       = "https://a.wykop.pl/entries/favorite/%d/appkey/%s/userkey/%s"

	RANK_INDEX = "https://a.wykop.pl/rank/index/appkey/%s/order/%s"

	MAIN_PAGE     = "https://a.wykop.pl/links/promoted/"
	UPCOMING_PAGE = "https://a.wykop.pl/links/upcoming/"

	ENTRIES_FROM_TAG = "https://a.wykop.pl/tag/entries/%s/appkey/%s/page/%d"

	FAVORITES_INDEX    = "https://a.wykop.pl/favorites/index/%s/appkey/%s/userkey/%s"
	FAVORITES_LISTS    = "https://a.wykop.pl/favorites/lists/appkey/%s/userkey/%s"
	FAVORITES_COMMENTS = "https://a.wykop.pl/favorites/comments/appkey/%s/userkey/%s"
	FAVORITES_ENTRIES  = "https://a.wykop.pl/favorites/entries/appkey/%s/userkey/%s"
)

func getLoginUrl(appkey string) string {
	return fmt.Sprintf(LOGIN_URL, appkey)
}

func getTagEntries(tag, appkey string, page uint) string {
	return fmt.Sprintf(ENTRIES_FROM_TAG, tag, appkey, page)
}
