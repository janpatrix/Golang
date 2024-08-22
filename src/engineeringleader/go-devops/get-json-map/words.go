package main

import "strings"

type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func (w Words) getResponse() string {
	return strings.Join(w.Words, ", \n")
}
