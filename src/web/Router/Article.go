package Router


type Article struct{
	Article_id int 
	Article_name string 
	Article_content string 
}

type ArticleResponse struct {
	Id int 
	Name string
}

type Comment struct{
	Comment_content string 
	Comment_publisher string 
	Article_id int 
}

