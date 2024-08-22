package api

import "net/http"

type MyJWTTransport struct {
	Transport http.RoundTripper
	Token     string
}

func (m MyJWTTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if m.Token != "" {
		r.Header.Add("Authorization", "Bearer "+m.Token)
	}
	return (m.Transport.RoundTrip(r))
}
