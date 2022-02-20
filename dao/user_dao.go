package dao

import (
	"database/sql"
	"douban/Struct"
	"douban/tool"
	"errors"
	"fmt"
)

// InsertUser 插入一条信息
func InsertUser(db *sql.DB, username string, password string, answer string, question string) error {
	sqlStr := "insert into user(username, password, secret, answer) values (?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, username, password, question, answer)
	if err != nil {
		fmt.Println("insert failed, err:", err)
		return err
	}
	return nil
}

// UpdateRowDemo 修改密码
func UpdateRowDemo(db *sql.DB, pwd string, username string) error {
	sqlStr := "update user set password=? where username = ?"
	_, err := db.Exec(sqlStr, pwd, username)
	if err != nil {
		fmt.Println("update failed, err:", err)
		return err
	}
	return nil
}

// FindUser 查找用户
func FindUser(db *sql.DB, username string) (Struct.User, bool) {
	sqlStr := "select username, password, secret, answer from user where username = ?"
	var (
		u        Struct.User
		question interface{}
		answer   interface{}
	)
	err := db.QueryRow(sqlStr, username).Scan(&u.Username, &u.Pwd, &question, &answer)
	u.Question = tool.TurnErr(question)
	u.Answer = tool.TurnErr(answer)
	if err != nil {
		fmt.Println("scan failed err:", err)
		return u, false
	}
	return u, true
}

// AddSecret 添加密保
func AddSecret(db *sql.DB, question string, answer string, username string) error {
	sqlStr := "update user set secret = ?, answer = ? where username = ?"
	_, err := db.Exec(sqlStr, question, answer, username)
	tool.CheckErr(err)
	return err
}

// CheckUser 检查账密是否正确
func CheckUser(db *sql.DB, username string, pwd string) bool {
	if username == "" || pwd == "" {
		return false
	}
	sqlStr := "select username, password from user where username = ?"
	row := db.QueryRow(sqlStr, username)
	err := row.Err()
	tool.CheckErr(err)
	if err != nil {
		return false
	}
	var usr, password string
	err = row.Scan(&usr, &password)
	tool.CheckErr(err)
	if usr == username && pwd == password {
		return true
	} else {
		return false
	}

}

// GetId 获取用户的ID
func GetId(db *sql.DB, username string) (int, error) {
	var id int
	sqlStr := "select ID from user where username = ?"
	err := db.QueryRow(sqlStr, username).Scan(&id)
	tool.CheckErr(err)
	return id, err
}

// InsertUserDetail 添加用户细节
func InsertUserDetail(db *sql.DB, id int) error {
	sqlStr := "insert user_detail(id) values (?)"
	_, err := db.Exec(sqlStr, id)
	return err
}

func UpdateHP(db *sql.DB, ID int, message string) error {
	sqlStr := "update user_detail set introduce = ? where ID = ?"
	_, err := db.Exec(sqlStr, message, ID)
	tool.CheckErr(err)
	return err
}

func GetHP(db *sql.DB, ID int) (string, error) {
	sqlStr := "select introduce from user_detail where ID = ?"
	var message string
	err := db.QueryRow(sqlStr, ID).Scan(&message)
	if err == sql.ErrNoRows {
		return "", errors.New("NoUser")
	}
	tool.CheckErr(err)
	if err != nil {
		return "", err
	}
	return message, nil
}

func GetDT(db *sql.DB, id int) (Struct.UserDetail, error) {
	sqlStr := "select mnum, unum, attwho, thumbnum, id, moviedone, moviewant, introduce, commentnum from user_detail where ID = ?"
	var u Struct.UserDetail
	err := db.QueryRow(sqlStr, id).Scan(&u.MNum, &u.UNum, &u.AttWho, &u.ThumbNum, &u.ID, &u.DMovie, &u.WMovie, &u.Introduce, &u.CommentNum)
	if err != nil {
		return u,err
	}
	return u,nil
}
