package gocra

import "encoding/xml"

// Auth stores the authentication data imported from cra_auth.json
type Auth struct {
	URL    string `json:"url"`
	AppID  string `json:"appid"`
	User   string `json:"username"`
	Passwd string `json:"password"`
}

// JobsList stores jobs list information retrieved from CIPRES server
type JobsList struct {
	XMLName xml.Name `xml:"joblist"`
	Text    string   `xml:",chardata"`
	Title   string   `xml:"title"`
	Jobs    struct {
		Text      string `xml:",chardata"`
		Jobstatus []struct {
			Text    string `xml:",chardata"`
			SelfURI struct {
				Text  string `xml:",chardata"`
				URL   string `xml:"url"`
				Rel   string `xml:"rel"`
				Title string `xml:"title"`
			} `xml:"selfUri"`
		} `xml:"jobstatus"`
	} `xml:"jobs"`
}

// JobStatus stores a job status information retrieved from CIPRES server
type JobStatus struct {
	XMLName xml.Name `xml:"jobstatus"`
	Text    string   `xml:",chardata"`
	SelfURI struct {
		Text  string `xml:",chardata"`
		URL   string `xml:"url"`
		Rel   string `xml:"rel"`
		Title string `xml:"title"`
	} `xml:"selfUri"`
	JobHandle     string `xml:"jobHandle"`
	JobStage      string `xml:"jobStage"`
	TerminalStage bool   `xml:"terminalStage"`
	Failed        string `xml:"failed"`
	Metadata      struct {
		Text  string `xml:",chardata"`
		Entry struct {
			Text  string `xml:",chardata"`
			Key   string `xml:"key"`
			Value string `xml:"value"`
		} `xml:"entry"`
	} `xml:"metadata"`
	DateSubmitted string `xml:"dateSubmitted"`
	ResultsURI    struct {
		Text  string `xml:",chardata"`
		URL   string `xml:"url"`
		Rel   string `xml:"rel"`
		Title string `xml:"title"`
	} `xml:"resultsUri"`
	WorkingDirURI struct {
		Text  string `xml:",chardata"`
		URL   string `xml:"url"`
		Rel   string `xml:"rel"`
		Title string `xml:"title"`
	} `xml:"workingDirUri"`
	Messages struct {
		Text    string `xml:",chardata"`
		Message struct {
			Chardata  string `xml:",chardata"`
			Timestamp string `xml:"timestamp"`
			Stage     string `xml:"stage"`
			Text      string `xml:"text"`
		} `xml:"message"`
	} `xml:"messages"`
	MinPollIntervalSeconds int `xml:"minPollIntervalSeconds"`
}

/*
JobResults stores a finished job result information retrieved
from CIPRES server
*/
type JobResults struct {
	XMLName  xml.Name `xml:"results"`
	Text     string   `xml:",chardata"`
	Jobfiles struct {
		Text    string `xml:",chardata"`
		Jobfile []struct {
			Text        string `xml:",chardata"`
			DownloadURI struct {
				Text  string `xml:",chardata"`
				URL   string `xml:"url"`
				Rel   string `xml:"rel"`
				Title string `xml:"title"`
			} `xml:"downloadUri"`
			JobHandle        string `xml:"jobHandle"`
			Filename         string `xml:"filename"`
			Length           int    `xml:"length"`
			ParameterName    string `xml:"parameterName"`
			OutputDocumentID string `xml:"outputDocumentId"`
		} `xml:"jobfile"`
	} `xml:"jobfiles"`
}

// WorkingDir stores a running job files list retrieved from CIPRES server
type WorkingDir struct {
	XMLName  xml.Name `xml:"workingdir"`
	Text     string   `xml:",chardata"`
	Jobfiles struct {
		Text    string `xml:",chardata"`
		Jobfile []struct {
			Text        string `xml:",chardata"`
			DownloadURI struct {
				Text  string `xml:",chardata"`
				URL   string `xml:"url"`
				Rel   string `xml:"rel"`
				Title string `xml:"title"`
			} `xml:"downloadUri"`
			Filename         string `xml:"filename"`
			Length           int    `xml:"length"`
			DateModified     string `xml:"dateModified"`
			ParameterName    string `xml:"parameterName"`
			OutputDocumentID string `xml:"outputDocumentId"`
		} `xml:"jobfile"`
	} `xml:"jobfiles"`
}
