package dao

import (
	"database/sql"
	"douban/Struct"
	"douban/tool"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// InsertMsg 添加留言
func InsertMsg(db *sql.DB, username string, msg string, pid int, movie string) (int, error) {
	sqlStr := "insert into info(username, msg, pid, time, type, movie) values (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, username, msg, pid, time.Now().Unix(), 1, movie)
	tool.CheckErr(err)
	if err != nil {
		return 0, err
	}
	if pid != 0 {
		commentAdd(db, pid)
	}
	err = ChangePoint(db, movie, 0, 1, username)
	tool.CheckErr(err)
	if err != nil {
		return 0, err
	}
	var id int
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)
	tool.CheckErr(err)
	return id, err
}

// FindTheMsg 查找留言或评论
func FindTheMsg(db *sql.DB, id string) (Struct.Info, error) {
	sqlStr := "select id, username, msg, pid, commentNum, time, thumbs_up, liker, movie, point, essay, type from info where id = ?"
	row := db.QueryRow(sqlStr, id)
	var u Struct.Info
	err := row.Scan(&u.Id, &u.Username, &u.Msg, &u.Pid, &u.CommentNum, &u.Time, &u.Thumb, &u.Liker, &u.Movie, &u.Point, &u.Essay, &u.Type)
	tool.CheckErr(err)
	if err == sql.ErrNoRows {
		return u, errors.New("NoMessage")
	}
	return u, nil
}

// CloseDb 关闭数据库
func CloseDb(row *sql.Rows) {
	err := row.Close()
	tool.CheckErr(err)
}

// commentAdd 添加评论数
func commentAdd(db *sql.DB, pid int) {
	var commentNum int
	sqlStr := "select commentNum from info where id = ?"
	row := db.QueryRow(sqlStr, pid)
	err := row.Scan(&commentNum)
	tool.CheckErr(err)
	commentNum++
	sqlStr2 := "update info set commentNum = ? where id = ?"
	_, err = db.Exec(sqlStr2, commentNum, pid)
	tool.CheckErr(err)
}

// DeleteMsg 删除留言
func DeleteMsg(db *sql.DB, id string) error {
	sqlStr := "delete from info where id = ?"
	idNum, _ := strconv.Atoi(id)
	_, err := db.Exec(sqlStr, idNum)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	sqlStr = "select id from info where pid = ?"
	var sid int
	rows, err := db.Query(sqlStr, idNum)
	tool.CheckErr(err)
	if err == sql.ErrNoRows {
		return nil
	}
	for rows.Next() {
		err = rows.Scan(&sid)
		tool.CheckErr(err)
		if err != nil {
			return err
		}
		err = DeleteMsg(db, strconv.Itoa(sid))
		if err != nil {
			return err
		}
	}
	tool.CheckErr(err)
	return err
}

// FindAll 查找所有留言或评论
func FindAll(db *sql.DB, pid int, movie string) ([]Struct.Info, error) {
	sqlStr := "select id, username, msg, pid, commentNum, time, thumbs_up, liker, movie, point, essay, type from info where movie = ?"
	rows, err := db.Query(sqlStr, movie)
	defer CloseDb(rows)
	var u0 Struct.Info
	u := make([]Struct.Info, 0)
	for rows.Next() {
		err = rows.Scan(&u0.Id, &u0.Username, &u0.Msg, &u0.Pid, &u0.CommentNum, &u0.Time, &u0.Thumb, &u0.Liker, &u0.Movie, &u0.Point, &u0.Essay, &u0.Type)
		tool.CheckErr(err)
		if err != nil {
			return nil, err
		}
		if u0.Pid == pid {
			u = append(u, u0)
		}
	}
	return u, nil
}

// ListComments 展示评论
func ListComments(db *sql.DB, pid int, u []Struct.Info) ([]Struct.Info, error) {
	var u0 Struct.Info
	sqlStr := "select id, username, msg, pid, commentNum, time, thumbs_up, liker, movie, point, essay, type from info where pid = ?"
	rows, err := db.Query(sqlStr, pid)
	tool.CheckErr(err)
	defer CloseDb(rows)
	if err == sql.ErrNoRows {
		return nil, errors.New("NoMessage")
	} else if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&u0.Id, &u0.Username, &u0.Msg, &u0.Pid, &u0.CommentNum, &u0.Time, &u0.Thumb, &u0.Liker, &u0.Movie, &u0.Point, &u0.Essay, &u0.Type)
		tool.CheckErr(err)
		if err != nil {
			return nil, err
		}
		u = append(u, u0)
		var commentNum int
		sqlStr = "select commentNum from info where id = ?"
		err = db.QueryRow(sqlStr, u0.Id).Scan(&commentNum)
		tool.CheckErr(err)
		if err != nil {
			return nil, err
		}
		if commentNum != 0 {
			u, err = ListComments(db, u0.Id, u)
			if err != nil {
				return nil, err
			}
		}
	}
	return u, nil
}

