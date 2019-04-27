package beets

const (
	AlbumTypeAlbum       = "album"
	AlbumTypeCompilation = "compilation"
)

/*
   {
     "disctotal": 1,
     "albumstatus": "Official",
     "month": 2,
     "original_day": 26,
     "albumartist": "1208",
     "year": 2002,
     "albumdisambig": "",
     "albumartist_sort": "1208",
     "id": 312,
     "album": "Feedback Is Payback",
     "asin": "B00005YKIT",
     "script": "Latn",
     "mb_albumid": "ade985e9-0e08-47f6-802d-c85aa8f23b4c",
     "label": "Epitaph",
     "rg_album_gain": -10.5,
     "mb_releasegroupid": "3bffc5d1-548b-36a9-8489-b0ab67d28b16",
     "rg_album_peak": 1.118833,
     "albumartist_credit": "1208",
     "catalognum": "E-86638-2",
     "added": 1490176131.131017,
     "original_month": 2,
     "comp": 0,
     "genre": "Punk Rock",
     "day": 26,
     "original_year": 2002,
     "language": "eng",
     "r128_album_gain": 0,
     "mb_albumartistid": "d4e58ad8-500d-433d-b0ff-ee8fd3459ad9",
     "country": "US",
     "albumtype": "album"
   },
*/
type Album struct {
	Name   string `json:"album"`
	Type   string `json:"albumtype"`
	Artist string `json:"albumartist"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
	OriginalYear int `json:"original_year"`
}
