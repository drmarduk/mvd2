package controller

import (
	"log"
	"net/http"

	"github.com/drmarduk/mvd2/api/model"
	"github.com/drmarduk/mvd2/shared/db"
	"github.com/julienschmidt/httprouter"
)

// QueueController is responsible for incoming pdf files
type QueueController struct {
	ctx *db.DBContext
}

// NewQueueController gets a db context and handles the handlers
func NewQueueController(ctx *db.DBContext) *QueueController {
	return &QueueController{ctx}
}

// IndexHandler returns a list with all files in the queeu
func (q *QueueController) IndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
		<form method="post" action="/queue/add">
			<input type="text" name="filename" />

			<input type="submit" value="adden" />
		</form>
	
	`))
	qs, err := model.OpenQueue(q.ctx)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	for _, e := range qs {
		w.Write([]byte(e.Filename))
	}
}

// AddHandler lets you insert a new pdf file by hand
func (q *QueueController) AddHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// process only submissions, otherwise or at success -> /queue

	err := model.AddQueue(q.ctx, "filename", "hash", 1234)
	if err != nil {
		log.Printf("error %v\n", err)
	}
	http.Redirect(w, r, "/queue", 201)
}
