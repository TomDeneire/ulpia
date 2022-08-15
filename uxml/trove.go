package uxml

import (
	"encoding/xml"
)

// Struct for Trove XML response
type Trove struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Query   string   `xml:"query"`
	Zone    struct {
		Text    string `xml:",chardata"`
		Name    string `xml:"name,attr"`
		Records struct {
			Text  string `xml:",chardata"`
			S     string `xml:"s,attr"`
			N     string `xml:"n,attr"`
			Total string `xml:"total,attr"`
			Work  []struct {
				Text           string `xml:",chardata"`
				ID             string `xml:"id,attr"`
				URL            string `xml:"url,attr"`
				TroveUrl       string `xml:"troveUrl"`
				Title          string `xml:"title"`
				Contributor    string `xml:"contributor"`
				Issued         string `xml:"issued"`
				Type           string `xml:"type"`
				HoldingsCount  string `xml:"holdingsCount"`
				VersionCount   string `xml:"versionCount"`
				HasCorrections string `xml:"hasCorrections"`
				Relevance      struct {
					Text  string `xml:",chardata"`
					Score string `xml:"score,attr"`
				} `xml:"relevance"`
				Snippet []string `xml:"snippet"`
			} `xml:"work"`
		} `xml:"records"`
	} `xml:"zone"`
}
