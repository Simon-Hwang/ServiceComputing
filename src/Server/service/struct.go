package service


type User struct{
	Username string `json:"Username"`
	Password string `json:"password"`
}

type Tag struct{
	Tag_content string `json:"Tag_content"`
}

type Date struct{
	Year, Month, Day, Hour, Minute int
}

type Comments struct{
	Comments_content []Comment `json:"Comments_content"`
}

type ArticlesResponse struct {
	Articles []ArticleResponse`json:"articles,omitempty"`
}

