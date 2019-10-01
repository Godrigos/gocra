package gocra

import (
	"fmt"
	"strings"
)

// Cancel send a request to cancel a running job without deleting files.
func Cancel(jh string) {

	url := auth.URL + "/job/" + auth.User + "/modify/" + jh

	resp := login("PUT", url, strings.NewReader("action=cancel"))

	if resp.StatusCode == 200 || resp.StatusCode == 202 ||
		resp.StatusCode == 204 {
		fmt.Println("Job canceled.")
	} else {
		fmt.Println(resp.Status)
	}
}
