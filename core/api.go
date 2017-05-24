package core

import (
	"strconv"
	"net/url"
)

const (
	SORT_LAST_ADDED = "last added"
	SORT_RATING = "rating"
	SORT_TITLE = "title"
	SORT_TRENDING = "trending"
	SORT_YEAR = "year"

	ORDER_NEWEST = -1
	ORDER_OLDEST = 1

	API_MOVIES = "/movies"
)

/*
	API is the base struct for talking to the API
 */
type API struct {
	Core Core
}

/**
	Movies will fetch movies from the API with the given parameters
 */
func (api API) Movies(page int, sort string, order int, genre string) (interface{}, error) {
	moviesURL := API_MOVIES + "/" + strconv.Itoa(page) + "?sort=" + url.QueryEscape(sort) + "&order=" + url.QueryEscape(strconv.Itoa(order))

	if genre != "" {
		moviesURL += "&genre=" + genre
	}

	var movies []Movie
	err := api.Core.Get(moviesURL, &movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}