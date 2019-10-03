package gocra

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"time"
)

/*
Submit a job and returns a JobStatus object.

vparams is a map where each key is a Visible parameter name, value
is the string value to assign the parameter.
For example: {"toolId" : "CLUSTALW", "runtime_" : "1"}

inputParams is a map where each key is an InFile parameter name and
the value is the full path of the file.
For example: {"infile_" : "/samplefiles/sample1_in.fasta",
              "usetree_" : "/samplefiles/guidetree.dnd"}

See https://www.phylo.org/restusers/docs/guide.html#ConfigureParams
for more info.
*/
func Submit(inputParams map[string]string,
	vParams map[string]string, validate bool) JobStatus {

	url := auth.URL + "/job/" + auth.User
	payload := make([]map[string]string, 0, 3)

	if validate == true {
		url += "/validate"
	}

	metadata := map[string]string{
		"metadata.clientJobId": vParams["input.infile_"] + "_" +
			time.Now().Format("2006-01-02T15:04:05"),
		"metadata.clientJobName": vParams["input.infile_"],
		"metadata.statusEmail":   "true",
	}

	payload = append(payload, inputParams, vParams, metadata)

	resp := login("POST", url, nil)

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var job JobStatus
	xml.Unmarshal(b, &job)

	return job
}
