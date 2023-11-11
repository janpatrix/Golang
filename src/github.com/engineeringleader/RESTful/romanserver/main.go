package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/engineeringleader/romanNumerals"
)

func main2() {
	http.HandleFunc("/", getRomanNumeral)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getRomanNumeral(w http.ResponseWriter, r *http.Request) {
	urlPathElements := strings.Split(r.URL.Path, "/")
	if urlPathElements[1] == "roman_number" {
		number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
		if number == 0 || number > 10 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not Found"))
		} else {
			fmt.Fprintf(w, "%q", html.EscapeString(romanNumerals.Numerals[number]))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad request"))
	}
}
