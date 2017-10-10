package go_wypok

type Link struct {
	Id              int
	Title           string
	Description     string
	Tags            string
	Url             string
	SourceUrl       string        `json:"source_url"`
	VoteCount       int           `json:"vote_count"`
	CommentCount    int           `json:"comment_count"`
	ReportCount     int           `json:"report_count"`
	Date            WypokShitDate `json:"date"`
	Author          string
	AuthorAvatar    string `json:"avatar_avatar"`
	AuthorAvatarMed string `json:"avatar_med"`
	AuthorAvatarLo  string `json:"avatar_lo"`
	AuthorGroup     int    `json:"author_group"`
	AuthorSex       string `json:"author_sex"`
	Preview         string
	UserLists       []int `json:"user_lists"`
	Plus18          bool  `json:"plus_18"`
	Status          string
	CanVote         bool `json:"can_vote"`
	IsHot           bool `json:"is_hot"`
	HasOwnContent   bool `json:"has_own_content"`
	Category        string
	CategoryName    string            `json:"category_name,omitempty"`
	UserVote        WykopShitUserVote `json:"user_vote,omitempty"`
	UserObserve     bool              `json:"user_observe,omitempty"`
	UserFavorite    bool              `json:"user_favorite,omitempty"`
}

func (wh *WykopHandler) GetMainPageLinks(page int) (mainPageLinks []Link, wypokError *WykopError) {
	urlAddress := getMainPageUrl() + appKeyPathElement + wh.appKey

	if wh.authResponse.Userkey != "" {
		urlAddress = urlAddress + userKeyPathElement + wh.authResponse.Userkey
	}

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &mainPageLinks)

	return
}

func (wh *WykopHandler) GetUpcomingLinks(page int) (mainPageLinks []Link, wypokError *WykopError) {
	urlAddress := getUpcomingPageUrl() + appKeyPathElement + wh.appKey

	if wh.authResponse.Userkey != "" {
		urlAddress = urlAddress + userKeyPathElement + wh.authResponse.Userkey
	}

	_, responseBody, _ := wh.preparePostRequest(urlAddress).End()

	wypokError = wh.getObjectFromJson(responseBody, &mainPageLinks)

	return
}

func getMainPageUrl() string {
	return MAIN_PAGE
}

func getUpcomingPageUrl() string {
	return UPCOMING_PAGE
}
