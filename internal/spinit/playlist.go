package spinit

import (
	"spin-it/internal/models"
	"spin-it/internal/spotifyclient"

	log "github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

// CreatePlaylist todo
func CreatePlaylist(token *oauth2.Token, name string, description string, public bool, tracks *[]models.Track) (*models.Playlist, error) {
	c := spotifyclient.NewClient(token)
	c.AutoRetry = true

	user, err := c.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Creating playlist for:", user.ID)

	var ids []string

	for _, t := range *tracks {
		ids = append(ids, t.ID)
	}

	pl, err := spotifyclient.CreatePlaylistForUser(&c, user.ID, name, description, public)
	log.Debug("Playlist Created: ", name, "; ID: ", pl)

	if err != nil {
		log.Error(err)
	}

	snapshotID, err := addTracksToPlaylist(&c, pl.ID, ids...)
	log.Debug(len(*tracks), " added to playlist ", name, "with snapshotID ", snapshotID)

	if err != nil {
		log.Error(err)
	}

	pl.TotalTracks = len(*tracks)

	return &pl, err
}

func addTracksToPlaylist(client *spotify.Client, playlistID string, trackIDs ...string) (snapshotID string, err error) {
	const maxTracks = 100
	var divided [][]string

	chunkSize := (len(trackIDs) + maxTracks - 1) / maxTracks

	for i := 0; i < len(trackIDs); i += chunkSize {
		end := i + chunkSize

		if end > len(trackIDs) {
			end = len(trackIDs)
		}

		divided = append(divided, trackIDs[i:end])
	}

	for _, ids := range divided {
		snapshotID, err = spotifyclient.AddTracksToPlaylist(client, playlistID, ids...)
	}

	return snapshotID, err
}
