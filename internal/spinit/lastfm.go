package spinit

import (
	"sort"
	"spin-it/internal/lastfmclient"
	"spin-it/internal/models"
	"spin-it/internal/spotifyclient"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

// LastFMMatchTrackChart Search lfm, match to spotify, return results
func LastFMMatchTrackChart(token *oauth2.Token, user string, from string, to string, limit int) (ft *models.FoundTracks, e error) {
	// clients
	c := spotifyclient.NewClient(token)
	c.AutoRetry = true

	pc := &c
	lfm := lastfmclient.Auth()

	// search tracks on lastfm
	lt, _ := lastfmclient.GetTrackChart(lfm, user, from, to)

	// match tracks on spotify
	var m []models.Track
	var u []models.Track

	// do this concurrently
	// guard := make(chan struct{}, 100)
	wg := sync.WaitGroup{}

	for i, t := range lt {

		if i == limit {
			log.Debug("Limit reached")
			break // break here
		}

		// guard <- struct{}{} // would block if guard channel is already filled

		wg.Add(1)

		go func(t models.Track) {
			st, err := spotifyclient.SearchTrack(pc, t.Artists[0].Name, t.Album.Name, t.Name)

			if err != nil {
				e = err
				wg.Done()
				return
			}

			// did we find something?
			if len(st) > 0 {
				mt := st[0]
				mt.Rank = t.Rank
				mt.PlayCount = t.PlayCount
				mt.SearchedFor = t.Artists[0].Name + " - " + t.Name
				m = append(m, mt)
			} else {
				u = append(u, t)
			}

			log.Debug("Searched for ", t.Artists[0].Name, " - ", t.Name)

			wg.Done()
			// <-guard
		}(t)
	}

	wg.Wait()

	sortByRank(m)
	sortByRank(u)

	// return matched and unmatched
	return &models.FoundTracks{
		Unmatched:      u,
		Matched:        m,
		Total:          len(u) + len(m),
		TotalMatched:   len(m),
		TotalUnmatched: len(u),
	}, e
}

// LastFMMatchTopTracks Search lfm, match to spotify, return results
func LastFMMatchTopTracks(token *oauth2.Token, user string, period string, limit int) (ft *models.FoundTracks, e error) {
	// clients
	c := spotifyclient.NewClient(token)
	c.AutoRetry = true

	pc := &c
	lfm := lastfmclient.Auth()

	// search tracks on lastfm
	lt, _ := lastfmclient.GetTopTracks(lfm, user, period, strconv.Itoa(limit))

	// match tracks on spotify
	var m []models.Track
	var u []models.Track

	// do this concurrently
	// guard := make(chan struct{}, 100)
	wg := sync.WaitGroup{}

	for i, t := range lt {

		if i == limit {
			log.Debug("Limit reached")
			break // break here
		}

		// guard <- struct{}{} // would block if guard channel is already filled

		wg.Add(1)

		go func(t models.Track) {
			st, err := spotifyclient.SearchTrack(pc, t.Artists[0].Name, t.Album.Name, t.Name)

			if err != nil {
				e = err
				wg.Done()
				return
			}

			// did we find something?
			if len(st) > 0 {
				mt := st[0]
				mt.Rank = t.Rank
				mt.PlayCount = t.PlayCount
				mt.SearchedFor = t.Artists[0].Name + " - " + t.Name
				m = append(m, mt)
			} else {
				u = append(u, t)
			}

			log.Debug("Searched for ", t.Artists[0].Name, " - ", t.Name)

			wg.Done()
			// <-guard
		}(t)
	}

	wg.Wait()

	sortByRank(m)
	sortByRank(u)

	// return matched and unmatched
	return &models.FoundTracks{
		Unmatched:      u,
		Matched:        m,
		Total:          len(u) + len(m),
		TotalMatched:   len(m),
		TotalUnmatched: len(u),
	}, e
}

func sortByRank(t []models.Track) {
	sort.Slice(t, func(i, j int) bool {
		ti, _ := strconv.Atoi(t[i].Rank)
		tj, _ := strconv.Atoi(t[j].Rank)
		return ti < tj
	})
}
