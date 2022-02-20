package Struct

type MovieLiker struct {
	IMDB       string  `json:"imdb"`
	Score      float32 `json:"score"`
	Num1       int     `json:"num_1"`
	Num2       int     `json:"num_2"`
	Num3       int     `json:"num_3"`
	Num4       int     `json:"num_4"`
	Num5       int     `json:"num_5"`
	NumWant    int     `json:"num_want"`
	NumDone    int     `json:"num_done"`
	CommentNum int     `json:"comment_num"`
	EssayNum   int     `json:"essay_num"`
}

type MovieMember struct {
	IMDB         string `json:"imdb"`
	Director     string `json:"director"`
	Scriptwriter string `json:"scriptwriter"`
	Player       string `json:"player"`
}

type MovieView struct {
	IMDB       string `json:"imdb"`
	Brief      string `json:"brief"`
	PictureNum int    `json:"picture_num"`
	Picture1   string `json:"picture_1"`
	Picture2   string `json:"picture_2"`
	Picture3   string `json:"picture_3"`
	Picture4   string `json:"picture_4"`
	Picture5   string `json:"picture_5"`
	Video      string `json:"video"`
}

type MovieInfo struct {
	IMDB  string `json:"imdb"`
	Name  string `json:"name"`
	Date  string  `json:"date"`
	Long  int    `json:"long"`
	Alias string `json:"alias"`
	Type  string `json:"type"`
}
