package service

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"strconv"
	"io/ioutil"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct{
	Article_id int `json:"Article_id"`
	Article_name string `json:"Article_name"`
	Article_tag []Tag `json:"Article_tag"`
	Article_date string `json:"Article_date"`
	Article_content string `json:"Article_content"` //
}

type ArticleResponse struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Comment struct{
	Comment_date string `json:"Comment_date"`
	Comment_content string `json:"Comment_content"`
	Comment_publisher string `json:"Comment_publisher"`
	Article_id int `json:"Article_id"`
}


func JsonResponse(response interface{}, w http.ResponseWriter, code int) {
    json, err := json.Marshal(response) // turn into json format
    if err != nil {
        log.Fatal(err)
        return
	}
	// set Handler args
    w.Header().Set("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,Authorization")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
    w.Write(json)
}

func getJSON(rows *sql.Rows)([]byte, error){ // turn sql info into json format
	columns, err := rows.Columns() // make sure there is at least one column and return all col's name
	if err != nil{
		return []byte(""), err
	}
	count := len(columns)
	table_data := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next(){  // check each row and write into table_data
		for i := 0; i < count; i++{
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...) // write into values according to the col's name
		entry := make(map[string]interface{}) //deal with the (key,value), while interface make value possible to be any kinds
		for i, col := range columns{
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok{
				v = string(b)
			}else{
				v = val
			}
			entry[col] = v
		}
		table_data = append(table_data, entry)
	}
	json_data, err := json.Marshal(table_data) // turn into json format
	if err != nil {
		return []byte(""), err
	}
	return json_data, nil
}

func GetArticleById(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql", "root:baby942.@tcp(mysql:3306)/?charset=utf8") //open sql with -u testuser -p 123
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	articleId := strings.Split(r.URL.Path, "/")[3] // "/v3/article/{id}/"
	_, err = strconv.Atoi(articleId)
	if err != nil{
		reponse := ErrorResponse{"Wrong ArticleId"}
		JsonResponse(reponse, w, http.StatusBadRequest)
		return
	}
	query, err := db.Query("select * from test.Article where id=" + articleId)//query info from sql
	if err != nil{
		log.Fatal(err)
	}
	defer query.Close()
	v, err := getJSON(query)
	if err != nil{
		log.Fatal(err)
	}
	if string(v) == "[]"{
		response := ErrorResponse{"Ariticle Not Exists"}
		JsonResponse(response,w , http.StatusBadRequest)
		return
	}
	v = v[1:len(v)-1]  // v[0] corresponding to columns name
	str := strings.Replace(string(v), "id\":\"", "id\":", -1)// replace id":" into id": and ","name into ,"name
	str = strings.Replace(str, "\",\"name", ",\"name", -1)// -1 means will replace all match string
	v = []byte(str)
	var article Article
	_ = json.Unmarshal(v, &article)
	file, err := ioutil.ReadFile(article.Article_content)
	if err != nil {
		log.Fatal(err)
	}
	article.Article_content = string(file)
	JsonResponse(article, w, http.StatusOK) // turn into json format and return 
}

func GetArticles(w http.ResponseWriter, r *http.Request){ // get articile catalogue
	db, err := sql.Open("mysql", "root:baby942.@tcp(mysql:3306)/?charset=utf8") //open sql with -u testuser -p 123
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	u, err := url.Parse(r.URL.String()) // parse router info 
	if err != nil{
		log.Fatal(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	page := m["page"][0]
	IdIndex, err := strconv.Atoi(page)
	IdIndex = (IdIndex - 1)* 10
	Id := strconv.Itoa(IdIndex)
	query, err := db.Query("select * from test.Article limit " + Id + ",10")//limit the scope [Id, 10]
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()
	v, err := getJSON(query)
	if err != nil {
		log.Fatal(err)
	}
	if string(v) == "[]" {
		reponse := ErrorResponse{"Page is out of index"}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}
	var article ArticleResponse 
	v = []byte("{\"articles\":"+ string(v) + "}") // { "article" : v }
	str := strings.Replace(string(v), "id\":\"", "id\":", -1)
	str = strings.Replace(str, "\",\"name", ",\"name", -1)
	v = []byte(str)
	_ = json.Unmarshal(v, &article)
	JsonResponse(article, w, http.StatusOK)
}

func GetCommentsOfArticle(w http.ResponseWriter, r *http.Request){
	// "/v3/article/{id}/comments"
	db, err := sql.Open("mysql", "root:baby942.@tcp(mysql:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()
	articleId := strings.Split(r.URL.Path, "/")[3]
	_, err = strconv.Atoi(articleId)
	if err != nil {
		reponse := ErrorResponse{"Wrong ArticleId"}
		JsonResponse(reponse, w, http.StatusBadRequest)
		return
	}
	// check if the article exist before finding the comments
	query, err := db.Query("select * from test.Article where id=" + articleId)
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()
	if !query.Next() {
		response := ErrorResponse{"Article Not Exists"}
		JsonResponse(response, w, http.StatusBadRequest)  
		return
	}
	query, err = db.Query("select * from test.Comment where articleId=" + articleId)
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()
	v, err := getJSON(query)
	if err != nil {
		log.Fatal(err)
	}
	var comments Comments
	v = []byte("{\"content\":" + string(v) + "}")
	str := strings.Replace(string(v), "Id\":\"", "Id\":", -1)
	str = strings.Replace(str, "\",\"author", ",\"author", -1)
	v = []byte(str)
	_ = json.Unmarshal(v, &comments)
	JsonResponse(comments, w, http.StatusOK)
}