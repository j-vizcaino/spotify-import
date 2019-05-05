package spotify

import (
	"github.com/zmb3/spotify"
	"strconv"
	"strings"
)

type Album struct {
	ID         spotify.ID
	Name       string
	Artist     string
	Type       string
	Year       int
	TrackCount int
}

func NewAlbum(a spotify.SimpleAlbum) Album {
	dateParts := strings.Split(a.ReleaseDate, "-")
	year, err := strconv.ParseInt(dateParts[0], 10, 32)
	if err != nil {
		year = 0
	}
	return Album{
		ID:     a.ID,
		Name:   a.Name,
		Artist: a.Artists[0].Name,
		Type:   a.AlbumType,
		Year:   int(year),
	}
}
