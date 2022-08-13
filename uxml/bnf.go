package uxml

import (
	"encoding/xml"
)

// Struct for BNF SRU response
type BNF struct {
	XMLName                     xml.Name `xml:"searchRetrieveResponse"`
	Text                        string   `xml:",chardata"`
	Srw                         string   `xml:"srw,attr"`
	Xmlns                       string   `xml:"xmlns,attr"`
	Ixm                         string   `xml:"ixm,attr"`
	Mn                          string   `xml:"mn,attr"`
	Sd                          string   `xml:"sd,attr"`
	Version                     string   `xml:"version"`
	EchoedSearchRetrieveRequest struct {
		Text    string `xml:",chardata"`
		Version string `xml:"version"`
		Query   string `xml:"query"`
	} `xml:"echoedSearchRetrieveRequest"`
	NumberOfRecords string `xml:"numberOfRecords"`
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
					OaiDc          string   `xml:"oai_dc,attr"`
					Dc             string   `xml:"dc,attr"`
					Xsi            string   `xml:"xsi,attr"`
					SchemaLocation string   `xml:"schemaLocation,attr"`
					Identifier     string   `xml:"identifier"`
					Title          string   `xml:"title"`
					Creator        string   `xml:"creator"`
					Publisher      []string `xml:"publisher"`
					Date           string   `xml:"date"`
					Format         string   `xml:"format"`
					Language       []string `xml:"language"`
					Type           []struct {
						Text string `xml:",chardata"`
						Lang string `xml:"lang,attr"`
					} `xml:"type"`
					Rights []struct {
						Text string `xml:",chardata"`
						Lang string `xml:"lang,attr"`
					} `xml:"rights"`
				} `xml:"dc"`
			} `xml:"recordData"`
			RecordIdentifier string `xml:"recordIdentifier"`
			RecordPosition   string `xml:"recordPosition"`
			ExtraRecordData  struct {
				Text string `xml:",chardata"`
				Attr []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"attr"`
				Score string `xml:"score"`
			} `xml:"extraRecordData"`
		} `xml:"record"`
	} `xml:"records"`
	NextRecordPosition string `xml:"nextRecordPosition"`
}
