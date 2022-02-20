package dao

import (
	"database/sql"
	"douban/Struct"
	"douban/tool"
	"errors"
	"strings"
)

// FindMemberIMDB 根据name获取影人的IMDB
func FindMemberIMDB(db *sql.DB, name string) ([]string, error) {
	sqlStr := "select IMDb from person_info where name like ?"
	rows, err := db.Query(sqlStr, "%"+name+"%")
	tool.CheckErr(err)
	if err == sql.ErrNoRows {
		return nil, errors.New("NoMember")
	} else if err != nil {
		return nil, err
	}
	IMDBList := make([]string, 0)
	var IMDB string
	defer CloseDb(rows)
	for rows.Next() {
		err := rows.Scan(&IMDB)
		tool.CheckErr(err)
		if err != nil {
			return nil, err
		}
		IMDBList = append(IMDBList, IMDB)
	}
	return IMDBList, nil
}

// ShowMember 根据imdb查找影人信息
func ShowMember(db *sql.DB, IMDB string, heading string) (interface{}, error) {
	var sqlStr string
	switch heading {
	case "info":
		sqlStr = "select name, byName, date, IMDb, job, cons, male, place from person_info where IMDb = ?"
	case "view":
		sqlStr = "select video, IMDb, brief, pictureNum, picture_1, picture_2, picture_3, picture_4, picture_5 from person_view where IMDb = ?"
	}
	if len(IMDB) == 0 {
		return nil, errors.New("NoIMDB")
	}
	IMDB = "/" + IMDB
	IMDBList := make([]string, 0)
	var TStr = IMDB[1:]
	i1 := 0
	for {
		TStr = IMDB[i1+1:]
		i2 := strings.Index(TStr, "/")
		if i2 == 0 || i2 == -1 {
			IMDBList = append(IMDBList, TStr)
			break
		}
		TStr = TStr[:i2]
		IMDBList = append(IMDBList, TStr)
		i1 = i2 + i1 + 1
	}
	switch heading {
	case "info":
		var u0 Struct.PersonInfo
		var date []uint8
		u := make([]Struct.PersonInfo, 0)
		for _, i := range IMDBList {
			err := db.QueryRow(sqlStr, i).Scan(&u0.Name, &u0.ByName, &date, &u0.IMDB, &u0.Job, &u0.Cons, &u0.Male, &u0.Place)
			tool.CheckErr(err)
			if err != nil {
				return nil, err
			}
			u0.Date = string(date)
			u = append(u, u0)
		}
		return u, nil
	case "view":
		var u0 Struct.PersonView
		u := make([]Struct.PersonView, 0)
		var (
			Picture1ITE interface{}
			Picture2ITE interface{}
			Picture3ITE interface{}
			Picture4ITE interface{}
			Picture5ITE interface{}
			VideoITE    interface{}
		)
		for _, i := range IMDBList {
			err := db.QueryRow(sqlStr, i).Scan(&VideoITE, &u0.IMDB, &u0.Brief, &u0.PictureNum, &Picture1ITE, &Picture2ITE, &Picture3ITE, &Picture4ITE, &Picture5ITE)
			tool.CheckErr(err)
			if err != nil {
				return nil, err
			}
			u0.Picture1 = tool.TurnErr(Picture1ITE)
			u0.Picture2 = tool.TurnErr(Picture2ITE)
			u0.Picture3 = tool.TurnErr(Picture3ITE)
			u0.Picture4 = tool.TurnErr(Picture4ITE)
			u0.Picture5 = tool.TurnErr(Picture5ITE)
			u0.Video = tool.TurnErr(VideoITE)
			u = append(u, u0)
		}
		return u, nil
	}
	return nil, nil
}
