package service

import (
	"database/sql"
	"douban/Struct"
	"douban/dao"
	"douban/tool"
)

// QueryRowDemoPassword 检验密码和账号是否正确。
func QueryRowDemoPassword(db *sql.DB, username string, pwd string) bool {
	flag := dao.CheckUser(db, username, pwd)
	return flag
}

// RegisterUser 注册服务
func RegisterUser(db *sql.DB, username string, password string, answer string, question string) bool {
	err := dao.InsertUser(db, username, password, answer, question)
	tool.CheckErr(err)
	if err != nil {
		return false
	}
	id, err := dao.GetId(db, username)
	tool.CheckErr(err)
	if err != nil {
		return false
	}
	err = dao.InsertUserDetail(db, id)
	tool.CheckErr(err)
	if err != nil {
		return false
	}
	return true
}

// QueryRowDemo 查找用户
func QueryRowDemo(db *sql.DB, username string) (Struct.User, bool) {
	u, flag := dao.FindUser(db, username)
	return u, flag
}

// AddSecret 添加密保
func AddSecret(db *sql.DB, question string, answer string, username string) error {
	err := dao.AddSecret(db, question, answer, username)
	return err
}

// ChangePassword 改密码
func ChangePassword(db *sql.DB, pwd string, username string) error {
	err := dao.UpdateRowDemo(db, pwd, username)
	return err
}

func UpdateHP(db *sql.DB, username string, message string) error {
	ID, err := dao.GetId(db, username)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	err = dao.UpdateHP(db, ID, message)
	return err
}

func GetHP(db *sql.DB, ID int) (string, error) {
	message, err := dao.GetHP(db, ID)
	return message, err
}

func GetDT(db *sql.DB, id int) (Struct.UserDetail, error) {
	u, err := dao.GetDT(db, id)
	return u, err
}
