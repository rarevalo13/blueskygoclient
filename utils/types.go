package utils

type BluSkySession struct {
	AccessJwt       string                 `json:"accessJwt"`
	RefreshJwt      string                 `json:"refreshJwt"`
	Handle          string                 `json:"handle"`
	Did             string                 `json:"did"`
	DidDoc          map[string]interface{} `json:"didDoc"` // Using map for a flexible JSON object
	Email           string                 `json:"email,omitempty"`
	EmailConfirmed  bool                   `json:"emailConfirmed,omitempty"`
	EmailAuthFactor bool                   `json:"emailAuthFactor,omitempty"`
	Active          bool                   `json:"active"`
	Status          string                 `json:"status"`
}

// Profile represents the main JSON object.
type Profile struct {
	Did            string     `json:"did"`
	Handle         string     `json:"handle"`
	DisplayName    string     `json:"displayName"`
	Avatar         string     `json:"avatar"`
	Associated     Associated `json:"associated"`
	Labels         []string   `json:"labels"`
	CreatedAt      string     `json:"createdAt"`
	Description    string     `json:"description"`
	IndexedAt      string     `json:"indexedAt"`
	FollowersCount int        `json:"followersCount"`
	FollowsCount   int        `json:"followsCount"`
	PostsCount     int        `json:"postsCount"`
}

// Associated represents the nested "associated" object.
type Associated struct {
	Lists        int  `json:"lists"`
	Feedgens     int  `json:"feedgens"`
	StarterPacks int  `json:"starterPacks"`
	Labeler      bool `json:"labeler"`
	Chat         Chat `json:"chat,omitempty"`
}

// Post represents the top-level JSON structure.
type Post struct {
	Repo       string `json:"repo"`
	Collection string `json:"collection"`
	Record     Record `json:"record"`
}
type TimlinePost struct {
	URI         string     `json:"uri"`
	CID         string     `json:"cid"`
	Author      Author     `json:"author"`
	Record      Record     `json:"record"`
	ReplyCount  int        `json:"replyCount"`
	RepostCount int        `json:"repostCount"`
	LikeCount   int        `json:"likeCount"`
	QuoteCount  int        `json:"quoteCount"`
	IndexedAt   string     `json:"indexedAt"`
	Viewer      PostViewer `json:"viewer"`
	Labels      []Label    `json:"labels"`
}
type PostViewer struct {
	ThreadMuted       bool `json:"threadMuted"`
	EmbeddingDisabled bool `json:"embeddingDisabled"`
}

// Record represents the nested JSON object.
type Record struct {
	Text      string `json:"text,omitempty"`
	CreatedAt string `json:"createdAt"`
	Type      string `json:"$type,omitempty"`
}

// TimelineResponse is the top-level response.
type TimelineResponse struct {
	Cursor string     `json:"cursor"`
	Feed   []FeedItem `json:"feed"`
}

// FeedItem represents one item in the feed.
type FeedItem struct {
	Post TimlinePost `json:"post"`
}

// Reply holds the reply details.
type Reply struct {
	GrandparentAuthor Author `json:"grandparentAuthor"`
}

// Reason holds the reason details.
type Reason struct {
	By        Author `json:"by"`
	IndexedAt string `json:"indexedAt"`
}

// Author represents a user (or author) structure.
// Author represents the author of the post.
type Author struct {
	DID         string       `json:"did"`
	Handle      string       `json:"handle"`
	DisplayName string       `json:"displayName"`
	Avatar      string       `json:"avatar"`
	Viewer      AuthorViewer `json:"viewer"`
	Labels      []Label      `json:"labels"`
	CreatedAt   string       `json:"createdAt"`
}

// AuthorViewer represents viewing details in the author's context.
type AuthorViewer struct {
	Muted     bool `json:"muted"`
	BlockedBy bool `json:"blockedBy"`
}

// Associated holds numbers and a nested Chat object.

// Chat holds chat-specific settings.
type Chat struct {
	AllowIncoming string `json:"allowIncoming"`
}

// Viewer holds information about the viewing relationship.
type Viewer struct {
	Muted          bool           `json:"muted"`
	MutedByList    ListInfo       `json:"mutedByList"`
	BlockedBy      bool           `json:"blockedBy"`
	Blocking       string         `json:"blocking"`
	BlockingByList ListInfo       `json:"blockingByList"`
	Following      string         `json:"following"`
	FollowedBy     string         `json:"followedBy"`
	KnownFollowers KnownFollowers `json:"knownFollowers"`
}

// ListInfo represents information about a list (for mutedByList, blockingByList).
type ListInfo struct {
	URI           string     `json:"uri"`
	CID           string     `json:"cid"`
	Name          string     `json:"name"`
	Purpose       string     `json:"purpose"`
	Avatar        string     `json:"avatar"`
	ListItemCount int        `json:"listItemCount"`
	Labels        []Label    `json:"labels"`
	Viewer        ListViewer `json:"viewer"`
	IndexedAt     string     `json:"indexedAt"`
}

// ListViewer holds the viewer details inside a list.
type ListViewer struct {
	Muted   bool   `json:"muted"`
	Blocked string `json:"blocked"`
}

// KnownFollowers holds a count and a list of followers.
type KnownFollowers struct {
	Count     int           `json:"count"`
	Followers []interface{} `json:"followers"`
}

// Label represents a single label.
type Label struct {
	Ver int    `json:"ver"`
	Src string `json:"src"`
	URI string `json:"uri"`
	CID string `json:"cid"`
	Val string `json:"val"`
	Neg bool   `json:"neg"`
	CTS string `json:"cts"`
	EXP string `json:"exp"`
	Sig string `json:"sig"`
}
