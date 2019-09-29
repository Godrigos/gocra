package gocra

import (
	"log"
	"net/http"
)

/* Authenticate prepares a http.Request that will be used by
   by other function to login in on CIPRES */
func authenticate(method string, url string) *http.Request {

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(auth.User, auth.Passwd)
	req.Header.Set("cipres-appkey", auth.AppID)

	return req
}

/* Login and access server data based on the URL provided,
can get jobs list, status and results. */
func login(method string, url string) *http.Response {

	req := authenticate(method, url)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
