package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func randomFloat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random number is %f ", rand.Float64())
}

func randomInt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random number is %d ", rand.Intn(100))
}

func main2() {
	mux := http.NewServeMux()
	mux.HandleFunc("/randomInt", randomInt)
	mux.HandleFunc("/randomFloat", randomFloat)
	http.ListenAndServe(":8000", mux)
}
