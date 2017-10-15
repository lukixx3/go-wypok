package go_wypok

import "fmt"

type Profile struct {
	Id              int
	Login           string
	Email           string
	PublicEmail     string `json:"public_email"`
	Name            string
	Www             string
	Jabber          string
	Gg              int
	City            string
	About           string
	AuthorGroup     int `json:"author_group"`
	LinksAdded      int `json:"links_added"`
	LinksPublished  int `json:"links_published"`
	Comments        int
	Rank            int
	Followers       int
	Following       int
	Entries         int
	EntriesComments int `json:"entries_comments"`
	Diggs           int
	Buries          int
	RelatedLinks    int `json:"related_links"`
	Groups          int
	Sex             string
	Avatar          string
	AvatarLo        string `json:"avatar_lo"`
	AvatarMed       string `json:"avatar_med"`
	AvatarBig       string `json:"avatar_big"`
	IsObserved      bool   `json:"is_observed"`
}

func getProfileUrl(username, appkey string) string {
	return fmt.Sprintf(PROFILE_INDEX, username, appkey)
}

func getProfileFavoritesUrl(uesrname, appkey, userkey string) string {
	return fmt.Sprintf(PROFILE_FAVORITES, appkey, userkey)
}

func getProfileEntriesUrl(username, appkey, userkey string, page int) string {
	return fmt.Sprintf(PROFILE_ENTRIES, username, appkey, userkey, page)
}

func getProfileCommentsUrl(username, appkey, userkey string, page int) string {
	return fmt.Sprintf(PROFILE_COMMENTS, username, appkey, userkey, page)
}

func getProfileEntriesCommentsUrl(username, appkey, userkey string, page int) string {
	return fmt.Sprintf(PROFILE_ENTRY_COMMENTS, username, appkey, userkey, page)
}
