package gocra

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

/*
ListJobs GETs a list of jobs information from the server
and returns a slice of JobList structures.
*/
func ListJobs() []JobsList {

	resp := login("GET", auth.URL+"/job/"+auth.User)

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var job JobsList
	var jobs []JobsList
	xml.Unmarshal(b, &job)

	for range job.Jobs.Jobstatus {
		jobs = append(jobs, job)
	}
	return jobs
}
