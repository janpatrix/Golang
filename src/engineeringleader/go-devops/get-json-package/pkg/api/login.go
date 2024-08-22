package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LoginRequest struct {
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func DoLoginRequest(client http.Client, requestURL, password string) (string, error) {

	loginRequest := LoginRequest{
		Password: password,
	}

	body, err := json.Marshal(loginRequest)
	if err != nil {
		return "", fmt.Errorf("marshal error: %s", err)
	}

	response, err := client.Post(requestURL, "application/json", bytes.NewBuffer(body))

	if err != nil {
		return "", fmt.Errorf("http post: %s", err)
	}

	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("ReadAll error: %s", err)
	}

	if response.StatusCode != 200 {
		return "", fmt.Errorf("invalid output (HTTP Code %d): %s", response.StatusCode, string(resBody))
	}

	if !json.Valid(resBody) {
		return "", UnmarshalErr("No valid JSON returned", err, response, resBody)
	}

	var loginResponse LoginResponse

	err = json.Unmarshal(resBody, &loginResponse)
	if err != nil {
		return "", UnmarshalErr("LoginResponse unmarshal error:", err, response, resBody)
	}

	if loginResponse.Token == "" {
		return "", UnmarshalErr("Login token is empty", err, response, resBody)
	}

	return loginResponse.Token, nil
}
