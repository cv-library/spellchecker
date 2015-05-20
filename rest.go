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

	suggestions := h.Suggest(word)

	result := map[string]interface{}{
		"suggestions": suggestions,
	}

	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json error:", err)
		return
	}

	w.Write(b)
}
