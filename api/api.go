// Package classification Spin-It API
//
// Documentation of the Spin-It API.
//
//     Schemes: http, https
//     BasePath: /
//     Version: 0.0.1
// 	   License: MIT http://opensource.org/licenses/MIT
//     Contact: Rick Roche<a@b.com> https://www.rickroche.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//     oauth2:
//         type: oauth2
//         authorizationUrl: /oauth2/auth
//         tokenUrl: /oauth2/token
//         in: header
//         scopes:
//           bar: foo
//         flow: accessCode
//
// swagger:meta
package api

import (
	"errors"
	"net/http"
	"spin-it/internal/lastfmclient"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/schema"

	"github.com/shkh/lastfm-go/lastfm"
	log "github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

var (
	client  *spotify.Client
	lfm     *lastfm.Api
	decoder = schema.NewDecoder()
)

// API todo
type API struct {
	// Router  *mux.Router
	Router *chi.Mux
}

// Initialise todo
func (a *API) Initialise() {
	log.SetLevel(log.DebugLevel)

	a.Router = chi.NewRouter()
	a.setupMiddlewares()
	a.initialiseRoutes()

	lfm = lastfmclient.Auth()
}

// Run todo
func (a *API) Run(addr string) {
	log.Info("Running...")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *API) setupMiddlewares() {

	// Setup middlewares: https://github.com/go-chi/chi#core-middlewares
	a.Router.Use(
		middleware.Compress(5, "application/json"),
		middleware.Logger,
		middleware.RealIP,
		middleware.Recoverer, // Panic recovery.
		middleware.RequestID,
		// middleware.AllowContentEncoding()
		middleware.AllowContentType("application/json"),
	)
}

func (a *API) initialiseRoutes() {
	a.initialiseHealth()
	a.initialiseSwagger()
	a.initialiseAuth()
	a.initialiseMatch()
	a.initialisePlaylists()
}

func extractToken(r *http.Request) (*oauth2.Token, error) {
	a := r.Header.Get("Authorization")
	s := strings.Split(a, "Bearer")
	if len(s) != 2 {
		return nil, errors.New("No token present")
	}
	t := strings.TrimSpace(s[1])

	return &oauth2.Token{
		AccessToken: t,
	}, nil
}

// func searchTracks(w http.ResponseWriter, r *http.Request) {
// 	var a, _ = extractToken(r)
// 	log.Debug(a.AccessToken)

// 	// spotifyclient.SearchTrack(client, "natural child", "for the love of the game", "dtv")
// 	c := spotifyclient.NewClient(a)
// 	spotifyclient.SearchTrack(&c, "natural child", "for the love of the game", "dtv")
// }

// func getTopTracks(w http.ResponseWriter, r *http.Request) {
// 	lastfmclient.GetTopTracks(lfm, "rickdisco", lastfmclient.Period7Day, "100")

// 	// result, _ := lfm.User.GetTopTracks(lastfm.P{"user": "rickdisco", "period": "overall"})
// 	// fmt.Println(result)
// 	// fmt.Println(len(result.Tracks))

// 	// for _, track := range result.Tracks {
// 	// 	fmt.Println(track.Artist.Name + " - " + track.Name)
// 	// }
// }

// func getTrackChart(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// todo proper validation and error struct
// 	values := r.URL.Query()
// 	user := values.Get("user")
// 	if user == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode("user: Last.FM User is missing")
// 		return
// 	}
// 	from := values.Get("from")
// 	if from == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode("from date missing")
// 		return
// 	}
// 	to := values.Get("to")
// 	if to == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode("to date missing")
// 		return
// 	}

// 	t, _ := extractToken(r)
// 	found, _ := spinit.LastFMMatchTrackChart(t, user, from, to)

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(found)
// }
