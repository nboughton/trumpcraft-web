package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nboughton/go-utils/fs"
	jfile "github.com/nboughton/go-utils/json/file"
	"github.com/nboughton/misc/markov"
)

// Config struct for configuration
type Config struct {
	Files  map[string]string `json:"files"`
	Port   int               `json:"port"`
	Crazy  int               `json:"crazy"` // 0 is craziest, sanity increases with value
	AppDir string
}

var (
	cfg  Config
	data map[string]*markov.Chain
)

func main() {
	// Read config
	c := flag.String("c", "config.json", "Path to config file")
	flag.Parse()

	if err := jfile.Scan(*c, &cfg); err != nil {
		log.Fatal(err)
	}

	// Set application directory
	var err error
	if cfg.AppDir, err = fs.AbsPath(); err != nil {
		log.Fatal(err)
	}

	// Load data into memory
	data = make(map[string]*markov.Chain)
	for k, v := range cfg.Files {
		f, err := os.Open(fmt.Sprintf("%v/sources/%v", cfg.AppDir, v))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		data[k] = markov.NewChain(cfg.Crazy)
		data[k].Build(f)
	}

	// Configure router
	r := mux.NewRouter()

	r.HandleFunc("/{source}/{words}", ReqHandler).Methods("GET")
	r.HandleFunc("/api/trumpcraft/{source}/{words}", ReqHandler).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public/")))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r))
}
