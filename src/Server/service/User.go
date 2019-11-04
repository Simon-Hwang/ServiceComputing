package service

import(
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"encoding/binary"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	Key = "256SHA64"
)

func fatal(err error){
	if err != nil {
		log.Fatal(err)
	}
}

type Response struct {
    Data string `json:"data"`
}

type Token struct {
    Token string `json:"token"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}

func itob(v int) []byte { //turn into btye
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(v))
    return b
}

func ByteSliceEqual(a, b []byte) bool { // compare two btye variable 
    if len(a) != len(b) {
        return false
    }
    if (a == nil) != (b == nil) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:baby942.@tcp(mysql:3306)/?charset=utf8")
	if err != nil {
			log.Fatal(err)
	}
	defer db.Close()
	// /v3/article/{id}/comment
	articleId  := strings.Split(r.URL.Path, "/")[3]
	Id, err:= strconv.Atoi(articleId)
	if err != nil {
		response := ErrorResponse{"Wrong ArticleId"}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}
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

	comment := &Comment{
		Comment_date:  time.Now().Format("2000-01-01 00:00:00"),
		Comment_content: "",
		Comment_publisher: "",
		Article_id: Id,
	}
	err = json.NewDecoder(r.Body).Decode(&comment)

	if err != nil  || comment.Comment_content == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err != nil {
			response := ErrorResponse{err.Error()}
			JsonResponse(response, w, http.StatusBadRequest)
		} else {
			response := ErrorResponse{"There is no content in your article"}
			JsonResponse(response, w, http.StatusBadRequest)
		} 
		return
	}
	// verify the info
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
        func(token *jwt.Token) (interface{}, error) {
            return []byte(comment.Comment_publisher), nil
        })

    if err == nil {
        if token.Valid {
			query, err = db.Query("INSERT INTO `test`.`Comment` (`date`, `author`, `articleId`, `content`) VALUES ('" + comment.Comment_date + "', '" + comment.Comment_publisher + "', " + articleId + ", '" + comment.Comment_content + "')")
			if err != nil {
				log.Fatal(err)
			}
			defer query.Close()
			JsonResponse(comment, w, http.StatusOK)
        } else {
			response := ErrorResponse{"Token is not valid"}
			JsonResponse(response, w, http.StatusUnauthorized)
        }
    } else {
		response := ErrorResponse{"Unauthorized access to this resource"}
		JsonResponse(response, w, http.StatusUnauthorized)
    }
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:baby942.@tcp(mysql:3306)/?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response := ErrorResponse{err.Error()}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}
	query, err := db.Query("select * from test.User where username='" + user.Username + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()
	v, err := getJSON(query)
	if err != nil {
		log.Fatal(err)
	}
	if string(v) == "[]" {
		reponse := ErrorResponse{"Wrong Username or Password"}
		JsonResponse(reponse, w, http.StatusNotFound)
		return
	}
	var userQuery User
	v = v[1:len(v)-1]
	_ = json.Unmarshal(v, &userQuery)
	if userQuery.Password != user.Password {
		response := ErrorResponse{"Wrong Username or Password"}
		JsonResponse(response, w, http.StatusNotFound)
		return
	}
	// return cookie to client 
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	if err != nil {
			fatal(err)
	}
	tokenString, err := token.SignedString([]byte(user.Username))
	if err != nil {
			fatal(err)
	}
	response := Token{tokenString}
	JsonResponse(response, w, http.StatusOK)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:baby942.@tcp(mysql:3306)/?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Password == "" || user.Username == "" {
		response := ErrorResponse{"Wrong Username or Password"}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}

	query, err := db.Query("select * from test.User where username='" + user.Username + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()
	// check whether user name duplicate
	if query.Next() { 
		response := ErrorResponse{"User Exists"}
		JsonResponse(response, w, http.StatusBadRequest)  
		return
	}
	query, err = db.Query("INSERT INTO `test`.`User` (`username`, `password`) VALUES ('" + user.Username + "', '" + user.Password + "')")
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}