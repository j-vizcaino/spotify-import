package main

import (
	"fmt"
	"github.com/j-vizcaino/spotify-import/pkg/beets"
)

func main() {

	clt := beets.NewClient("http://nas.vco.mooo.com:8337")

	albums, err := clt.GetAlbums()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", albums)
}
