package service

import (
    "net/http"
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
    // New constructs a new Render instance with the supplied options.
    formatter := render.New(render.Options{ 
        IndentJSON: true,
    }) // rendering a web format

    n := negroni.Classic()
    mx := mux.NewRouter()

    initRoutes(mx, formatter)
    // set the router 
    n.UseHandler(mx)
    return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
    // Methods adds a matcher for HTTP methods.
    // It accepts a sequence of one or more methods to be matched, e.g.:
    // "GET", "POST", "PUT".
    mx.HandleFunc("/{status}", testHandler(formatter)).Methods("GET")
    //provide a hanlder func to deal with specific router
    //corresponding to func(w http.ResponseWriter, req *http.Request)
}

func testHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        vars := mux.Vars(req)
        id := vars["id"]
        status := vars["status"]
        formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + status + id})
    }
}