// ThumbAdd 为评论添加点赞
func ThumbAdd(db *sql.DB, id string, username string) error {
	idNum, _ := strconv.Atoi(id)
	sqlStr := "select thumbs_up, liker from info where id = ?"
	row := db.QueryRow(sqlStr, idNum)
	var thump int
	var liker string
	err := row.Scan(&thump, &liker)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	if liker != "" {
		var TLiker = liker[1:]
		i1 := 0
		for {
			TLiker = liker[i1+1:]
			i2 := strings.Index(TLiker, ",")
			if i2 == 0 || i2 == -1 {
				if TLiker == username {
					return errors.New("done")
				}
				break
			}
			TLiker = TLiker[:i2]
			if TLiker == username {
				return errors.New("done")
			}
			i1 = i2 + i1 + 1
		}
	}
	thump++
	liker = liker + "," + username
	sqlStr = "update info set thumbs_up = ?, liker = ? where id = ?"
	_, err = db.Exec(sqlStr, thump, liker, id)
	tool.CheckErr(err)
	return err
}

// InsertEssay 添加留言
func InsertEssay(db *sql.DB, username string, msg string, movie string, point int) (int, error) {
	sqlStr := "insert into info(username, essay, time, type, movie, point, pid) values (?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, username, msg, time.Now().UnixNano(), 2, movie, point, 0)
	tool.CheckErr(err)
	if err != nil {
		return 0, err
	}
	err = ChangePoint(db, movie, point, 2, username)
	tool.CheckErr(err)
	if err != nil {
		return 0, err
	}
	var id int
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)
	tool.CheckErr(err)
	return id, err
}

// ChangePoint 调整评分
func ChangePoint(db *sql.DB, movie string, point int, _type int, username string) error {
	var sqlStr string
	switch _type {
	case 2:
		var num, essayNum int
		var score float32
		switch point {
		case 1:
			sqlStr = "select score, num_1, essay_num from movie_liker where IMDb = ?"
		case 2:
			sqlStr = "select score, num_2, essay_num from movie_liker where IMDb = ?"
		case 3:
			sqlStr = "select score, num_3, essay_num from movie_liker where IMDb = ?"
		case 4:
			sqlStr = "select score, num_4, essay_num from movie_liker where IMDb = ?"
		case 5:
			sqlStr = "select score, num_5, essay_num from movie_liker where IMDb = ?"
		}
		row := db.QueryRow(sqlStr, movie)
		err := row.Scan(&score, &num, &essayNum)
		tool.CheckErr(err)
		if err != nil {
			return err
		}
		num++
		essayNum++
		score = (score*float32(essayNum) + float32(point*2)) / (float32(essayNum) + 1)
		switch point {
		case 1:
			sqlStr = "update movie_liker set score = ?,essay_Num = ?,num_1 = ? where IMDb = ?"
		case 2:
			sqlStr = "update movie_liker set score = ?,essay_Num = ?,num_2 = ? where IMDb = ?"
		case 3:
			sqlStr = "update movie_liker set score = ?,essay_Num = ?,num_3 = ? where IMDb = ?"
		case 4:
			sqlStr = "update movie_liker set score = ?,essay_Num = ?,num_4 = ? where IMDb = ?"
		case 5:
			sqlStr = "update movie_liker set score = ?,essay_Num = ?,num_5 = ? where IMDb = ?"
		}
		_, err = db.Exec(sqlStr, score, essayNum, num, movie)
		tool.CheckErr(err)
		if err != nil {
			return err
		}
	case 1:
		sqlStr = "select commentNum from movie_liker where IMDb = ?"
		row := db.QueryRow(sqlStr, movie)
		var num int
		err := row.Scan(&num)
		tool.CheckErr(err)
		if err != nil {
			return err
		}
		num++
		sqlStr = "update movie_liker set commentNum = ? where IMDb = ?"
		_, err = db.Exec(sqlStr, num, movie)
		tool.CheckErr(err)
		if err != nil {
			return err
		}
	}
	if username != "anonymousUser" {
		var ID int
		sqlStr = "select ID from user where username = ?"
		err := db.QueryRow(sqlStr, username).Scan(&ID)
		tool.CheckErr(err)
		if err != nil {
			return err
		}
		var num int
		sqlStr = "select commentNum from user_detail where ID = ?"
		err = db.QueryRow(sqlStr, ID).Scan(&num)
		tool.CheckErr(err)
		if err != nil {
			return err
		}
		num++
		sqlStr = "update user_detail set commentNum = ? where ID = ?"
		_, err = db.Exec(sqlStr, num, ID)
		tool.CheckErr(err)
		if err != nil {
			return err
		}
	}
	return nil
}

