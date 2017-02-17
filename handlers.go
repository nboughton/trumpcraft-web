package main

import (
	"strconv"

	"github.com/nboughton/misc/markov"
	"github.com/pilu/traffic"
)

// JSONData is exactly what you think it is
type JSONData struct {
	Text string `json:"text"`
}

// RootHandler renders and returns the main page
func RootHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.Render("index", cfg)
}

// ReqHandler returns generated markov strings from the source specified
// in the request URL or picks a source if the one given is invalid
func ReqHandler(w traffic.ResponseWriter, r *traffic.Request) {
	var (
		source   = r.Param("source")
		words, _ = strconv.Atoi(r.Param("words"))
		text     string
	)

	// check that source is valid before checking string length
	_, ok := data[source]
	if !ok {
		// pseudo-randomly pick a source since maps do not return in order
		// of assignment
		for k := range data {
			source = k
			break
		}
	}

	// words should never be higher than 1000 or lower than 10
	if words > 1000 {
		words = 1000
	} else if words < 10 {
		words = 10
	}

	// ensure we have return output
	for len(text) < 10 {
		text = markov.TrimToSentence(data[source].Generate(words))
		words += 5
	}

	w.WriteJSON(&JSONData{Text: text})
}
