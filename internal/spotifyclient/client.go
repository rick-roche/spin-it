package spotifyclient

import "github.com/zmb3/spotify"

type spotifyClient interface {
	CreatePlaylistForUser(userID, playlistName, description string, public bool) (*spotify.FullPlaylist, error)
	AddTracksToPlaylist(playlistID spotify.ID, trackIDs ...spotify.ID) (snapshotID string, err error)
	Search(query string, t spotify.SearchType) (*spotify.SearchResult, error)
}
