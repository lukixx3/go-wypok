package go_wypok

import (
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

type Entry struct {
	Id                int
	Author            string
	AuthorAvatar      string        `json:"author_avatar"`
	AuthorAvatarBig   string        `json:"author_big"`
	AuthorAvatarMed   string        `json:"author_med"`
	AuthorAvatarLo    string        `json:"author_lo"`
	AuthorGroup       int           `json:"author_group"`
	AuthorSex         string        `json:"author_sex"`
	Date              WypokShitDate `json:"date"`
	Body              string
	Source            string
	Url               string
	Receiver          string
	ReceiverAvatar    string `json:"receiver_avatar"`
	ReceiverAvatarBig string `json:"receiver_avatar_big"`
	ReceiverAvatarMed string `json:"receiver_avatar_med"`
	ReceiverAvatarLo  string `json:"receiver_avatar_lo"`
	ReceiverGroup     string `json:"receiver_group"`
	ReceiverSex       string `json:"receiver_sex"`
	Comments          []EntryComment
	Blocked           bool
	VoteCount         int `json:"vote_count"`
	UserVote          int `json:"user_vote"`
	Voters            []Voter
	UserFavorite      bool   `json:"user_favorite"`
	TypeE             string `json:"type"`
	Embed             Embed
	Deleted           bool
	ViolationUrl      string `json:"violation_url"`
	CanComment        bool   `json:"can_comment"`
	App               string
	CommentCount      int `json:"comment_count"`
}

type EntryComment struct {
	ID              int           `json:"id"`
	Author          string        `json:"author"`
	AuthorAvatar    string        `json:"author_avatar"`
	AuthorAvatarBig string        `json:"author_avatar_big"`
	AuthorAvatarMed string        `json:"author_avatar_med"`
	AuthorAvatarLo  string        `json:"author_avatar_lo"`
	AuthorGroup     int           `json:"author_group"`
	AuthorSex       string        `json:"author_sex"`
	Date            WypokShitDate `json:"date"`
	Body            string        `json:"body"`
	Source          string        `json:"source"`
	EntryID         int           `json:"entry_id"`
	Blocked         bool          `json:"blocked"`
	Deleted         bool          `json:"deleted"`
	VoteCount       int           `json:"vote_count"`
	UserVote        int           `json:"user_vote"`
	Voters          []Voter
	Embed           Embed  `json:"embed"`
	Type            string `json:"type"`
	App             string `json:"app"`
	ViolationURL    string `json:"violation_url"`
	Entry           Entry
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

type AuthenticationResponse struct {
	Login        string
	Email        string
	ViolationUrl string `json:"violation_url"`
	Userkey      string
}

type EntryResponse struct {
	// for some stupid reason wypok returns string here, convert it manually to int
	Id string `json:"id"`
}

type CommentResponse struct {
	// and here, it should be int
	Id string `json:"id"`
}

type WykopError struct {
	ErrorObject WykopErrorMessage `json:"error"`
}

type WykopErrorMessage struct {
	Code    int
	Message string
}
