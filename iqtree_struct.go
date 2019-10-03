package gocra

import "encoding/xml"

/*
IQTree parameters structure
Usage: /projects/ps-ngbt/opt/comet/iqtree-omp-1.5.4-Linux/bin/iqtree-omp -s <alignment> [OPTIONS]
Details at http://www.phylo.org/index.php/rest/iqtree_xsede.html
*/
type IQTree struct {
	XMLName xml.Name `xml:"pise"`
	Text    string   `xml:",chardata"`
	Head    struct {
		Text        string `xml:",chardata"`
		Title       string `xml:"title"`
		Version     string `xml:"version"`
		Description string `xml:"description"`
		Authors     string `xml:"authors"`
		Reference   string `xml:"reference"`
		Category    string `xml:"category"`
	} `xml:"head"`
	Command    string `xml:"command"`
	Parameters struct {
		Text      string `xml:",chardata"`
		Parameter []struct {
			Text        string `xml:",chardata"`
			Ismandatory string `xml:"ismandatory,attr"`
			Ishidden    string `xml:"ishidden,attr"`
			Type        string `xml:"type,attr"`
			Issimple    string `xml:"issimple,attr"`
			Isinput     string `xml:"isinput,attr"`
			Name        string `xml:"name"`
			Attributes  struct {
				Text    string `xml:",chardata"`
				Precond struct {
					Text     string `xml:",chardata"`
					Language string `xml:"language"`
					Code     string `xml:"code"`
				} `xml:"precond"`
				Format struct {
					Text     string `xml:",chardata"`
					Language string `xml:"language"`
					Code     string `xml:"code"`
				} `xml:"format"`
				Group     string `xml:"group"`
				Paramfile string `xml:"paramfile"`
				Prompt    string `xml:"prompt"`
				Filenames string `xml:"filenames"`
				Vdef      struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value"`
				} `xml:"vdef"`
				Ctrls struct {
					Text string `xml:",chardata"`
					Ctrl []struct {
						Text     string `xml:",chardata"`
						Message  string `xml:"message"`
						Language string `xml:"language"`
						Code     string `xml:"code"`
					} `xml:"ctrl"`
				} `xml:"ctrls"`
				Warns struct {
					Text string `xml:",chardata"`
					Warn []struct {
						Text     string `xml:",chardata"`
						Message  string `xml:"message"`
						Language string `xml:"language"`
						Code     string `xml:"code"`
					} `xml:"warn"`
				} `xml:"warns"`
				Comment struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value"`
				} `xml:"comment"`
				Vlist struct {
					Text  string   `xml:",chardata"`
					Value []string `xml:"value"`
					Label []string `xml:"label"`
				} `xml:"vlist"`
			} `xml:"attributes"`
			Paragraph struct {
				Text       string `xml:",chardata"`
				Name       string `xml:"name"`
				Prompt     string `xml:"prompt"`
				Parameters struct {
					Text      string `xml:",chardata"`
					Parameter []struct {
						Text       string `xml:",chardata"`
						Issimple   string `xml:"issimple,attr"`
						Type       string `xml:"type,attr"`
						Ishidden   string `xml:"ishidden,attr"`
						Name       string `xml:"name"`
						Attributes struct {
							Text   string `xml:",chardata"`
							Prompt string `xml:"prompt"`
							Vlist  struct {
								Text  string   `xml:",chardata"`
								Value []string `xml:"value"`
								Label []string `xml:"label"`
							} `xml:"vlist"`
							Format struct {
								Text     string `xml:",chardata"`
								Language string `xml:"language"`
								Code     string `xml:"code"`
							} `xml:"format"`
							Group   string `xml:"group"`
							Precond struct {
								Text     string `xml:",chardata"`
								Language string `xml:"language"`
								Code     string `xml:"code"`
							} `xml:"precond"`
							Comment struct {
								Text  string `xml:",chardata"`
								Value string `xml:"value"`
							} `xml:"comment"`
							Ctrls struct {
								Text string `xml:",chardata"`
								Ctrl []struct {
									Text     string `xml:",chardata"`
									Message  string `xml:"message"`
									Language string `xml:"language"`
									Code     string `xml:"code"`
								} `xml:"ctrl"`
							} `xml:"ctrls"`
							Filenames string `xml:"filenames"`
							Vdef      struct {
								Text  string `xml:",chardata"`
								Value string `xml:"value"`
							} `xml:"vdef"`
						} `xml:"attributes"`
					} `xml:"parameter"`
				} `xml:"parameters"`
			} `xml:"paragraph"`
		} `xml:"parameter"`
	} `xml:"parameters"`
}
