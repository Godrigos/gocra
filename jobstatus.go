package gocra

import (
	"encoding/xml"
	"io/ioutil"
)

/*
JobStat GETs a job handle url (as s string) information from CIPRES server
provided by listJobs function and returns a JobStatus structure
*/
func JobStat(jh string) JobStatus {

	resp := login("GET", jh, nil)

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}

	var job JobStatus
	xml.Unmarshal(b, &job)
	return job
}
