package main

import (
	"fmt"
	"github.com/j-vizcaino/spotify-import/pkg/beets"
	spot "github.com/j-vizcaino/spotify-import/pkg/spotify"
	"golang.org/x/text/unicode/norm"
	"sort"
	"strings"
	"unicode"
)

// normalizeName to keep only letters (with accents stripped), digits and spaces.
func normalizeName(in string) string {
	out := strings.Builder{}

	for _, r := range norm.NFD.String(in) {
		if unicode.IsLetter(r) {
			out.WriteRune(unicode.ToLower(r))
		} else if unicode.IsDigit(r) || unicode.IsSpace(r) {
			out.WriteRune(r)
		}
	}
	return out.String()
}

func findAlbumMatch(local beets.Album, candidates []spot.Album) int {

	localName := normalizeName(local.Name)
	localArtist := normalizeName(local.Artist)

	for idx, candidate := range candidates {
		if normalizeName(candidate.Name) == localName &&
			normalizeName(candidate.Artist) == localArtist {
			return idx
		}
	}
	return -1
}

func main() {
	spotClt, err := spot.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println("Getting Spotify library content...")
	//albums, err := spot.GetLibraryAlbums(spotClt)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("Library: %v", albums)
	//
	clt := beets.NewClient("http://nas.home:8337")

	localAlbums, err := clt.GetAlbums()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, album := range localAlbums {
		fmt.Printf("🔍  Looking for album %q, artist %q on Spotify...\n", album.Name, album.Artist)
		results, err := spot.SearchAlbums(spotClt, album.Name)
		if err != nil {
			fmt.Printf("Failed: %s\n", err)
			continue
		}
		idx := findAlbumMatch(album, results)
		if idx == -1 {
			strResults := make([]string, 0, len(results))
			for _, r := range results {
				strResults = append(strResults, fmt.Sprintf("%s - %s", r.Artist, r.Name))
			}
			sort.Strings(strResults)
			fmt.Printf("😭  No match found. Results were %v\n", strResults)
			continue
		}

		newAlbum := results[idx]
		fmt.Printf("🎉  Found! %+v\n", newAlbum)
	}
}
