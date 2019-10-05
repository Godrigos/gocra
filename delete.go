package gocra

import (
	"fmt"
)

/*
Delete send a request to delete a completed job and its corresponding files.
Must be only executes after the results have been already downloaded.
*/
func Delete(url string) {

	resp := login("DELETE", url, nil)

	if resp.StatusCode == 200 || resp.StatusCode == 202 ||
		resp.StatusCode == 204 {
		fmt.Println("Job deleted from Cipres.")
	} else {
		fmt.Println(resp.Status)
	}
}
