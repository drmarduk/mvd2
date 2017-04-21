package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/drmarduk/mvd2/api/controller"
	"github.com/julienschmidt/httprouter"
)

func main() {
	configfile := flag.String("config", "config.json", "the configuration file to use")
	flag.Parse()

	cfg, err := NewConfig(*configfile)
	if err != nil {
		log.Fatalf("config could not be loaded: %sv\n", err)
	}

	// Create our main router
	router := httprouter.New()

	// /index Handler
	index := controller.NewIndexController()
	router.GET("/", index.IndexHandler)

	// /notensatz handler
	notensatz := controller.NewNotenSatzController()
	router.GET("/notensatz", notensatz.Get)
	router.GET("/notensatz/:id", notensatz.Get)
	router.POST("/notensatz", notensatz.Add)
	router.PUT("/notensatz/:id", notensatz.Update)
	// router.PATCH("/notensatz/:id", notensatz.Update) // only partial update
	router.DELETE("/notensatz/:id", notensatz.Delete)
	router.Handle("TRACE", "/notensatz", notensatz.Trace)

	http.ListenAndServe(cfg.HTTPAddress(), router)
}
