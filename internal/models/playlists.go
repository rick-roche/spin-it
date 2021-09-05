package models

// Playlist ...
type Playlist struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	IsPublic    bool   `json:"public"`
	TotalTracks int    `json:"totalTracks"`
}