// MoviePointChange 改变电影想看/看过
func MoviePointChange(db *sql.DB, _type string, movie string) error {
	var sqlStr string
	switch _type {
	case "W":
		sqlStr = "select num_want from movie_liker where IMDb = ?"
	case "D":
		sqlStr = "select num_done from movie_liker where IMDb = ?"
	}
	var num int
	err := db.QueryRow(sqlStr, movie).Scan(&num)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	num++
	switch _type {
	case "W":
		sqlStr = "update movie_liker set num_want = ? where IMDb = ?"
	case "D":
		sqlStr = "update movie_liker set num_done = ? where IMDb = ?"
	}
	_, err = db.Exec(sqlStr, num, movie)
	tool.CheckErr(err)
	return err
}

// UserPointChange 改变用户想看/看过
func UserPointChange(db *sql.DB, username string, _type string, movie string) error {
	var (
		sqlStr string
		ID     int
	)
	sqlStr = "select ID from user where username = ?"
	err := db.QueryRow(sqlStr, username).Scan(&ID)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	switch _type {
	case "W":
		sqlStr = "select UNum from user_detail where ID = ?"
	case "D":
		sqlStr = "select MNum from user_detail where ID = ?"
	}
	var num int
	err = db.QueryRow(sqlStr, ID).Scan(&num)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	num++
	switch _type {
	case "W":
		sqlStr = "update user_detail set UNum = ? where ID = ?"
	case "D":
		sqlStr = "update user_detail set MNum = ? where ID = ?"
	}
	_, err = db.Exec(sqlStr, num, ID)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	switch _type {
	case "W":
		sqlStr = "select MovieWant from user_detail where ID = ?"
	case "D":
		sqlStr = "select MovieDone from user_detail where ID = ?"
	}
	var str string
	err = db.QueryRow(sqlStr, ID).Scan(&str)
	tool.CheckErr(err)
	var build strings.Builder
	build.WriteString(str)
	build.WriteString("/")
	build.WriteString(movie)
	s := build.String()
	switch _type {
	case "W":
		sqlStr = "update user_detail set MovieWant = ? where ID = ?"
	case "D":
		sqlStr = "update user_detail set MovieDone = ? where ID = ?"
	}
	_, err = db.Exec(sqlStr, s, ID)
	tool.CheckErr(err)
	return err
}

// JudgeWOD 判断是否想看/看过
func JudgeWOD(db *sql.DB, username string, _type string, movie string) bool {
	var str, sqlStr string
	ID, err := GetId(db, username)
	tool.CheckErr(err)
	if err != nil {
		return false
	}
	switch _type {
	case "W":
		sqlStr = "select MovieWant from user_detail where ID = ?"
	case "D":
		sqlStr = "select MovieDone from user_detail where ID = ?"
	}
	err = db.QueryRow(sqlStr, ID).Scan(&str)
	tool.CheckErr(err)
	if err != nil {
		return false
	}
	if len(str) == 0 {
		return true
	}
	var TStr = str[1:]
	i1 := 0
	for {
		TStr = str[i1+1:]
		i2 := strings.Index(TStr, "/")
		if i2 == 0 || i2 == -1 {
			if TStr == movie {
				return false
			}
			break
		}
		TStr = TStr[:i2]
		if TStr == movie {
			return false
		}
		i1 = i2 + i1 + 1
	}
	return true
}

func DeleteThumb(db *sql.DB, username string, ID int) error {
	var (
		thumbUp int
		liker   string
	)
	sqlStr := "select thumbs_up, liker from info where id = ?"
	err := db.QueryRow(sqlStr, ID).Scan(&thumbUp, &liker)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	if len(liker) == 0 {
		return errors.New("NoThumb")
	}
	var TLiker = liker[1:]
	i1 := 0
	for {
		TLiker = liker[i1+1:]
		i2 := strings.Index(TLiker, ",")
		if i2 == 0 || i2 == -1 {
			if TLiker == username {
				liker = liker[:i1]
			} else {
				return errors.New("NoThumb")
			}
			break
		}
		TLiker = TLiker[:i2]
		if TLiker == username {
			var build strings.Builder
			build.WriteString(liker[:i1])
			build.WriteString(liker[i2+i1+1:])
			liker = build.String()
			break
		}
		i1 = i2 + i1 + 1
	}
	thumbUp--
	sqlStr = "update info set liker = ?, thumbs_up = ? where id = ?"
	_, err = db.Exec(sqlStr, liker, thumbUp, ID)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	return nil
}

