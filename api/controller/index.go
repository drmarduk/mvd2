package controller

import (
	"net/http"

	"github.com/drmarduk/mvd2/shared/db"
	"github.com/julienschmidt/httprouter"
)

// IndexController handles all request to /index
type IndexController struct {
	ctx *db.DBContext
}

// NewIndexController returns a new controller for /index
func NewIndexController(ctx *db.DBContext) *IndexController {
	return &IndexController{ctx}
}

// IndexHandler handles /index
func (i *IndexController) IndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("hello world"))
}
