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
}

// Post represents the top-level JSON structure.
type Post struct {
	Repo       string `json:"repo"`
	Collection string `json:"collection"`
	Record     Record `json:"record"`
}

// Record represents the nested JSON object.
type Record struct {
	Text      string `json:"text,omitempty"`
	CreatedAt string `json:"createdAt"`
}
