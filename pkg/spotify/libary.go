package spotify

import (
	"github.com/zmb3/spotify"
)

func GetLibraryAlbums(client spotify.Client) ([]Album, error) {
	out := make([]Album, 0, 512)
	offset := 0

	for {
		opts := spotify.Options{
			Offset: &offset,
		}
		page, err := client.CurrentUsersAlbumsOpt(&opts)
		if err != nil {
			return nil, err
		}

		offset += len(page.Albums)
		for _, album := range page.Albums {
			a := NewAlbum(album.SimpleAlbum)
			a.TrackCount = album.Tracks.Total
			out = append(out, a)
		}

		if offset >= page.Total {
			break
		}
	}
	return out, nil
}
