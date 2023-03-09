package http

import (
	"net/http"
	"time"
)

//var client *http.Client

type MyRoundTripper struct {
	r   http.RoundTripper
	JWT string
}

func (mrt MyRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Authorization", mrt.JWT)
	return mrt.r.RoundTrip(r)
}

func NewCustomHttpClient() *http.Client {
	return &http.Client{
		Transport: MyRoundTripper{r: http.DefaultTransport, JWT: "231231"},
		Timeout:   time.Second * 20,
	}
}
