package Router

import(
	"fmt"
    "html/template"
	"net/http"
	"strings"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"os"
	"io"
	"io/ioutil"
)

var (
	cur_username string
	cur_password string
)

var html = `
<!DOCTYPE html>
 
<html lang="zh-ch">
	<head>
		<meta charset="utf-8">
		<title>主页</title>
	</head>
	<body>
		<embed src={{.Addr}} width="800" height="600" ></embed>
	</body>
</html>
`
type Data struct{
	Addr string
}

func Login(w http.ResponseWriter, r *http.Request) {  
    if r.Method == "GET" {  
        t, _ := template.ParseFiles("Resource/login.gtpl")  
        t.Execute(w, nil)  
    } else if r.Method == "POST" {  
		r.ParseForm()
		in_username := r.Form.Get("username")
		in_password := r.Form.Get("password")
		if check_user(in_username, in_password){
			cur_username = in_username
			cur_password = in_password
			http.Redirect(w, r, "/upload", http.StatusFound)
		}else{
			http.Redirect(w, r, "/register", http.StatusFound)
		}
    }  
}

func Upload(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
		if !check_status() {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
		t, _ := template.ParseFiles("Resource/upload.gtpl")
		t.Execute(w, nil)
    } else {
        r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("Resource/uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        fmt.Fprintf(w, "%v", handler.Header)
        f, err := os.OpenFile("./test/" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)
    }
}

func Register(w http.ResponseWriter, r *http.Request){
    if r.Method == "GET" {
		t, _ := template.ParseFiles("Resource/register.gtpl")
		t.Execute(w, nil)
    }else{
		r.ParseForm()
		in_username := r.Form.Get("username")
		in_password := r.Form.Get("password")
		fmt.Println("Insert: ", in_username, in_password)
		if check_user(in_username, in_password) {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		db, err := sql.Open("mysql", "root:baby942.@tcp(localhost)/test?charset=utf8")
		stmt, err := db.Prepare("INSERT user SET username=?,password=?")
		checkErr(err)
		res, err := stmt.Exec(in_username, in_password)
		checkErr(err)
		id, err := res.LastInsertId()
		checkErr(err)
		fmt.Println("Sign in ",id)
		db.Close()
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func ArticlesByID(w http.ResponseWriter, r *http.Request){ // SOme problem
	// if r.Method == "GET" {
	// 	if !check_status() {
	// 		http.Redirect(w, r, "/login", http.StatusFound)
	// 		return
	// 	}
	// }
	articleId := strings.Split(r.URL.Path, "/")[2] // "/article/{id}/"
	id, err := strconv.Atoi(articleId)
	checkErr(err)
	db, err := sql.Open("mysql", "root:baby942.@tcp(localhost)/test?charset=utf8")
	checkErr(err)
	rows, err := db.Query("SELECT * FROM article")
	checkErr(err)
	for rows.Next() {
		var article_id int
		var article_name, article_content string
		err = rows.Scan(&article_id, &article_name, &article_content)
		checkErr(err)
		if id == article_id {
			db.Close()
			data := Data{Addr: "./" + article_content}
			fmt.Println(data.Addr)
			var err error
   			var t *template.Template
			t = template.New("Products")
			t, err = t.Parse(html) 
			checkErr(err)
			f, err := os.OpenFile("test_2.gtpl",os.O_WRONLY|os.O_CREATE,0666)
			defer f.Close()
			err = t.Execute(f, data)
			checkErr(err)
			t, _ = template.ParseFiles("test_2.gtpl")  
        	t.Execute(w, nil)  
		}
	}
}

func Articles(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		if !check_status() {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		t, _ := template.ParseFiles("Resource/articles.gtpl")
		t.Execute(w, nil)
    }else{
		r.ParseForm()
		articles_id := r.Form.Get("article_id")
		id, err := strconv.Atoi(articles_id)
		checkErr(err)
		db, err := sql.Open("mysql", "root:baby942.@tcp(localhost)/test?charset=utf8")
		checkErr(err)
		rows, err := db.Query("SELECT * FROM article")
		checkErr(err)
		for rows.Next() {
			var article_id int
			var article_name, article_content string
			err = rows.Scan(&article_id, &article_name, &article_content)
			checkErr(err)
			if id == article_id {
				db.Close()
				http.Redirect(w,r, "/article/" + articles_id, http.StatusFound)
				return
			}
		}
	}
}

func Skip(w http.ResponseWriter, r *http.Request){
	str, _ := ioutil.ReadFile("resource/api.json")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content_Type", "application/json") // set response type
	w.Write(str)
}

func check_user(in_username, in_password string) bool{
	db, err := sql.Open("mysql", "root:baby942.@tcp(localhost)/test?charset=utf8")
	checkErr(err)
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)
    for rows.Next() {
		var username string
		var password string
        err = rows.Scan(&username, &password)
		checkErr(err)
		if len(username) != 0 && len(password) != 0 && in_username == username && in_password == password {
			db.Close()
			return true
		}
	}
	db.Close()
	return false
}

func check_status() bool {
	return check_user(cur_username, cur_password)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}