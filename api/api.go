package api

const (
	LAST_ADDED = "last added"
	RATING = "rating"
	TITLE = "title"
	TRENDING = "trending"
	YEAR = "year"
	NEWEST = -1
	OLDEST = 1
)

type API struct {

}

func (api API) Movies(page int, sort string, order int, genre string) interface{} {
	core.Core{}
}