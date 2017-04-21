package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	configfile := flag.String("config", "config.json", "the configuration file to use")
	flag.Parse()

	cfg, err := NewConfig(*configfile)
	if err != nil {
		log.Fatalf("config could not be loaded: %sv\n", err)
	}

	router := httprouter.New()

	http.ListenAndServe(cfg.HTTPAddress(), router)
}
