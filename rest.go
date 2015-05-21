package main

import (
	"encoding/json"
	"fmt"
	"github.com/sthorne/go-hunspell"
	"net/http"
	"regexp"
)

var h = hunspell.Hunspell("/usr/share/hunspell/en_GB.aff", "/usr/share/hunspell/en_GB.dic")

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/webservices/php/SpellChecker.php", legacyHandler)
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/suggest", suggestHandler)
	http.Handle("/", fs)

	http.ListenAndServe(":80", nil)
}

var wordRegexp = regexp.MustCompile(`\b[\w']+\b`)

func checkHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")

	typos := []string{}
	for _, word := range wordRegexp.FindAllString(text, -1) {
		if !h.Spell(word) {
			typos = append(typos, word)
		}
	}

	result := map[string]interface{}{
		"typos": typos,
	}

	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json error:", err)
		return
	}

	w.Write(b)
}

func suggestHandler(w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("word")

	result := h.Suggest(word)

	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json error:", err)
		return
	}

	w.Write(b)
}

func legacyHandler(w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("action")
	var result interface{}

	switch (action) {
	case "get_incorrect_words":
		text := r.FormValue("text[]")
		typos := []string{}
		for _, word := range wordRegexp.FindAllString(text, -1) {
			if !h.Spell(word) {
				typos = append(typos, word)
			}
		}

		result = map[string]interface{}{
			"outcome": "success",
			"data": [][]string{typos},
		}

	case "get_suggestions":
		word := r.FormValue("word")
		result = h.Suggest(word)

	default:
		result = map[string]interface{}{
			"outcome": "failure",
		}
	}

	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json error:", err)
		return
	}

	w.Write(b)
}
