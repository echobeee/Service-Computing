package service

import (
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	server := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	server.UseHandler(mx)
	return server
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/hello/{name}", testHandler(formatter)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["name"]
		formatter.JSON(w, http.StatusOK, map[string]string{"name": name})
	}
}
