package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Page struct {
	Name string `json:"page"`
}

type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

type Occurrence struct {
	Words map[string]int `json:"words"`
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: ./hello-get <url>")
		os.Exit(1)
	}

	myUrl, err := url.ParseRequestURI(args[1])
	if err != nil {
		fmt.Println("URL is not a valid url")
		os.Exit(1)
	}

	fmt.Printf("MyURL: %v\n", myUrl)

	response, err := http.Get(myUrl.String())
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Code %d): %s\n", response.StatusCode, string(body))
		os.Exit(1)
	}

	var page Page

	err = json.Unmarshal(body, &page)
	if err != nil {
		log.Fatal(err)
	}

	switch page.Name {
	case "words":
		var words Words

		err = json.Unmarshal(body, &words)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("JSON parsed:\nPage: %s\nWords: %s\n", page.Name, strings.Join(words.Words, ", "))

	case "occurrence":
		var occurrence Occurrence

		err = json.Unmarshal(body, &occurrence)
		if err != nil {
			log.Fatal(err)
		}
		for word, occurrence := range occurrence.Words {
			fmt.Printf("%s: %d\n", word, occurrence)
		}
	default:
		fmt.Printf("Page not found")
	}

}
