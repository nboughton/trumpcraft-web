package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/nboughton/config/parser"
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

func init() {
	// Read config
	cfgFile := flag.String("c", "config.json", "Path to config file")
	flag.Parse()

	p, err := parser.NewParser(*cfgFile)
	if err != nil {
		log.Fatal(err)
	}
	if err := p.Scan(&cfg); err != nil {
		log.Fatal(err)
	}

	// Set application directory
	cfg.AppDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
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
}

func main() {
	// Configure router
	r := mux.NewRouter()

	r.HandleFunc("/api/{source}/{words}", ReqHandler).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public/")))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r))
}
