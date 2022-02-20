package tool

import "fmt"

// CheckErr 检查错误
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

// TurnErr 检查nil并修改为string
func TurnErr(str interface{}) string {
	if str == nil {
		return ""
	}
	r := string(str.([]uint8))
	return r
}
