package uxml

import (
	"encoding/xml"
)

// Struct for HPB SRU response
type HPB struct {
	XMLName         xml.Name `xml:"searchRetrieveResponse"`
	Text            string   `xml:",chardata"`
	Zs              string   `xml:"zs,attr"`
	NumberOfRecords string   `xml:"numberOfRecords"`
	Records         struct {
		Text   string `xml:",chardata"`
		Record []struct {
			Text              string `xml:",chardata"`
			RecordSchema      string `xml:"recordSchema"`
			RecordXMLEscaping string `xml:"recordXMLEscaping"`
			RecordData        struct {
				Text string `xml:",chardata"`
				Dc   struct {
					Text           string `xml:",chardata"`
					OaiDc          string `xml:"oai_dc,attr"`
					Xsi            string `xml:"xsi,attr"`
					SchemaLocation string `xml:"schemaLocation,attr"`
					Title          []struct {
						Text  string `xml:",chardata"`
						Dc    string `xml:"dc,attr"`
						SrwDc string `xml:"srw_dc,attr"`
					} `xml:"title"`
					Contributor []struct {
						Text string `xml:",chardata"`
						Dc   string `xml:"dc,attr"`
					} `xml:"contributor"`
					Type struct {
						Text string `xml:",chardata"`
						Dc   string `xml:"dc,attr"`
					} `xml:"type"`
					Date []struct {
						Text  string `xml:",chardata"`
						Dc    string `xml:"dc,attr"`
						SrwDc string `xml:"srw_dc,attr"`
					} `xml:"date"`
					Language []struct {
						Text  string `xml:",chardata"`
						Dc    string `xml:"dc,attr"`
						SrwDc string `xml:"srw_dc,attr"`
					} `xml:"language"`
					Format []struct {
						Text string `xml:",chardata"`
						Dc   string `xml:"dc,attr"`
					} `xml:"format"`
					Description []struct {
						Text  string `xml:",chardata"`
						Dc    string `xml:"dc,attr"`
						SrwDc string `xml:"srw_dc,attr"`
					} `xml:"description"`
					Identifier []struct {
						Text  string `xml:",chardata"`
						Dc    string `xml:"dc,attr"`
						SrwDc string `xml:"srw_dc,attr"`
					} `xml:"identifier"`
					Publisher struct {
						Text string `xml:",chardata"`
						Dc   string `xml:"dc,attr"`
					} `xml:"publisher"`
					Relation struct {
						Text string `xml:",chardata"`
						Dc   string `xml:"dc,attr"`
					} `xml:"relation"`
				} `xml:"dc"`
			} `xml:"recordData"`
			RecordPosition string `xml:"recordPosition"`
		} `xml:"record"`
	} `xml:"records"`
	EchoedSearchRetrieveRequest struct {
		Text              string `xml:",chardata"`
		Version           string `xml:"version"`
		Query             string `xml:"query"`
		StartRecord       string `xml:"startRecord"`
		MaximumRecords    string `xml:"maximumRecords"`
		RecordXMLEscaping string `xml:"recordXMLEscaping"`
		RecordSchema      string `xml:"recordSchema"`
	} `xml:"echoedSearchRetrieveRequest"`
	ResultCountPrecision string `xml:"resultCountPrecision"`
}
