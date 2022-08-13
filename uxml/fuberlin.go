package uxml

import "encoding/xml"

type FUBerlin struct {
	XMLName         xml.Name `xml:"searchRetrieveResponse"`
	Text            string   `xml:",chardata"`
	Xmlns           string   `xml:"xmlns,attr"`
	Version         string   `xml:"version"`
	NumberOfRecords string   `xml:"numberOfRecords"`
	Records         struct {
		Text   string `xml:",chardata"`
		Record []struct {
			Text          string `xml:",chardata"`
			RecordSchema  string `xml:"recordSchema"`
			RecordPacking string `xml:"recordPacking"`
			RecordData    struct {
				Text string `xml:",chardata"`
				Dc   struct {
					Text           string   `xml:",chardata"`
					Dc             string   `xml:"dc,attr"`
					SrwDc          string   `xml:"srw_dc,attr"`
					Xsi            string   `xml:"xsi,attr"`
					SchemaLocation string   `xml:"schemaLocation,attr"`
					Title          []string `xml:"title"`
					Contributor    []string `xml:"contributor"`
					Type           string   `xml:"type"`
					Publisher      []string `xml:"publisher"`
					Date           []string `xml:"date"`
					Language       []string `xml:"language"`
					Description    []string `xml:"description"`
					Subject        []string `xml:"subject"`
					Coverage       []string `xml:"coverage"`
					Identifier     []string `xml:"identifier"`
				} `xml:"dc"`
			} `xml:"recordData"`
			RecordIdentifier string `xml:"recordIdentifier"`
			RecordPosition   string `xml:"recordPosition"`
		} `xml:"record"`
	} `xml:"records"`
	ExtraResponseData struct {
		Text         string `xml:",chardata"`
		Xb           string `xml:"xb,attr"`
		Exact        string `xml:"exact"`
		ResponseDate string `xml:"responseDate"`
	} `xml:"extraResponseData"`
}
