package service

import (
	"database/sql"
	"douban/dao"
)

func FindMemberIMDB(db *sql.DB, name string) ([]string, error) {
	IMDB, err := dao.FindMemberIMDB(db, name)
	return IMDB, err
}

func ShowMember(db *sql.DB, IMDB string, heading string) (interface{}, error) {
	u, err := dao.ShowMember(db, IMDB, heading)
	return u, err
}
