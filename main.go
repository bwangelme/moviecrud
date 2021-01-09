package main

import (
	"fmt"
	"log"
	"moviedemo/movieservice"
	"time"
)

func Add() {
	pubDate, _ := time.Parse("2006-01-02", "1994-09-10")
	err := movieservice.Add("肖申克的救赎", pubDate, "美国")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	movies, err := movieservice.Gets([]int64{1})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(len(movies), movies[0])
}
