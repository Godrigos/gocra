package gocra

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

/*
Submit a job and returns a JobStatus object and pointer to a bytes.Buffer.

params is a map where each key is a Visible parameter or metadata
and the value is the string to assign to the parameter.

For example:
{
  "vparam.toolId" : "CLUSTALW",
  "vparam.runtime_" : "1",
}

path is a map where each key is an InFile parameter and the value
is the full path of the file.

For example:
{
  "input.infile_" : "/samplefiles/sample1_in.fasta,
  "input.usetree_" : "/samplefiles/guidetree.dnd",
}


See https://www.phylo.org/restusers/docs/guide.html#ConfigureParams
for more info on files, parameters and metadata accepted by each tool.

If you want to validate you submission before actually doing it,
set validate to true.
*/
func Submit(params map[string]string, path map[string]string,
	validate bool) (JobStatus, *bytes.Buffer) {

	uri := auth.URL + "/job/" + auth.User
	if validate == true {
		uri += "/validate"
	}

	var file *os.File
	body := &bytes.Buffer{}
	bo := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	var part interface{ io.Writer }
	var err error

	for key := range path {

		file, err = os.Open(path[key])
		if err != nil {
			panic(err)
		}
		defer file.Close()

		part, err = writer.CreateFormFile(key, filepath.Base(path[key]))
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			panic(err)
		}
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", uri, body)
	req.SetBasicAuth(auth.User, auth.Passwd)
	req.Header.Set("cipres-appkey", auth.AppID)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if validate == true {
		_, err := bo.ReadFrom(resp.Body)
		if err != nil {
			panic(err)
		}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var job JobStatus
	xml.Unmarshal(b, &job)

	return job, bo
}
