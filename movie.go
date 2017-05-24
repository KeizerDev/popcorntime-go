package main

type Movie struct {
	ID string `json:"_id"`
	IMDB_ID string
	Title string
	Year string
	Synopsis string
	Runtime string
	Released int32
	Trailer string
	Certification string
	Torrents map[string]map[string]Torrent
}

type Torrent struct {
	Provider string
	Filesize string
	Size int
	Peer int
	Seed int
	URL string
}