package api

import (
	"fmt"
	"net/http"
	"os"
)

type RequestError struct {
	HTTPCode int
	Body     string
	Err      string
}

func (e RequestError) Error() string {
	return e.Err
}

func UnmarshalErr(errString string, err error, response *http.Response, body []byte) RequestError {
	return RequestError{
		HTTPCode: response.StatusCode,
		Body:     string(body),
		Err:      fmt.Sprintf(errString + fmt.Sprintf("%s", err)),
	}
}

func RequestErr(err error) {
	if requestErr, ok := err.(RequestError); ok {
		fmt.Printf("Error: %s HTTP Code: %d, BODY: %s\n", requestErr.Err, requestErr.HTTPCode, requestErr.Body)
		os.Exit(1)
	}
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}
