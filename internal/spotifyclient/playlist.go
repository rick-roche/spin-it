package spotifyclient

import (
	"spin-it/internal/models"

	"github.com/zmb3/spotify"
)

// CreatePlaylistForUser todo
func CreatePlaylistForUser(client spotifyClient, userID string, playlistName string, description string, public bool) (models.Playlist, error) {

	res, err := client.CreatePlaylistForUser(userID, playlistName, description, public)

	return models.Playlist{
		ID:       res.SimplePlaylist.ID.String(),
		Name:     res.SimplePlaylist.Name,
		IsPublic: res.SimplePlaylist.IsPublic,
	}, err
}

// AddTracksToPlaylist todo
func AddTracksToPlaylist(client spotifyClient, playlistID string, trackIDs ...string) (snapshotID string, err error) {

	var ids []spotify.ID

	for _, i := range trackIDs {
		ids = append(ids, spotify.ID(i))
	}

	return client.AddTracksToPlaylist(spotify.ID(playlistID), ids...)
}
