package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func apihandle(w http.ResponseWriter, req *http.Request) {
	rdr := render.New()
	rdr.JSON(w, http.StatusOK, struct {
		ID      string
		Content string
	}{ID: "0825", Content: "echobee"})
}
