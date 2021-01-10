package moviedao

import (
	"log"
	"moviedemo/libs"
	"moviedemo/movieerror"
	"strings"
	"time"

	xerrors "github.com/pkg/errors"
)

type Movie struct {
	ID      int64     `json:"id"`
	Title   string    `json:"title"`
	Pubdate time.Time `json:"pubdate"`
	Country string    `json:"country"`
}

func Get(movieID int64) (*Movie, error) {
	movies, err := Gets([]int64{movieID})
	if err != nil {
		return nil, err
	}

	if len(movies) == 0 {
		return nil, xerrors.Wrapf(movieerror.MovieNotFound, "%d not found", movieID)
	}

	return movies[0], nil
}

func Gets(movieIDs []int64) ([]*Movie, error) {
	movies := make([]*Movie, 0)

	db, err := libs.GetMySQLDB()
	if err != nil {
		return movies, err
	}

	sql := `select id, title, pubdate, country from movie 
	where id in (?` + strings.Repeat(", ?", len(movieIDs)-1) + `)
`
	sqlArgs := make([]interface{}, len(movieIDs))
	for i, id := range movieIDs {
		sqlArgs[i] = id
	}

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
		// 注意由于 movies 中存的是指针，所以 movie 要在每个循环中重新定义
		var movie = new(Movie)
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Pubdate, &movie.Country)
		if err != nil {
			return movies, xerrors.Wrapf(err, "scan failed: %v, %v", sql, sqlArgs)
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func Add(movie *Movie) (int64, error) {
	db, err := libs.GetMySQLDB()
	if err != nil {
		return 0, err
	}

	sql := `insert into movie(title, pubdate, country) values (?, ?, ?)`
	pubdate := movie.Pubdate.Format("2006-01-02 15:04:05")
	sqlArgs := []interface{}{
		movie.Title, pubdate, movie.Country,
	}
	res, err := db.Exec(sql, sqlArgs...)
	if err != nil {
		return 0, xerrors.Wrapf(err, "exec sql:%v failed with args: %v", sql, sqlArgs)
	}

	movieID, err := res.LastInsertId()
	if err != nil {
		return 0, xerrors.Wrapf(err, "get last insert id failed")
	}
	return movieID, nil
}

func List(start, limit int) (total int64, movieIDs []int64, err error) {
	movieIDs = make([]int64, 0)
	db, err := libs.GetMySQLDB()
	if err != nil {
		return total, movieIDs, err
	}

	sql := `select count(*) from movie`
	row := db.QueryRow(sql)
	if err := row.Scan(&total); err != nil {
		return total, movieIDs, xerrors.Wrapf(err, "scan failed: %v", sql)
	}

	sql = `select id from movie limit ?,?`
	sqlArgs := []interface{}{
		start, limit,
	}
	rows, err := db.Query(sql, sqlArgs...)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		return total, movieIDs, xerrors.Wrapf(err, "query failed: %v, %v", sql, sqlArgs)
	}

	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			return total, movieIDs, xerrors.Wrapf(err, "scan failed: %v, %v", sql, sqlArgs)
		}
		movieIDs = append(movieIDs, id)
	}

	return total, movieIDs, nil
}

func Delete(movieID int64) error {
	db, err := libs.GetMySQLDB()
	if err != nil {
		return err
	}

	sql := `delete from movie where id = ?`
	sqlArgs := movieID

	res, err := db.Exec(sql, sqlArgs)
	if err != nil {
		return xerrors.Wrapf(err, "delete failed: %v, %v", sql, sqlArgs)
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return xerrors.Wrapf(err, "get row affected failed: %v, %v", sql, sqlArgs)
	}

	log.Printf("Delete row id=%d, affected rows %d\n", movieID, affectedRows)

	return nil
}
