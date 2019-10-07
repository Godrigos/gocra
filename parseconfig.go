package gocra

import (
	"bufio"
	"log"
	"os"
	"strings"
)

/*
ParseConfig can extract Visible Parameters and Metadata from a plain text file
and return a map ready to use for submission.
It is usefull to use with data copied from Cipres Tool Configuration Helper
(https://www.phylo.org/restusers/docs/cipresXml) output
and saved to a text file.

The data should be in the following format: tool = IQTREE_XSEDE, with each
parameter or metadata in a different line.

Infile Parameters should be entered separately as the file parameter of submit
function, once they are processed differently by the function.
*/
func ParseConfig(path string) map[string]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	config := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		config[(strings.TrimSpace(strings.Split(scanner.Text(), "=")[0]))] =
			strings.TrimSpace(strings.Split(scanner.Text(), "=")[1])
	}
	return config
}
