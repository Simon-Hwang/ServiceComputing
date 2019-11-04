package main

import (
    "log"
    "net/http"
	"web/Helper"
	"web/Router"
	// _ "github.com/go-sql-driver/mysql"
	// "database/sql"
)

//处理登录  
type User struct {
	Username string
	Password string
}

var sessionMgr *Helper.SessionMgr = nil //session管理器


func main() {
	sessionMgr = Helper.NewSessionMgr("TestCookieName", 3600)
    //http.HandleFunc("/", Router.Skip)       //设置访问的路由
	http.HandleFunc("/login", Router.Login)         //设置访问的路由
	http.HandleFunc("/upload", Router.Upload)
	http.HandleFunc("/register", Router.Register)
	http.HandleFunc("/article/3", Router.ArticlesByID)
	http.HandleFunc("/articles", Router.Articles)
    err := http.ListenAndServe(":3333", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

