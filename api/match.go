package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"spin-it/internal/models"
	"spin-it/internal/spinit"

	log "github.com/sirupsen/logrus"
)

type MatchParams struct {
	Source string `schema:"source,required"`
	Type   string `schema:"type,required"`
	User   string `schema:"user,required"`
	From   string `schema:"from"`
	To     string `schema:"to"`
	Period string `schema:"period"`
	Limit  int    `schema:"limit"`
}

// Match response with found tracks
// swagger:model MatchResponse
type MatchResponse struct {
	// in: body
	*models.FoundTracks
}

func (a *API) initialiseMatch() {
	// swagger:operation GET /v1/match match
	//
	// Match songs from last.fm or Discogs to Spotify
	//
	// Match songs from your favourite sources to tracks on Spotify
	// ---
	// parameters:
	// - name: source
	//   in: query
	//   description: source to match against. e.g. lastfm
	//   type: string
	//   enum:
	//    - lastfm
	//    - discogs
	//   required: true
	// - name: type
	//   in: query
	//   description: type of match. e.g. chart
	//   type: string
	//   enum:
	//    - chart
	//    - top-tracks
	//   required: true
	// - name: user
	//   in: query
	//   description: user to use on the source platform. e.g. your last.fm username
	//   type: string
	//   required: true
	// - name: from
	//   in: query
	//   description: from date
	//   type: string
	//   format: date
	//   required: false
	// - name: to
	//   in: query
	//   description: to date
	//   type: string
	//   format: date
	//   required: false
	// - name: period
	//   in: query
	//   description: Period enum. e.g. overall
	//   type: string
	//   enum:
	//    - overall
	//    - 7day
	//    - 1month
	//    - 3month
	//    - 6month
	//    - 12month
	//   required: false
	// - name: limit
	//   in: query
	//   description: The number of results to fetch per page
	//   type: number
	//   required: false
	// responses:
	//   '200':
	//     description: Match response
	//     schema:
	//       "$ref": "#/definitions/MatchResponse"
	a.Router.HandleFunc("/v1/match", matchGetHandler)
	// a.Router.HandleFunc("/v1/match", matchGetHandler).Methods("GET")
}

func matchGetHandler(w http.ResponseWriter, r *http.Request) {
	token, err := extractToken(r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var mp MatchParams
	err = decoder.Decode(&mp, r.URL.Query())

	if err != nil {
		log.Info("Error in GET parameters : ", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	} else {
		log.Println("GET parameters : ", mp)
	}

	var res *models.FoundTracks

	switch mp.Source {
	case "lastfm":
		switch mp.Type {
		case "chart":
			fmt.Println("last fm: top charts")
			res, err = spinit.LastFMMatchTrackChart(token, mp.User, mp.From, mp.To, mp.Limit)

		case "top-tracks":
			fmt.Println("last fm: top tracks")
			res, err = spinit.LastFMMatchTopTracks(token, mp.User, mp.Period, mp.Limit)

		default:
			fmt.Printf("Invalid type: %s\n", mp.Type)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case "discogs":
		fmt.Println("discogs")

	default:
		fmt.Printf("Invalid source: %s\n", mp.Source)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(MatchResponse{res})
}
