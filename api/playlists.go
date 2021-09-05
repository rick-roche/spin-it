package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"spin-it/internal/models"
	"spin-it/internal/spinit"
)

// Create a new playlist
// swagger:model CreatePlaylistRequest
type CreatePlaylistRequest struct {
	// in: body
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Public      bool           `json:"public"`
	Tracks      []models.Track `json:"tracks"`
}

// Create a new playlist response
// swagger:model CreatePlaylistResponse
type CreatePlaylistResponse struct {
	// in: body
	*models.Playlist
}

func (a *API) initialisePlaylists() {
	// swagger:operation POST /v1/playlists playlists
	//
	// Create a playlist on Spotify using the matched tracks
	//
	// Create a new playlist on Spotify using the matched tracks from the match endpoint
	// ---
	// parameters:
	// - name: createPlaylist
	//   in: body
	//   description: The playlist to be created
	//   schema:
	//     "$ref": "#/definitions/CreatePlaylistRequest"
	//   required: true
	// responses:
	//   '200':
	//     description: Match response
	//     schema:
	//       "$ref": "#/definitions/CreatePlaylistResponse"
	a.Router.HandleFunc("/v1/playlists", playlistsPostHandler)
	// a.Router.HandleFunc("/v1/playlists", playlistsPostHandler).Methods("POST")
}

func playlistsPostHandler(w http.ResponseWriter, r *http.Request) {
	token, err := extractToken(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var req CreatePlaylistRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// handle error
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}

	pl, err := spinit.CreatePlaylist(token, req.Name, req.Description, req.Public, &req.Tracks)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(CreatePlaylistResponse{pl})
}
