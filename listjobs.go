package gocra

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

/*
ListJobs GETs a list of jobs information from the server
and returns a JobList structure.
*/
func ListJobs() JobsList {

	resp := login("GET", auth.URL+"/job/"+auth.User, nil)

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var job JobsList
	xml.Unmarshal(b, &job)

	return job
}
