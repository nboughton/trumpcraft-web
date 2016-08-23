package main

import (
	"strconv"

	"github.com/nboughton/misc/markov"
	"github.com/pilu/traffic"
)

// ResponseData is exactly what you think it is
type ResponseData struct {
	Text string
}

// RootHandler renders and returns the main page
func RootHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.Render("index")
}

// ReqHandler returns generated markov strings from the source specified
// in the request URL or picks a source if the one given is invalid
func ReqHandler(w traffic.ResponseWriter, r *traffic.Request) {
	var (
		file     = r.Param("source")
		words, _ = strconv.Atoi(r.Param("words"))
		text     = markov.TrimToSentence(data[file].Generate(words))
	)

	// check that source is valid before checking string length
	_, ok := data[file]
	if !ok {
		// pseudo-randomly pick a source since maps do not return in order
		// of assignment
		for k := range data {
			file = k
			break
		}
	}

	// ensure we have return output
	for len(text) < 10 {
		words += 5
		text = markov.TrimToSentence(data[file].Generate(words))
	}

	w.WriteJSON(&ResponseData{Text: text})
}
