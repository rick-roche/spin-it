package api

import (
	"encoding/json"
	"net/http"
	"spin-it/internal/spotifyclient"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

// Auth response with token
// swagger:model AuthResponse
type authResponse struct {
	// in: body
	Token oauth2.Token `json:"token"`
}

func (a *API) initialiseAuth() {
	// swagger:operation GET /v1/auth authenticate
	//
	// Initiates the OAuth2 flow
	//
	// Authenticate with Spotify using your account
	// ---
	// responses:
	//   '200':
	//     description: Auth response
	//     schema:
	//       "$ref": "#/definitions/AuthResponse"
	a.Router.HandleFunc("/v1/auth", authGetHandler)
	// a.Router.HandleFunc("/v1/auth", authGetHandler).Methods("GET")

	a.Router.HandleFunc("/oauth2/callback", authCallbackHandler)
}

func authGetHandler(w http.ResponseWriter, r *http.Request) {
	url, c := spotifyclient.Auth()

	http.SetCookie(w, c)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

// the user will eventually be redirected back to your redirect URL
// typically you'll have a handler set up like the following:
func authCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Get oauth2 tokens
	stateCookie, err := r.Cookie("spin-it.session")
	if err != nil {
		log.Error("failed to get state cookie for oauth2: ", "err", err.Error())
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	// use the same state string here that you used to generate the URL
	token, err := spotifyclient.Token(stateCookie.Value, r)
	if err != nil {
		log.Error("Couldn't get token", "err", err.Error())
		http.Error(w, "Couldn't get token", http.StatusNotFound)
		return
	}
	// create a client using the specified token
	c := spotifyclient.NewClient(token)
	client = &c

	// the client can now be used to make authenticated requests
	// use the client to make calls that require authorization
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("You are logged in as:", user.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authResponse{*token})
}
