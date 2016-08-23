package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/nboughton/config/parser"
	"github.com/nboughton/misc/markov"
	"github.com/pilu/traffic"
)

// Config struct for configuration
type Config struct {
	Files  map[string]string `json:"files"`
	Port   int               `json:"port"`
	AppDir string
}

var (
	router *traffic.Router
	cfg    Config
	data   map[string]*markov.Chain
)

func init() {
	// Read config
	cfgFile := flag.String("c", "config.json", "Path to config file")
	flag.Parse()

	p := parser.NewParser(*cfgFile)
	p.Scan(&cfg)

	// Set application directory
	var err error
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

		data[k] = markov.NewChain(2)
		data[k].Build(f)
	}

	// Configure router
	router = traffic.New()
	traffic.SetPort(cfg.Port)
	router.Get("/", RootHandler)
	router.Get("/api/:source/:words", ReqHandler)
}

func main() {
	router.Run()
}
