package uxml

import (
	"encoding/xml"
)

// Struct for Anet SRU response
type Anet struct {
	XMLName         xml.Name `xml:"searchRetrieveResponse"`
	Text            string   `xml:",chardata"`
	Zs              string   `xml:"zs,attr"`
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
					Text           string `xml:",chardata"`
					SrwDc          string `xml:"srw_dc,attr"`
					Xsi            string `xml:"xsi,attr"`
					SchemaLocation string `xml:"schemaLocation,attr"`
					Title          struct {
						Text  string `xml:",chardata"`
						Xmlns string `xml:"xmlns,attr"`
					} `xml:"title"`
					Creator []struct {
						Text  string `xml:",chardata"`
						Xmlns string `xml:"xmlns,attr"`
					} `xml:"creator"`
					Type struct {
						Text  string `xml:",chardata"`
						Xmlns string `xml:"xmlns,attr"`
					} `xml:"type"`
					Publisher struct {
						Text  string `xml:",chardata"`
						Xmlns string `xml:"xmlns,attr"`
					} `xml:"publisher"`
					Language struct {
						Text  string `xml:",chardata"`
						Xmlns string `xml:"xmlns,attr"`
					} `xml:"language"`
					Description []struct {
						Text  string `xml:",chardata"`
						Xmlns string `xml:"xmlns,attr"`
					} `xml:"description"`
					Subject []struct {
						Text  string `xml:",chardata"`
						Xmlns string `xml:"xmlns,attr"`
					} `xml:"subject"`
					Identifier []struct {
						Text  string `xml:",chardata"`
						Xmlns string `xml:"xmlns,attr"`
					} `xml:"identifier"`
					Relation []struct {
						Text  string `xml:",chardata"`
						Xmlns string `xml:"xmlns,attr"`
					} `xml:"relation"`
					Format struct {
						Text  string `xml:",chardata"`
						Xmlns string `xml:"xmlns,attr"`
					} `xml:"format"`
					Coverage struct {
						Text  string `xml:",chardata"`
						Xmlns string `xml:"xmlns,attr"`
					} `xml:"coverage"`
				} `xml:"dc"`
			} `xml:"recordData"`
			RecordPosition string `xml:"recordPosition"`
		} `xml:"record"`
	} `xml:"records"`
	EchoedSearchRetrieveRequest struct {
		Text           string `xml:",chardata"`
		Version        string `xml:"version"`
		Query          string `xml:"query"`
		StartRecord    string `xml:"startRecord"`
		MaximumRecords string `xml:"maximumRecords"`
		RecordPacking  string `xml:"recordPacking"`
		RecordSchema   string `xml:"recordSchema"`
	} `xml:"echoedSearchRetrieveRequest"`
}
