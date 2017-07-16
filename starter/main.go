package main

import (
	"flag"
	"log"

	"github.com/drmarduk/mvd2/shared"
	"github.com/drmarduk/mvd2/shared/db"
	"github.com/drmarduk/mvd2/starter/tables"
)

func main() {
	configfile := flag.String("config", "config.json", "the configuration file to use")
	flag.Parse()

	cfg, err := shared.NewConfig(*configfile)
	if err != nil {
		log.Fatalf("config could not be loaded: %sv\n", err)
	}

	ctxDB, err := db.NewDBContext(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBDatabase, cfg.DBDriver)
	if err != nil {
		log.Fatalf("error while creating db context: %v\n", err)
	}

	tables := []tables.Table{
		new(tables.Queue{}),
	}
}
