package main

import (
	"flag"
	"log"
)

func main() {
	importDir := flag.String("import", "", "the directory to watch and import from")

	if *importDir == "" {
		log.Fatal("no import directory found")
	}

	config := &Config{
		ImportDir: *importDir,
	}

}
