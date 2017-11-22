package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func InfoHandler(formatter *render.Render) http.HandlerFunc{
	return func (w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			ID      string
			Content string
		}{ID: "0825", Content: "echobee"})
	}
}
