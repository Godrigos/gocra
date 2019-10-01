package gocra

import (
	"fmt"
)

/*
Delete send a request to delete a completed job.
Must be only executes after the results have been already downloaded.
*/
func Delete(url string) {

	resp := login("DELETE", url)

	if resp.StatusCode == 204 || resp.StatusCode == 200 {
		fmt.Println("Job deleted from Cipres.")
	} else {
		fmt.Println(resp.Status)
	}
}
