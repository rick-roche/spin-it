package models

// Album ...
type Album struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

// Artist ...
type Artist struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

// Track ...
type Track struct {
	ID          string   `json:"id,omitempty"`
	Artists     []Artist `json:"artists"`
	Album       Album    `json:"album,omitempty"`
	Name        string   `json:"name"`
	Rank        string   `json:"rank,omitempty"`
	PlayCount   string   `json:"playcount,omitempty"`
	Confidence  string   `json:"confidence,omitempty"`
	SearchedFor string   `json:"searchedFor,omitempty"`
}

// SearchedTracks ...
type SearchedTracks struct {
	Matched   []Track `json:"matched"`
	Unmatched []Track `json:"unmatched"`
}

// FoundTracks ...
type FoundTracks struct {
	Matched        []Track `json:"matched"`
	Unmatched      []Track `json:"unmatched"`
	Total          int     `json:"total"`
	TotalMatched   int     `json:"totalMatched"`
	TotalUnmatched int     `json:"totalUnmatched"`
}
