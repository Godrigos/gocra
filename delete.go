package gocra

import (
	"fmt"
)

/*
Delete send a requesto to delete a completed job.
Must be only executes after the results have been already downloaded.
*/
func Delete(url string) {

	login("DELETE", url)

	fmt.Println("Job deleted from Cipres.")
}
