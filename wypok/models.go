package Wykop

import "time"

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

type Entry struct {
	Id                  int
	Author              string
	Author_avatar       string
	Author_avatar_big   string
	Author_avatar_med   string
	Author_avatar_lo    string
	Author_group        int
	Author_sex          string
	Date                WypokShitDate `json:"date"`
	Body                string
	Source              string
	Url                 string
	Receiver            string
	Receiver_avatar     string
	Receiver_avatar_big string
	Receiver_avatar_med string
	Receiver_avatar_lo  string
	Receiver_group      string
	Receiver_sex        string
	Comments            []EntryComment
	Blocked             bool
	Vote_count          int
	User_vote           int
	Voters              []Voter
	User_favorite       bool
	TypeE               string `json:"type"`
	Embed               Embed
	Deleted             bool
	Violation_url       string
	Can_comment         bool
	App                 string
	Comment_count       int
}

type EntryComment struct {
	Id                int
	Author            string
	Author_avatar     string
	Author_avatar_big string
	Author_avatar_med string
	Author_avatar_lo  string
	Author_group      int
	Author_sex        string
	Date              WypokShitDate `json:"date"`
	Body              string
	Source            string
	Entry_id          int
	Blocked           bool
	Deleted           bool
	Vote_count        int
	User_vote         int
	Voters            []Voter
	Embed             Embed
	Entry             Entry
	Type              string
	App               string
	Violation_url     string
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
	Author            string
	Author_group      int
	author_avatar     string
	author_avatar_big string
	author_avatar_med string
	author_avatar_lo  string
	author_sex        string
	date              string
}

type Embed struct {
	EmbedType string `json:"type"`
	Preview   string
	Url       string
	Source    string
	Plus18    bool
}

type TagsEntries struct {
	Meta  Meta
	Items []Entry
}

type Meta struct {
	Tags        string
	Is_Observed bool
	Is_Blocked  bool
	Counters    Counters
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
	Id string `json:"id"`
}

type CommentResponse struct {
	Id string `json:"id"`
}

type WykopError struct {
	ErrorObject WykopErrorMessage `json:"error"`
}

type WykopErrorMessage struct {
	Code    int
	Message string
}
