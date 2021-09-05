package lastfmclient

import (
	"spin-it/internal/models"
	"time"

	"github.com/shkh/lastfm-go/lastfm"
	log "github.com/sirupsen/logrus"
)

// Period constants
const (
	PeriodOverall = "overall"
	Period7Day    = "7day"
	Period1month  = "1month"
	Period3month  = "3month"
	Period6month  = "6month"
	Period12month = "12month"
)

const layoutISO = "2006-01-02"

// GetTrackChart todo
func GetTrackChart(api *lastfm.Api, user string, from string, to string) (tracks []models.Track, err error) {
	f, _ := time.Parse(layoutISO, from)
	t, _ := time.Parse(layoutISO, to)

	result, err := api.User.GetWeeklyTrackChart(lastfm.P{"user": user, "from": f.Unix(), "to": t.Unix()})
	log.Debug("total: ", len(result.Tracks), " from: ", f.Unix(), " to: ", t.Unix())

	var ts []models.Track

	for _, t := range result.Tracks {
		var as []models.Artist
		var a models.Artist
		a.Name = t.Artist.Name
		as = append(as, a)
		track := models.Track{
			Name:      t.Name,
			Artists:   as,
			Rank:      t.Rank,
			PlayCount: t.PlayCount,
		}

		ts = append(ts, track)
	}

	return ts, err
}

// GetTopTracks todo
// func GetTrackChart(api *lastfm.Api, user string, from string, to string) (tracks []models.Track, err error) {
func GetTopTracks(api *lastfm.Api, user string, period string, limit string) (tracks []models.Track, err error) {
	result, err := api.User.GetTopTracks(lastfm.P{"user": user, "period": period, "limit": limit})
	log.Debug("total: ", len(result.Tracks))

	var ts []models.Track

	for _, t := range result.Tracks {
		var as []models.Artist
		var a models.Artist
		a.Name = t.Artist.Name
		as = append(as, a)
		track := models.Track{
			Name:      t.Name,
			Artists:   as,
			Rank:      t.Rank,
			PlayCount: t.PlayCount,
		}

		ts = append(ts, track)
	}

	return ts, err
}

// func mapLastFMToTracks() []models.Track {

// }
