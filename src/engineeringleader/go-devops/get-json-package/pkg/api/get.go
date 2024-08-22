package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Page struct {
	Name string `json:"page"`
}

type Response interface {
	GetResponse() string
}

func DoRequest(client http.Client, requestURL string) (Response, error) {

	response, err := client.Get(requestURL)
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
		return nil, UnmarshalErr("No valid JSON returned", err, response, body)
	}

	var page Page

	err = json.Unmarshal(body, &page)
	if err != nil {
		return nil, UnmarshalErr("Page unmarshal error:", err, response, body)
	}

	switch page.Name {
	case "words":
		var words Words

		err = json.Unmarshal(body, &words)
		if err != nil {
			return nil, UnmarshalErr("Unmarshal error:", err, response, body)
		}

		return words, nil
	case "occurrence":
		var occurrence Occurrence

		err = json.Unmarshal(body, &occurrence)
		if err != nil {
			return nil, UnmarshalErr("Occurrence error:", err, response, body)
		}
		return occurrence, nil
	}
	return nil, nil
}
