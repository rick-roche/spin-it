package lastfmclient

import (
	"os"

	"github.com/shkh/lastfm-go/lastfm"
)

// Auth todo
func Auth() *lastfm.Api {
	a := lastfm.New(os.Getenv("LASTFM_API_KEY"), "")

	return a
}
