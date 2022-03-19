package service

import (
	"database/sql"
	"douban/Struct"
	"douban/dao"
	"douban/middleware"
	"douban/tool"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const ClientId = "2b945d7341827be023c4"
const ClientSecret = "e2299bbc41e5bfb23e2fa5c5a13988d05863a4be"

// QueryRowDemoPassword 检验密码和账号是否正确。
func QueryRowDemoPassword(db *sql.DB, username string, pwd string) bool {
	flag := dao.CheckUser(db, username, pwd)
	return flag
}

// RegisterUser 注册服务
func RegisterUser(db *sql.DB, username string, password string, answer string, question string) error {
	err := dao.InsertUser(db, username, password, answer, question)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	id, err := dao.GetId(db, username)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	err = dao.InsertUserDetail(db, id)
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	return nil
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

func AddCap(db *sql.DB, url string, username string) error {
	err := dao.AddCap(db, url, username)
	tool.CheckErr(err)
	return err
}

func GitLogin(clientId string) error {
	str := "https://github.com/login/oauth/authorize?client_id=" + clientId + "?state=" + tool.RandString(8)
	_, err := http.NewRequest(http.MethodGet, str, nil)
	return err
}

func OauthCheck(c *gin.Context, code string) error {

	postValue := url.Values{
		"client_id":     {ClientId},
		"client_secret": {ClientSecret},
		"code":          {code},
	}

	var user Struct.OriginUser
	resp, err := http.PostForm("https://github.com/login/oauth/access_token", postValue)
	tool.CheckErr(err)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		tool.CheckErr(err)
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	tool.CheckErr(err)
	if err != nil {
		return err
	}

	str := string(body)
	var str1, str2 string
	if str[:13] == "access_token=" {
		i := strings.Index(str, "&")
		str1 = str[13:i]
		i = strings.Index(str, "token_type=")
		fmt.Println(str[i+len("token_type="):])
		str2 = str[i+len("token_type="):]
	}
	str = str2 + " " + str1

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", strings.NewReader(""))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", str)
	resp, err = client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_ = json.Unmarshal(body, &user)

	db := tool.GetDb()
	flag := dao.JudgeGitID(db, user.Id)
	if flag {
		middleware.GenerateToken(c, user.Name)
		return nil
	}

	username := user.Name
	password := ClientSecret
	err = dao.InsertUser(db, username, password, "", "")
	tool.CheckErr(err)
	if err != nil {
		return err
	}
	middleware.GenerateToken(c, username)

	return nil
}
