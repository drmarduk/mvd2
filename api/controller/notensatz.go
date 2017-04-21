package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/drmarduk/mvd2/api/model"
	"github.com/julienschmidt/httprouter"
)

// NotenSatzController handles all requests to /notensatz
type NotenSatzController struct{}

// NewNotenSatzController returns a handler for /notensatz
func NewNotenSatzController() *NotenSatzController {
	return &NotenSatzController{}
}

// Get returns a single NotenSatz instance based on the id
func (n *NotenSatzController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_id := p.ByName("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		log.Printf("error, invalid atoi conversion of %s: %v\n", _id, err)
		return
	}

	ns, err := model.OpenNotenSatz(id)
	if err != nil {
		log.Printf("error while opening notensatz %d: %v\n", id, err)
		return
	}
	JSONRender(ns, w)
	return
}

// Add adds a new NotenSatz instance to the database
func (n *NotenSatzController) Add(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// Update updates an existing NotenSatz
func (n *NotenSatzController) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// Delete deletes a instance from the database
func (n *NotenSatzController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// Trace is a debug function
func (n *NotenSatzController) Trace(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
