package gocra

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var auth Auth

/* Initializes the the authentication data and import job data
   and count the number of job on server side. */
func init() {

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	jsonFile, err := os.Open(home + "/cra_auth.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteValue, &auth)
}