func DeletePoint(db *sql.DB, username string, movie string, _type string) error {
	ID, err := GetId(db, username)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	var (
		sqlStr string
		str    string
		num    int
	)
	switch _type {
	case "W":
		sqlStr = "select MovieWant,UNum from user_detail where ID = ?"
	case "D":
		sqlStr = "select MovieDone,MNum from user_detail where ID = ?"
	}
	err = db.QueryRow(sqlStr, ID).Scan(&str, &num)
	if str == "" {
		return errors.New("NoPoint")
	}
	var TStr = str[1:]
	i1 := 0
	for {
		TStr = str[i1+1:]
		i2 := strings.Index(TStr, "/")
		if i2 == 0 || i2 == -1 {
			fmt.Println(TStr)
			if TStr == movie {
				str = str[:i1]
			} else {
				return errors.New("NoPoint")
			}
			break
		}
		TStr = TStr[:i2]
		if TStr == movie {
			var build strings.Builder
			build.WriteString(str[:i1])
			build.WriteString(str[i2+i1+1:])
			str = build.String()
			break
		}
		i1 = i2 + i1 + 1
	}
	num--
	switch _type {
	case "W":
		sqlStr = "update user_detail set MovieWant = ?,UNum = ? where ID = ?"
	case "D":
		sqlStr = "update user_detail set MovieDone = ?,MNum = ? where ID = ?"
	}
	_, err = db.Exec(sqlStr, str, num, ID)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	switch _type {
	case "W":
		sqlStr = "select num_want from movie_liker where IMDb = ?"
	case "D":
		sqlStr = "select num_done from movie_liker where IMDb = ?"
	}
	err = db.QueryRow(sqlStr, movie).Scan(&num)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	num--
	switch _type {
	case "W":
		sqlStr = "update movie_liker set num_want = ? where IMDb = ?"
	case "D":
		sqlStr = "update movie_liker set num_done = ? where IMDb = ?"
	}
	_, err = db.Exec(sqlStr, num)
	tool.CheckErr(err)
	return err
}

func ThumbUser(db *sql.DB, user string, username string) error {
	sqlStr := "select attWho from user_detail where ID = ?"
	var (
		num    int
		attWho string
	)
	ID, err := GetId(db, username)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	err = db.QueryRow(sqlStr, ID).Scan(&attWho)
	if attWho != "" {
		var TStr = attWho[1:]
		i1 := 0
		for {
			TStr = attWho[i1+1:]
			i2 := strings.Index(TStr, ",")
			if i2 == 0 || i2 == -1 {
				if TStr == user {
					return errors.New("done")
				}
				break
			}
			TStr = TStr[:i2]
			if TStr == user {
				return errors.New("done")
			}
			i1 = i2 + i1 + 1
		}
	}
	var build strings.Builder
	build.WriteString(attWho)
	build.WriteString(",")
	build.WriteString(user)
	attWho = build.String()
	sqlStr = "update user_detail set attWho = ? where ID = ?"
	_, err = db.Exec(sqlStr, attWho, ID)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	sqlStr = "select thumbNum from user_detail where ID = ?"
	ID, _ = strconv.Atoi(user)
	tool.CheckErr(err)
	if err != nil {
		fmt.Println(1)
		return err
	}
	err = db.QueryRow(sqlStr, ID).Scan(&num)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	num++
	sqlStr = "update user_detail set thumbNum = ? where ID = ?"
	_, err = db.Exec(sqlStr, num, ID)
	tool.CheckErr(err)
	return err
}

func DeleteUT(db *sql.DB, user string, username string) error {
	var (
		attWho string
		num    int
	)
	sqlStr := "select attWho from user_detail where ID = ?"
	ID, err := GetId(db, username)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	err = db.QueryRow(sqlStr, ID).Scan(&attWho)
	if len(attWho) == 0 {
		return errors.New("NoThumb")
	}
	var TStr = attWho[1:]
	i1 := 0
	for {
		TStr = attWho[i1+1:]
		i2 := strings.Index(TStr, ",")
		if i2 == 0 || i2 == -1 {
			if TStr == user {
				attWho = attWho[:i1]
			} else {
				return errors.New("NoThumb")
			}
			break
		}
		TStr = TStr[:i2]
		if TStr == user {
			var build strings.Builder
			build.WriteString(attWho[:i1])
			build.WriteString(attWho[i2+i1+1:])
			attWho = build.String()
			break
		}
		i1 = i2 + i1 + 1
	}
	sqlStr = "update user_detail set attWho = ? where ID = ?"
	_, err = db.Exec(sqlStr, attWho, ID)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	ID, err = GetId(db, user)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	sqlStr = "select thumbNum from user_detail where ID = ?"
	err = db.QueryRow(sqlStr, ID).Scan(&num)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	num--
	sqlStr = "update user_detail set thumbNum = ? where ID = ?"
	_, err = db.Exec(sqlStr, num, ID)
	tool.CheckErr(err)
	return err
}
