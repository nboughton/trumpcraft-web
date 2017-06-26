package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nboughton/misc/markov"
)

// JSONData is exactly what you think it is
type JSONData struct {
	Text string `json:"text"`
}

// ReqHandler returns generated markov strings from the source specified
// in the request URL or picks a source if the one given is invalid
func ReqHandler(w http.ResponseWriter, r *http.Request) {
	var (
		p        = mux.Vars(r)
		source   = p["source"]
		words, _ = strconv.Atoi(p["words"])
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

	JSON{Status: 200, Data: text}.write(w)
}
