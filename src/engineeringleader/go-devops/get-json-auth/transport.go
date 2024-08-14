package main

import "net/http"

type MyJWTTransport struct {
	transport http.RoundTripper
	token     string
}

func (m MyJWTTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if m.token != "" {
		r.Header.Add("Authorization", "Bearer "+m.token)
	}
	return (m.transport.RoundTrip(r))
}
