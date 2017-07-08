package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	importDir := flag.String("import", "", "the directory to watch and import from")
	singleRun := flag.Bool("singlerun", false, "true to run only once, default is false")
	flag.Parse()

	config := &Config{
		ImportDir: *importDir,
	}

	if !check(config.ImportDir) {
		log.Fatal("no valid directory found")
	}

	for {
		// get new files to process
		files, err := ioutil.ReadDir(config.ImportDir)
		if err != nil {
			log.Printf("error while scanning %s: %v\n", config.ImportDir, err)
			continue
		}

		// send them to the api
		go processFiles(files)

		if *singleRun {
			os.Exit(0)
		} else {
			time.Sleep(time.Second * 4)
		}
	}
}

func check(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func processFiles(files []os.FileInfo) {

}
