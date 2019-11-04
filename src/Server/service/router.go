package service

import (
    "net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"strings"
)

type Router struct{
	Router_name string
	Router_method string
	Router_pattern string
	Router_handlerFunc http.HandlerFunc
}

type Routes []Router

var routes = Routes{//deal with different routes
	Router{
		"Index",
		"GET",
		"/v3/",
		Index,
	},

	Router{
		"GetArticleById",
		strings.ToUpper("Get"),
		"/v3/article/{id}",
		GetArticleById,
	},

	Router{
		"GetArticles",
		strings.ToUpper("Get"),
		"/v3/articles",
		GetArticles,
	},

	Router{
		"GetCommentsOfArticle",
		strings.ToUpper("Get"),
		"/v3/article/{id}/comments",
		GetCommentsOfArticle,
	},

	Router{
		"CreateComment",
		strings.ToUpper("Post"),
		"/v3/article/{id}/comment",
		CreateComment,
	},

	Router{
		"SignIn",
		strings.ToUpper("Post"),
		"/v3/auth/signin",
		SignIn,
	},

	Router{
		"SignUp",
		strings.ToUpper("Post"),
		"/v3/auth/signup",
		SignUp,
	},

	Router{
		"OPTIONS",
		strings.ToUpper("options"),
		"/v3/auth/signin",
		Options,
	},

	Router{
		"OPTIONS",
		strings.ToUpper("options"),
		"/v3/article/{id}/comment",
		Options,
	},

	Router{
		"OPTIONS",
		strings.ToUpper("options"),
		"/v3/auth/signup",
		Options,
	},
}

func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes{
		var handler http.Handler
		handler = route.Router_handlerFunc
		handler = Logger(handler, route.Router_name)
		router.Methods(route.Router_method).Path(route.Router_pattern).Name(route.Router_name).Handler(handler)
	}
	return router
}

func Index(w http.ResponseWriter, r *http.Request){
	str, _ := ioutil.ReadFile("./service/api.json")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content_Type", "application/json") // set response type
	w.Write(str)
}

func Options(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,Authorization")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
}
