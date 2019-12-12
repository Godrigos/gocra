package gocra

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var auth Auth

/*
Initializes the the authentication data and import job data
and count the number of job on server side.
*/
func init() {

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Open(home + "/cra_auth.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(byteValue, &auth)
}
