package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/drmarduk/mvd2/api/controller"
	"github.com/drmarduk/mvd2/shared"
	"github.com/drmarduk/mvd2/shared/db"
	"github.com/julienschmidt/httprouter"
)

func main() {
	configfile := flag.String("config", "config.json", "the configuration file to use")
	flag.Parse()

	cfg, err := shared.NewConfig(*configfile)
	if err != nil {
		log.Fatalf("config could not be loaded: %sv\n", err)
	}

	ctxDB, err := db.NewDBContext(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBDatabase, cfg.DBDriver)

	// Create our main router
	router := httprouter.New()

	// /index Handler
	index := controller.NewIndexController(ctxDB)
	router.GET("/", index.IndexHandler)

	// /queue handler
	queue := controller.NewQueueController(ctxDB)
	router.GET("/queue", queue.IndexHandler)
	router.POST("/queue/add", queue.AddHandler)

	// /notensatz handler
	notensatz := controller.NewNotenSatzController(ctxDB)
	router.GET("/api/notensatz", notensatz.Get)
	router.GET("/api/notensatz/:id", notensatz.Get)
	router.POST("/api/notensatz", notensatz.Add)
	router.PUT("/api/notensatz/:id", notensatz.Update)
	router.DELETE("/api/notensatz/:id", notensatz.Delete)

	http.ListenAndServe(cfg.HTTPAddress(), router)
}
