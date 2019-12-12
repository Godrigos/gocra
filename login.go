package gocra

import (
	"io"
	"net/http"
)

/*
Authenticate prepares a http.Request that will be used by
by other function to login on CIPRES
*/
func authenticate(method string, url string, body io.Reader) *http.Request {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(auth.User, auth.Passwd)
	req.Header.Set("cipres-appkey", auth.AppID)

	return req
}

/*
Login and access server data based on the URL provided,
can get jobs list, status and results.
*/
func login(method string, url string, body io.Reader) *http.Response {

	req := authenticate(method, url, body)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return resp
}
