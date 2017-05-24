package core

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
	Genres []string
	Images map[string]string
	Rating Rating
}

type Torrent struct {
	Provider string
	Filesize string
	Size int
	Peer int
	Seed int
	URL string
}

type Rating struct {
	Percentage int
	Watching int
	Votes int
	Loved int
	Hated int
}