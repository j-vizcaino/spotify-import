package spotify

import (
	"github.com/zmb3/spotify"
)

func SearchAlbums(clt spotify.Client, name string) ([]Album, error) {
	results, err := clt.Search(name, spotify.SearchTypeAlbum)
	if err != nil {
		return nil, err
	}
	out := make([]Album, 0)
	for _, result := range results.Albums.Albums {
		out = append(out, NewAlbum(result))
	}
	return out, nil
}
