package uxml

import (
	"encoding/xml"
)

// Struct for Unicat SRU response
type Unicat struct {
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
					SrwDc          string   `xml:"srw_dc,attr"`
					Xsi            string   `xml:"xsi,attr"`
					Xmlns          string   `xml:"xmlns,attr"`
					SchemaLocation string   `xml:"schemaLocation,attr"`
					Title          string   `xml:"title"`
					Identifier     []string `xml:"identifier"`
					Creator        []string `xml:"creator"`
					Subject        string   `xml:"subject"`
					Language       string   `xml:"language"`
					Publisher      string   `xml:"publisher"`
					Date           string   `xml:"date"`
				} `xml:"dc"`
			} `xml:"recordData"`
			RecordPosition string `xml:"recordPosition"`
		} `xml:"record"`
	} `xml:"records"`
	ExtraResponseData           string `xml:"extraResponseData"`
	EchoedSearchRetrieveRequest struct {
		Text           string `xml:",chardata"`
		Version        string `xml:"version"`
		Query          string `xml:"query"`
		StartRecord    string `xml:"startRecord"`
		MaximumRecords string `xml:"maximumRecords"`
		RecordSchema   string `xml:"recordSchema"`
		XQuery         struct {
			Text   string `xml:",chardata"`
			Triple struct {
				Text    string `xml:",chardata"`
				Xmlns   string `xml:"xmlns,attr"`
				Boolean struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value"`
				} `xml:"boolean"`
				LeftOperand struct {
					Text   string `xml:",chardata"`
					Triple struct {
						Text    string `xml:",chardata"`
						Boolean struct {
							Text  string `xml:",chardata"`
							Value string `xml:"value"`
						} `xml:"boolean"`
						LeftOperand struct {
							Text         string `xml:",chardata"`
							SearchClause struct {
								Text     string `xml:",chardata"`
								Index    string `xml:"index"`
								Relation struct {
									Text  string `xml:",chardata"`
									Value string `xml:"value"`
								} `xml:"relation"`
								Term string `xml:"term"`
							} `xml:"searchClause"`
						} `xml:"leftOperand"`
						RightOperand struct {
							Text         string `xml:",chardata"`
							SearchClause struct {
								Text     string `xml:",chardata"`
								Index    string `xml:"index"`
								Relation struct {
									Text  string `xml:",chardata"`
									Value string `xml:"value"`
								} `xml:"relation"`
								Term string `xml:"term"`
							} `xml:"searchClause"`
						} `xml:"rightOperand"`
					} `xml:"triple"`
				} `xml:"leftOperand"`
				RightOperand struct {
					Text         string `xml:",chardata"`
					SearchClause struct {
						Text     string `xml:",chardata"`
						Index    string `xml:"index"`
						Relation struct {
							Text  string `xml:",chardata"`
							Value string `xml:"value"`
						} `xml:"relation"`
						Term string `xml:"term"`
					} `xml:"searchClause"`
				} `xml:"rightOperand"`
			} `xml:"triple"`
		} `xml:"xQuery"`
	} `xml:"echoedSearchRetrieveRequest"`
}
