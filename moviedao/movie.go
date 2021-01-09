package moviedao

import (
	"moviedemo/libs"
	"strings"
	"time"

	xerrors "github.com/pkg/errors"
)

type Movie struct {
	ID      int64
	Title   string
	Pubdate time.Time
	Country string
}

func Get(movieID int64) (*Movie, error) {
	movies, err := Gets([]int64{movieID})
	if err != nil {
		return nil, err
	}

	return movies[0], nil
}

func Gets(movieIDs []int64) ([]*Movie, error) {
	movies := make([]*Movie, 0)

	db, err := libs.GetMySQLDB()
	if err != nil {
		return movies, err
	}

	sql := `select id, title, pubdate, country from movie where id in (?` + strings.Repeat(", ?", len(movieIDs)-1) + `)`
	sqlArgs := make([]interface{}, len(movieIDs))
	for i, id := range movieIDs {
		sqlArgs[i] = id
	}

	movie := new(Movie)
	rows, err := db.Query(sql, sqlArgs...)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		return movies, xerrors.Wrapf(err, "query failed: %v, %v", sql, sqlArgs)
	}

	for rows.Next() {
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Pubdate, &movie.Country)
		if err != nil {
			return movies, xerrors.Wrapf(err, "scan failed: %v, %v", sql, sqlArgs)
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func Add(movie *Movie) error {
	db, err := libs.GetMySQLDB()
	if err != nil {
		return err
	}

	sql := `insert into movie(title, pubdate, country) values (?, ?, ?)`
	pubdate := movie.Pubdate.Format("2006-01-02 15:04:05")
	sqlArgs := []interface{}{
		movie.Title, pubdate, movie.Country,
	}
	_, err = db.Exec(sql, sqlArgs...)

	if err != nil {
		return xerrors.Wrapf(err, "exec sql:%v failed with args: %v", sql, sqlArgs)
	}

	return nil
}
