package main

import (
	"fmt"
	"strings"
)

type Occurrence struct {
	Words map[string]int `json:"words"`
}

func (o Occurrence) getResponse() string {
	out := []string{}
	for word, occurence := range o.Words {
		out = append(out, fmt.Sprintf("%s (%d)", word, occurence))
	}
	return (strings.Join(out, ", \n"))
}
