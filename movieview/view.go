package movieview

import (
	"fmt"
	"log"
	"moviedemo/movieerror"
	"moviedemo/movieservice"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func AddMovie(ctx *gin.Context) {
	var err error
	var args struct {
		Title   string `form:"title"`
		Pubdate string `form:"pubdate"`
		Country string `form:"country"`
	}
	err = ctx.Bind(&args)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	pubDate, err := time.Parse("2006-01-02", args.Pubdate)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id, err := movieservice.Add(args.Title, pubDate, args.Country)
	if err != nil {
		log.Printf("add movie error %+v\n", err)
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("add movie failed"))
		return
	}
	movie, err := movieservice.Get(id)
	if err != nil {
		log.Printf("get movie error %+v\n", err)
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("get movie failed"))
		return
	}
	ctx.JSON(200, gin.H{
		"error": nil,
		"movie": movie,
	})
}

func GetMovie(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	movie, err := movieservice.Get(id)
	if err != nil {
		if errors.Is(err, movieerror.MovieNotFound) {
			ctx.AbortWithError(404, err)
			return
		}
		log.Printf("get movie error %+v\n", err)
		ctx.AbortWithError(500, fmt.Errorf("get movie failed"))
		return
	}
	ctx.JSON(200, gin.H{
		"error": nil,
		"movie": movie,
	})
}

func ListMovies(ctx *gin.Context) {
	var args struct {
		Start int `form:"start"`
		Limit int `form:"limit"`
	}
	// 设置请求参数的默认值
	args.Start = 0
	args.Limit = 20
	err := ctx.BindQuery(&args)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	total, movieIDs, err := movieservice.List(args.Start, args.Limit)
	if err != nil {
		log.Printf("list movie error %+v\n", err)
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("list movie failed"))
		return
	}
	movies, err := movieservice.Gets(movieIDs)
	if err != nil {
		log.Printf("gets movie error %+v\n", err)
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("gets movie failed"))
		return
	}
	ctx.JSON(200, gin.H{
		"error":  nil,
		"total":  total,
		"movies": movies,
	})
}

func DeleteMovie(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	err = movieservice.Delete(id)
	if err != nil {
		log.Printf("delete movie error %+v\n", err)
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("delete movie failed"))
		return
	}

	ctx.JSON(200, gin.H{
		"error": nil,
	})
}
