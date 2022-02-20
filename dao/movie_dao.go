package dao

import (
	"database/sql"
	"douban/Struct"
	"douban/tool"
	"errors"
	"math/rand"
	"time"
)

// FindMoviesInfo 获取指定名字电影的info
func FindMoviesInfo(db *sql.DB, IMDB []string) ([]Struct.MovieInfo, error) {
	sqlStr := "select IMDb, name, alias, date, `long`, type from movies_info where IMDb = ?"
	u := make([]Struct.MovieInfo, 0)
	var u0 Struct.MovieInfo
	var Date []uint8
	var AliasITE interface{}
	for _, i := range IMDB {
		row := db.QueryRow(sqlStr, i)
		err := row.Scan(&u0.IMDB, &u0.Name, &AliasITE, &Date, &u0.Long, &u0.Type)
		u0.Date = string(Date)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			tool.CheckErr(err)
			return nil, err
		}
		u0.Alias = tool.TurnErr(AliasITE)
		u = append(u, u0)
	}
	return u, nil
}

// FindMoviesView 获取指定名字的电影view
func FindMoviesView(db *sql.DB, IMDB []string) ([]Struct.MovieView, error) {
	u := make([]Struct.MovieView, 0)
	var u0 Struct.MovieView
	sqlStr := "select IMDb, brief, pictureNum, picture_1, picture_2, picture_3, picture_4, picture_5, video from movie_view where IMDB = ?"
	var (
		Picture1ITE interface{}
		Picture2ITE interface{}
		Picture3ITE interface{}
		Picture4ITE interface{}
		Picture5ITE interface{}
	)
	for _, i := range IMDB {
		err := db.QueryRow(sqlStr, i).Scan(&u0.IMDB, &u0.Brief, &u0.PictureNum, &Picture1ITE, &Picture2ITE, &Picture3ITE, &Picture4ITE, &Picture5ITE, &u0.Video)
		tool.CheckErr(err)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		u0.Picture1 = tool.TurnErr(Picture1ITE)
		u0.Picture2 = tool.TurnErr(Picture2ITE)
		u0.Picture3 = tool.TurnErr(Picture3ITE)
		u0.Picture4 = tool.TurnErr(Picture4ITE)
		u0.Picture5 = tool.TurnErr(Picture5ITE)
		u = append(u, u0)
	}
	return u, nil
}

// FindMoviesMember 获取指定IMDB电影的member
func FindMoviesMember(db *sql.DB, IMDB []string) ([]Struct.MovieMember, error) {
	u := make([]Struct.MovieMember, 0)
	var u0 Struct.MovieMember
	sqlStr := "select IMDb, director, scriptwriter, player from movie_member where IMDb = ?"
	for _, i := range IMDB {
		err := db.QueryRow(sqlStr, i).Scan(&u0.IMDB, &u0.Director, &u0.Scriptwriter, &u0.Player)
		tool.CheckErr(err)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		u = append(u, u0)
	}
	return u, nil
}

// FindMoviesLiker 获取指定IMDB电影的liker
func FindMoviesLiker(db *sql.DB, IMDB []string) ([]Struct.MovieLiker, error) {
	u := make([]Struct.MovieLiker, 0)
	var u0 Struct.MovieLiker
	sqlStr := "select IMDb, score, num_1, num_2, num_3, num_4, num_5, num_want, num_done, commentnum, essay_num from movie_liker where IMDb = ?"
	for _, i := range IMDB {
		err := db.QueryRow(sqlStr, i).Scan(&u0.IMDB, &u0.Score, &u0.Num1, &u0.Num2, &u0.Num3, &u0.Num4, &u0.Num5, &u0.NumWant, &u0.NumDone, &u0.CommentNum, &u0.EssayNum)
		tool.CheckErr(err)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		u = append(u, u0)
	}
	return u, nil
}

// GetMovie 获取指定类型的电影IMDb
func GetMovie(db *sql.DB, _type string) ([]string, error) {
	IMDbList := make([]string, 0)
	sqlStr := "select IMDb from movies_info where type like ?"
	rows, err := db.Query(sqlStr, "%"+_type+"%")
	tool.CheckErr(err)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	var IMDb string
	defer CloseDb(rows)
	for rows.Next() {
		err = rows.Scan(&IMDb)
		tool.CheckErr(err)
		if err != nil {
			return nil, err
		}
		IMDbList = append(IMDbList, IMDb)
	}
	return IMDbList, err
}

// FindMovieIMdb 获取指定名字的电影IMDb
func FindMovieIMdb(db *sql.DB, name string) ([]string, error) {
	sqlStr := "select IMDb from movies_info where name like ?"
	rows, err := db.Query(sqlStr, "%"+name+"%")
	tool.CheckErr(err)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	var IMDBTem string
	IMDB := make([]string, 0)
	defer CloseDb(rows)
	for rows.Next() {
		err = rows.Scan(&IMDBTem)
		tool.CheckErr(err)
		if err != nil {
			return nil, err
		}
		IMDB = append(IMDB, IMDBTem)
	}
	return IMDB, err
}

// GetRandInfo 获取电影信息
func GetRandInfo(db *sql.DB, num int, IMDBList []string, heading string) (interface{}, error) {
	sum := len(IMDBList)
	randList := make([]int, num)
	var flag bool
	if num != sum {
		for i := 0; i < num; i++ {
			flag = true
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			tem := r.Intn(sum)
			var t int
			for t = 0; t < i; t++ {
				if randList[t] == tem {
					i--
					flag = false
					break
				}
			}
			if flag {
				randList[i] = tem
			}
		}
	} else {
		for i := 0; i < num; i++ {
			randList[i] = i
		}
	}
	IMDB := make([]string, 0)
	for i := 0; i < num; i++ {
		IMDB = append(IMDB, IMDBList[randList[i]])
	}
	switch heading {
	case "info":
		MovieLists, err := FindMoviesInfo(db, IMDB)
		tool.CheckErr(err)
		return MovieLists, err
	case "view":
		MovieLists, err := FindMoviesView(db, IMDB)
		tool.CheckErr(err)
		return MovieLists, err
	case "member":
		MovieLists, err := FindMoviesMember(db, IMDB)
		tool.CheckErr(err)
		return MovieLists, err
	case "liker":
		MovieLists, err := FindMoviesLiker(db, IMDB)
		tool.CheckErr(err)
		return MovieLists, err
	}
	return nil, nil
}

func FindByMIMDB(db *sql.DB, IMDB string) ([]string, error) {
	sqlStr := "select IMDb from movie_member where director like ? or player like ? or scriptwriter like ?"
	rows, err := db.Query(sqlStr, "%"+IMDB+"%", "%"+IMDB+"%", "%"+IMDB+"%")
	tool.CheckErr(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("NoRows")
		}
		return nil, err
	}
	var MemberIMDB string
	IMDBList := make([]string, 0)
	for rows.Next() {
		err = rows.Scan(&MemberIMDB)
		tool.CheckErr(err)
		if err != nil {
			return nil, err
		}
		IMDBList = append(IMDBList, MemberIMDB)
	}
	return IMDBList, nil
}
