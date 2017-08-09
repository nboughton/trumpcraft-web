package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	jweb "github.com/nboughton/go-utils/json/web"
	"github.com/nboughton/misc/markov"
)

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
	if _, ok := data[source]; !ok {
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

	jweb.New(http.StatusOK, text).Write(w)
}
