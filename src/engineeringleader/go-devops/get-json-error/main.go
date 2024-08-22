package main

import (
	"encoding/json"
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
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: ./hello-get <url>")
		os.Exit(1)
	}

	res, err := doRequest(args[1])
	if err != nil {
		if requestErr, ok := err.(RequestError); ok {
			fmt.Printf("Error: %s HTTP Code: %d, BODY: %s\n", requestErr.Err, requestErr.HTTPCode, requestErr.Body)
			os.Exit(1)
		}
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if res == nil {
		fmt.Printf("No response\n")
		os.Exit(1)
	}

	fmt.Printf("Response %s", res.getResponse())
}

func doRequest(requestURL string) (Response, error) {

	var page Page

	myUrl, err := url.ParseRequestURI(requestURL)
	if err != nil {
		return nil, fmt.Errorf("URL Validation error: URL is not valid %s", err)
	}

	fmt.Printf("MyURL: %v\n", myUrl)

	response, err := http.Get(myUrl.String())
	if err != nil {
		return nil, fmt.Errorf("HTTPGet error: %s", err)
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
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      "No valid JSON returned.",
		}
	}

	err = json.Unmarshal(body, &page)
	if err != nil {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Sprintf("Page unmarshal error: %s", err),
		}
	}

	switch page.Name {
	case "words":
		var words Words

		err = json.Unmarshal(body, &words)
		if err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Unmarshal error: %s", err),
			}
		}

		return words, nil
	case "occurrence":
		var occurrence Occurrence

		err = json.Unmarshal(body, &occurrence)
		if err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Occurrence error: %s", err),
			}
		}
		return occurrence, nil
	}
	return nil, nil
}
