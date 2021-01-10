package movieservice

import (
	"moviedemo/moviedao"
	"time"
)

func Add(title string, pubdate time.Time, country string) (int64, error) {
	movie := &moviedao.Movie{
		Title:   title,
		Pubdate: pubdate,
		Country: country,
	}
	return moviedao.Add(movie)
}

func Gets(movieIDs []int64) ([]*moviedao.Movie, error) {
	if len(movieIDs) == 0 {
		return make([]*moviedao.Movie, 0), nil
	}
	return moviedao.Gets(movieIDs)
}

func Get(movieID int64) (*moviedao.Movie, error) {
	// 电影找不到应该返回404,如果电影找不到也返回 error的话，那就没法简单区分内部错误(500)和找不的的错误(404)了
	// 此时为了区分可能需要增加一个 sentinel error，这样就让包的接口变宽了
	return moviedao.Get(movieID)
}

func List(start, limit int) (total int64, movieIDs []int64, err error) {
	return moviedao.List(start, limit)
}

func Delete(movieID int64) error {
	return moviedao.Delete(movieID)
}