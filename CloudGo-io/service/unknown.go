package service

import (
	"net/http"
)

func UnknownHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusNotImplemented, "501 NotImplemented")
	}
}
