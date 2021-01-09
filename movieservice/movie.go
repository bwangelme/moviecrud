package movieservice

import (
	"moviedemo/moviedao"
	"time"
)

func Add(title string, pubdate time.Time, country string) error {
	movie := &moviedao.Movie{
		Title:   title,
		Pubdate: pubdate,
		Country: country,
	}
	return moviedao.Add(movie)
}

func Gets(movieIDs []int64) ([]*moviedao.Movie, error) {
	return moviedao.Gets(movieIDs)
}

func Get(movieID int64) (*moviedao.Movie, error) {
	return moviedao.Get(movieID)
}
