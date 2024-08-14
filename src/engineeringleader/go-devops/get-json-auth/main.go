package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Page struct {
	Name string `json:"page"`
}

type Response interface {
	getResponse() string
}

func main() {
	var (
		requestURL string
		password   string
		parsedURL  *url.URL
		err        error
	)

	flag.StringVar(&requestURL, "url", "", "url to access")
	flag.StringVar(&password, "password", "", "password to access the api")

	flag.Parse()

	if parsedURL, err = url.ParseRequestURI(requestURL); err != nil {
		fmt.Printf("URL Validation error: URL is not valid %s\n", err)
		flag.Usage()
		os.Exit(1)
	}

	if password != "" {
		token, err := doLoginRequest(parsedURL.Scheme+"://"+parsedURL.Host+"/login", password)
		if err != nil {
			requestErr(err)
		}
		fmt.Printf("token: %s\n", token)
		os.Exit(1)
	}

	res, err := doRequest(parsedURL.String())
	if err != nil {
		requestErr(err)
	}

	if res == nil {
		fmt.Printf("No response\n")
		os.Exit(1)
	}

	fmt.Printf("Response %s", res.getResponse())
}

func doRequest(requestURL string) (Response, error) {

	response, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("http get: %s", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll error: %s", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("invalid output (http Code %d): %s", response.StatusCode, string(body))
	}

	if !json.Valid(body) {
		return nil, unmarshalErr("No valid JSON returned", err, response, body)
	}

	var page Page

	err = json.Unmarshal(body, &page)
	if err != nil {
		return nil, unmarshalErr("Page unmarshal error:", err, response, body)
	}

	switch page.Name {
	case "words":
		var words Words

		err = json.Unmarshal(body, &words)
		if err != nil {
			return nil, unmarshalErr("Unmarshal error:", err, response, body)
		}

		return words, nil
	case "occurrence":
		var occurrence Occurrence

		err = json.Unmarshal(body, &occurrence)
		if err != nil {
			return nil, unmarshalErr("Occurrence error:", err, response, body)
		}
		return occurrence, nil
	}
	return nil, nil
}
