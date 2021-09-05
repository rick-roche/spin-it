package spotifyclient

import (
	"regexp"
	"spin-it/internal/models"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/zmb3/spotify"
)

// SearchTrack todo
func SearchTrack(client spotifyClient, artist string, album string, track string) (tracks []models.Track, err error) {
	return searchRecursive(0, client, artist, album, track)
}

func searchRecursive(iteration int, client spotifyClient, artist string, album string, track string) (tracks []models.Track, err error) {
	iteration++
	var query, confidence = generateQuery(iteration, artist, album, track)

	if query == "" {
		return tracks, nil
	}

	found, err := search(client, query)

	if err != nil {
		return nil, err
	}

	if len(found) > 0 {
		tracks = found
		tracks[0].Confidence = confidence
		return tracks, nil
	}

	return searchRecursive(iteration, client, artist, album, track)
}

func generateQuery(iteration int, artist string, album string, track string) (query string, confidence string) {
	switch iteration {
	case 1:
		if len(album) > 0 {
			query = "artist:" + artist + " album:" + album + " track:" + removePunctuation(track)
		} else {
			query = "artist:" + artist + " track:" + removePunctuation(track)
		}
		confidence = "10"
	case 2:
		query = "artist:" + artist + " track:" + removeKeywords(track)
		confidence = "9"
	case 3:
		query = "artist:" + artist + " track:" + sanitise(track)
		confidence = "8"
	case 4:
		query = "artist:" + splitArtist(artist) + " track:" + removePunctuation(track)
		confidence = "7"
	case 5:
		query = "artist:" + artist + " track:" + removeBracketSections(track)
		confidence = "6"
	// case 6:
	// 	query = "track:" + removePunctuation(track)
	// 	confidence = "2"
	// case 7:
	// 	query = "track:" + removePunctuation(removeKeywords(track))
	// 	confidence = "1"

	default:
		query = ""
	}

	return query, confidence
}

func search(client spotifyClient, query string) (tracks []models.Track, err error) {
	res, err := client.Search(query, spotify.SearchTypeTrack)

	if err != nil {
		return nil, err
	}

	for _, track := range res.Tracks.Tracks {
		track := models.Track{
			ID:   track.SimpleTrack.ID.String(),
			Name: track.SimpleTrack.Name,
			Artists: []models.Artist{
				models.Artist{
					Name: track.SimpleTrack.Artists[0].Name,
				},
			},
			Album: models.Album{
				Name: track.Album.Name,
			},
		}

		tracks = append(tracks, track)
	}

	return tracks, nil
}

func sanitise(s string) string {
	return removePunctuation(removeKeywords(s))
}

func removePunctuation(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9: -]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(s, "")
}

func removeKeywords(s string) string {
	// reg, err := regexp.Compile(`(?i)mix`)
	reg, err := regexp.Compile(`(?i)mix|(?i)remix|ft.|the`)
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(s, "")
}

func removeBracketSections(s string) string {
	reg, err := regexp.Compile(`[\(\[][^)]*[\)\]]`)
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(s, "")
}

func splitArtist(s string) string {
	reg, err := regexp.Compile(`&`)
	if err != nil {
		log.Fatal(err)
	}
	a := reg.Split(s, -1)

	return strings.TrimSpace(a[0])
}
