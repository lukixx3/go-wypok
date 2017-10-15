package go_wypok

import (
	"strconv"
	"time"
)

type Profile struct {
	ID              int64
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

type LinkComment struct {
	ID              int    `json:"id"`
	Date            string `json:"date"`
	Author          string `json:"author"`
	AuthorGroup     int    `json:"author_group"`
	AuthorAvatar    string `json:"author_avatar"`
	AuthorAvatarBig string `json:"author_avatar_big"`
	AuthorAvatarMed string `json:"author_avatar_med"`
	AuthorAvatarLo  string `json:"author_avatar_lo"`
	AuthorSex       string `json:"author_sex"`
	VoteCount       int    `json:"vote_count"`
	VoteCountPlus   int    `json:"vote_count_plus"`
	VoteCountMinus  int    `json:"vote_count_minus"`
	Body            string `json:"body"`
	Source          string `json:"source"`
	ParentID        int    `json:"parent_id"`
	Status          string `json:"status"`
	CanVote         bool   `json:"can_vote"`
	UserVote        bool   `json:"user_vote"`
	Blocked         bool   `json:"blocked"`
	Deleted         bool   `json:"deleted"`
	Embed           Embed  `json:"embed"`
	Type            string `json:"type"`
	App             string `json:"app"`
	UserFavorite    bool   `json:"user_favorite"`
	ViolationURL    string `json:"violation_url"`
	Link            Link
}

type WykopShitUserVote string

// when user is NOT logged in, user_vote field will not be provided
// when user is logged in, user_vote might have value of "bury" or "dig"
// if user made such action in the past, however, when user didn't vote
// on the link wypok.pl provides user_vote as false (see no quotes), which
// golang unmarshaller tries to read as `bool` not `string` and panics.
// This wrapper type ensures that bool will be converted to string and string treated as string.
func (value *WykopShitUserVote) UnmarshalJSON(data []byte) error {
	asString := string(data)
	if asString == "dig" {
		*value = "dig"
	} else if asString == "bury" {
		*value = "bury"
	} else {
		*value = "false"
	}
	return nil
}

type WypokShitDate struct {
	time.Time
}

func (self *WypokShitDate) UnmarshalJSON(b []byte) (err error) {
	s := string(b)

	// Get rid of the quotes "" around the value.
	// A second option would be to include them
	// in the date format string instead, like so below:
	//   time.Parse(`"`+time.RFC3339Nano+`"`, s)
	s = s[1 : len(s)-1]

	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		t, err = time.Parse("2006-01-02 15:04:05", s)
	}
	self.Time = t
	return
}

type Voter struct {
	Author          string
	AuthorAvatar    string        `json:"author_avatar"`
	AuthorAvatarBig string        `json:"author_big"`
	AuthorAvatarMed string        `json:"author_med"`
	AuthorAvatarLo  string        `json:"author_lo"`
	AuthorSex       string        `json:"author_sex"`
	AuthorGroup     int           `json:"author_group"`
	Date            WypokShitDate `json:"date"`
}

type Embed struct {
	EmbedType string `json:"type"`
	Preview   string
	Url       string
	Source    string
	Plus18    bool `json:"plus_18"`
}

type TagsEntries struct {
	Meta  Meta
	Items []Entry
}

type Meta struct {
	Tags       string
	IsObserved bool `json:"is_observed"`
	IsBlocked  bool `json:"is_blocked"`
	Counters   Counters
}

type Counters struct {
	Total   int
	Entries int
	Links   int
}

type VoteResponse struct {
	Vote int
	//Voters string
}

type FavoriteResponse struct {
	UserFavorite bool `json:"user_favorite"`
}

type AuthenticationResponse struct {
	Login        string
	Email        string
	ViolationUrl string `json:"violation_url"`
	Userkey      string
}

type EntryResponse struct {
	Id int
}

type CommentResponse struct {
	Id int
}

// WTF wykop.pl API is so retarded. Everywhere it returns int as Id on any response class
// by comment/entry submitted response returns id in quotes "", literally - integer inside quotes,
// which makes it a string, so these struct wrapper classes unmarshall it to string
// then the struct is converted to a new one, where id is an int. bravo białkov. bravo.
// thanks for reading this rant, please enjoy comment at the bottom of this file.
type stringEntryResponse struct {
	// for some stupid reason wypok returns string here
	Id string `json:"id"`
}

type stringCommentResponse struct {
	// and here, it should be int
	Id string `json:"id"`
}

func newEntryResponse(stringIdResponse stringEntryResponse) EntryResponse {
	entryResponse := EntryResponse{}
	theId, err := strconv.Atoi(stringIdResponse.Id)
	if err != nil {
		entryResponse.Id = 0
	} else {
		entryResponse.Id = theId
	}
	return entryResponse
}

func newCommentResponse(stringIdResponse stringCommentResponse) CommentResponse {
	commentResponse := CommentResponse{}
	theId, err := strconv.Atoi(stringIdResponse.Id)
	if err != nil {
		commentResponse.Id = 0
	} else {
		commentResponse.Id = theId
	}
	return commentResponse
}

type WykopError struct {
	ErrorObject WykopErrorMessage `json:"error"`
}

type WykopErrorMessage struct {
	Code    int
	Message string
}

type ConversationsList struct {
	Conversations []Conversation
}
type Conversation struct {
	Author string `json:author`
	LastUpdate string `json:last_update`
	AuthorAvatar string `json:author_avatar`
	AuthorAvatarBig string `json:author_avatar_big`
	AuthorAvatarMed string `json:author_avatar_med`
	AuthorAvatarLow string `json:author_avatar_Lo`
	AuthorGroup string `json:author_group`
	AuthorSex string `json:author_sex`
	Status string `json:status`
}

type PrivateMessage struct {
	Author string `json:author`
	AuthorAvatar string `json:author_avatar`
	AuthorAvatarBig string `json:author_avatar_big`
	AuthorAvatarMed string `json:author_avatar_med`
	AuthorAvatarLow string `json:author_avatar_Lo`
	AuthorGroup string `json:author_group`
	AuthorSex string `json:author_sex`
	Status string `json:status`
	Date string `json:date`
	Content string `json:body`
	Direction string `json:direction`
	Embed string `json:embed`
	App string `json:app`
}