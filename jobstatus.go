package gocra

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

/*
JobStat GETs a job status information from CIPRES server
based on a job handle provided by listJobs function and returns
a JobStatus structure
*/
func JobStat(jb JobsList, i int) JobStatus {

	resp := login("GET", jb.Jobs.Jobstatus[i].SelfURI.URL)

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var job JobStatus
	xml.Unmarshal(b, &job)
	return job
}
