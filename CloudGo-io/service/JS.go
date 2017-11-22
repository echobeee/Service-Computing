package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func JsHandler(formatter *render.Render) http.HandlerFunc {
    return func(rw http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        formatter.HTML(rw, http.StatusOK, "table", struct {
          Un string
          Ag string
        }{Un: r.Form["username"][0], Ag: r.Form["age"][0]})
    }


}
