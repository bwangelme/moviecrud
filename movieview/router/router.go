package router

import (
	"fmt"
	"moviedemo/movieview"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	r := gin.New()

	r.Use(gin.Recovery())

	r.GET("/ping", func(context *gin.Context) {
		fmt.Fprintf(context.Writer, "pong")
		return
	})

	api := r.Group("/api")
	api.Use(func(c *gin.Context) {
		c.Next()
		//TODO: 如果 error 有多层，那么或生成一个 error 数组，应该处理成只在一个error中返回
		if len(c.Errors) > 0 {
			c.JSON(http.StatusBadRequest, c.Errors)
		}
	})
	{
		api.GET("movie/:id", movieview.GetMovie)
		api.DELETE("movie/:id", movieview.DeleteMovie)
	}
	{
		api.GET("movies/list", movieview.ListMovies)
		api.POST("movies", movieview.AddMovie)
	}

	return r
}
