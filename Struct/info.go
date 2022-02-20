package Struct

type Info struct {
	Username   string `json:"username"`
	Id         int    `json:"id"`
	Pid        int    `json:"pid"`
	Msg        string `json:"msg"`
	Time       int64  `json:"time"`
	CommentNum int    `json:"comment_num"`
	Thumb      int    `json:"thumb"`
	Liker      string `json:"liker"`
	Movie      string `json:"movie"`
	Point      int    `json:"point"`
	Essay      string `json:"essay"`
	Type       string `json:"type"`
}
