package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"engineeringleader.com/http-login-package/pkg/api/pkg/api"
)

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

	client := http.Client{}

	if password != "" {
		token, err := api.DoLoginRequest(client, parsedURL.Scheme+"://"+parsedURL.Host+"/login", password)
		if err != nil {
			api.RequestErr(err)
		}
		client.Transport = api.MyJWTTransport{
			Transport: http.DefaultTransport,
			Token:     token,
		}
	}

	res, err := api.DoRequest(client, parsedURL.String())
	if err != nil {
		api.RequestErr(err)
	}

	if res == nil {
		fmt.Printf("No response\n")
		os.Exit(1)
	}

	fmt.Printf("Response: %s", res.GetResponse())
}
