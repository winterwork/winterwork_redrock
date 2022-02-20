package service

import (
	"database/sql"
	"douban/dao"
	"douban/tool"
	"strings"
)

// FindMovies 查找电影
func FindMovies(db *sql.DB, name string, _type string) (interface{}, error) {
	var (
		u   interface{}
		err error
	)
	IMDB, err := dao.FindMovieIMdb(db, name)
	tool.CheckErr(err)
	if err != nil {
		return nil, err
	}
	switch _type {
	case "info":
		u, err = dao.FindMoviesInfo(db, IMDB)
	case "view":
		u, err = dao.FindMoviesView(db, IMDB)
	case "member":
		u, err = dao.FindMoviesMember(db, IMDB)
	case "liker":
		u, err = dao.FindMoviesLiker(db, IMDB)
	}
	tool.CheckErr(err)
	return u, err
}

// FindByIMDB 根据imdb查找电影
func FindByIMDB(db *sql.DB, IMDB string, _type string) (interface{}, error) {
	var (
		u   interface{}
		err error
	)
	var IMDBList = []string{IMDB}
	switch _type {
	case "info":
		u, err = dao.FindMoviesInfo(db, IMDBList)
	case "view":
		u, err = dao.FindMoviesView(db, IMDBList)
	case "member":
		u, err = dao.FindMoviesMember(db, IMDBList)
	case "liker":
		u, err = dao.FindMoviesLiker(db, IMDBList)
	}
	tool.CheckErr(err)
	return u, err
}

// GetRandMovie 获取指定数量的随机电影
func GetRandMovie(db *sql.DB, _type string, num int, heading string) (interface{}, error) {
	IMDbList, err := dao.GetMovie(db, _type)
	tool.CheckErr(err)
	if err != nil {
		return nil, err
	}
	var findNum int
	if len(IMDbList) < num || num == 0 {
		findNum = len(IMDbList)
	} else {
		findNum = num
	}
	movies, err := dao.GetRandInfo(db, findNum, IMDbList, heading)
	tool.CheckErr(err)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func FindByMIMDB(db *sql.DB, IMDB string) (string, error) {
	IMDBList, err := dao.FindByMIMDB(db, IMDB)
	tool.CheckErr(err)
	if err != nil {
		return "", err
	}
	var build strings.Builder
	for _, i := range IMDBList {
		build.WriteString(",")
		build.WriteString(i)
	}
	IMDBStr := build.String()
	return IMDBStr, nil
}
