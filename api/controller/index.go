package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// IndexController handles all request to /index
type IndexController struct{}

// NewIndexController returns a new controller for /index
func NewIndexController() *IndexController {
	return &IndexController{}
}

// IndexHandler handles /index
func (i *IndexController) IndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("hello world"))
}
