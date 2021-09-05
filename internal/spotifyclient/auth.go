package spotifyclient

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

const redirectURI = "http://localhost:8080/oauth2/callback"

var (
	auth   = spotify.NewAuthenticator(redirectURI, scopes...)
	scopes = []string{
		spotify.ScopePlaylistReadCollaborative,
		spotify.ScopePlaylistReadPrivate,
		spotify.ScopePlaylistModifyPublic,
		spotify.ScopePlaylistModifyPrivate,
		spotify.ScopeUserReadPrivate,
	}
)

// Auth todo
func Auth() (string, *http.Cookie) {
	state := uuid.New().String()

	log.Debug(state)

	cookie := http.Cookie{
		Name:    "spin-it.session",
		Value:   state,
		Expires: time.Now().Add(30 * time.Minute),
		Path:    "/",
	}

	url := auth.AuthURL(state)
	log.Debug("Please log in to Spotify by visiting the following page in your browser:", url)

	return url, &cookie
}

// Token todo
func Token(state string, r *http.Request) (*oauth2.Token, error) {
	return auth.Token(state, r)
}

// NewClient todo
func NewClient(token *oauth2.Token) spotify.Client {
	return auth.NewClient(token)
}
