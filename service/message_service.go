package service

import (
	"database/sql"
	"douban/Struct"
	"douban/dao"
	"douban/tool"
	"errors"
	"strconv"
)

// FindAllPMsg 查找所有留言或评论
func FindAllPMsg(db *sql.DB, pid int, movie string) ([]Struct.Info, error) {
	u, err := dao.FindAll(db, pid, movie)
	return u, err
}

// SendMsg 添加留言
func SendMsg(db *sql.DB, username string, msg string, pid int, _type int, movie string, point int) (int, error) {
	if _type == 1 {
		if movie == "comment" {
			u, _ := dao.FindTheMsg(db, strconv.Itoa(pid))
			movieR := u.Movie
			id, err := dao.InsertMsg(db, username, msg, pid, movieR)
			return id, err
		}
		id, err := dao.InsertMsg(db, username, msg, pid, movie)
		return id, err
	} else {
		id, err := dao.InsertEssay(db, username, msg, movie, point)
		return id, err
	}
}

// FindAllComments 展示评论
func FindAllComments(db *sql.DB, pid int) ([]Struct.Info, error) {
	u0 := make([]Struct.Info, 0)
	u, err := dao.ListComments(db, pid, u0)
	return u, err
}

// GetOneMsg 查找留言或评论
func GetOneMsg(db *sql.DB, id string) (Struct.Info, error) {
	u, err := dao.FindTheMsg(db, id)
	return u, err
}

// DeleteMsg 删除留言
func DeleteMsg(db *sql.DB, id string) error {
	err := dao.DeleteMsg(db, id)
	return err
}

// ThumbAdd 点赞
func ThumbAdd(db *sql.DB, id string, username string) error {
	err := dao.ThumbAdd(db, id, username)
	return err
}

// UserPoint 电影想看/看过登记
func UserPoint(db *sql.DB, username string, _type string, movie string) error {
	flag := dao.JudgeWOD(db, username, _type, movie)
	if !flag {
		return errors.New("done")
	}
	err := dao.MoviePointChange(db, _type, movie)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	err = dao.UserPointChange(db, username, _type, movie)
	tool.CheckErr(err)
	return err
}

// DeleteThumb 删除评论点赞
func DeleteThumb(db *sql.DB, username string, ID int) error {
	err := dao.DeleteThumb(db, username, ID)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	return nil
}

func DeletePoint(db *sql.DB, username string, movie string, _type string) error {
	err := dao.DeletePoint(db, username, movie, _type)
	tool.CheckErr(err)
	return err
}

func ThumbUser(db *sql.DB, user string, username string) error {
	err := dao.ThumbUser(db, user, username)
	return err
}

func DeleteUT(db *sql.DB, user string, username string) error {
	err := dao.DeleteUT(db, user, username)
	return err
}
