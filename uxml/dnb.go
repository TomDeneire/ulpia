package uxml

import (
	"encoding/xml"
)

// Struct for DNB SRU response
type DNB struct {
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
				RDF  struct {
					Text        string `xml:",chardata"`
					Schema      string `xml:"schema,attr"`
					Gndo        string `xml:"gndo,attr"`
					Lib         string `xml:"lib,attr"`
					Owl         string `xml:"owl,attr"`
					Xsd         string `xml:"xsd,attr"`
					Skos        string `xml:"skos,attr"`
					Rdfs        string `xml:"rdfs,attr"`
					Editeur     string `xml:"editeur,attr"`
					Geo         string `xml:"geo,attr"`
					Umbel       string `xml:"umbel,attr"`
					Rdau        string `xml:"rdau,attr"`
					Sf          string `xml:"sf,attr"`
					Bflc        string `xml:"bflc,attr"`
					Dcterms     string `xml:"dcterms,attr"`
					Isbd        string `xml:"isbd,attr"`
					Mesh        string `xml:"mesh,attr"`
					Foaf        string `xml:"foaf,attr"`
					Mo          string `xml:"mo,attr"`
					MarcRole    string `xml:"marcRole,attr"`
					Agrelon     string `xml:"agrelon,attr"`
					Dcmitype    string `xml:"dcmitype,attr"`
					Dbp         string `xml:"dbp,attr"`
					Dnbt        string `xml:"dnbt,attr"`
					Madsrdf     string `xml:"madsrdf,attr"`
					DnbIntern   string `xml:"dnb_intern,attr"`
					Rdf         string `xml:"rdf,attr"`
					V           string `xml:"v,attr"`
					Wdrs        string `xml:"wdrs,attr"`
					Ebu         string `xml:"ebu,attr"`
					Bibo        string `xml:"bibo,attr"`
					Gbv         string `xml:"gbv,attr"`
					Dc          string `xml:"dc,attr"`
					Description struct {
						Text  string `xml:",chardata"`
						About string `xml:"about,attr"`
						Type  struct {
							Text     string `xml:",chardata"`
							Resource string `xml:"resource,attr"`
						} `xml:"type"`
						Medium struct {
							Text     string `xml:",chardata"`
							Resource string `xml:"resource,attr"`
						} `xml:"medium"`
						P60049 []struct {
							Text     string `xml:",chardata"`
							Resource string `xml:"resource,attr"`
						} `xml:"P60049"`
						P60050 struct {
							Text     string `xml:",chardata"`
							Resource string `xml:"resource,attr"`
						} `xml:"P60050"`
						P60048 struct {
							Text     string `xml:",chardata"`
							Resource string `xml:"resource,attr"`
						} `xml:"P60048"`
						Identifier []struct {
							Text     string `xml:",chardata"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"identifier"`
						Isbn13 struct {
							Text     string `xml:",chardata"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"isbn13"`
						P60521 struct {
							Text     string `xml:",chardata"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"P60521"`
						Gtin14 struct {
							Text     string `xml:",chardata"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"gtin14"`
						Language struct {
							Text     string `xml:",chardata"`
							Resource string `xml:"resource,attr"`
						} `xml:"language"`
						P60327 struct {
							Text     string `xml:",chardata"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"P60327"`
						Title struct {
							Text     string `xml:",chardata"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"title"`
						Creator struct {
							Text     string `xml:",chardata"`
							Resource string `xml:"resource,attr"`
						} `xml:"creator"`
						Aut struct {
							Text     string `xml:",chardata"`
							Resource string `xml:"resource,attr"`
						} `xml:"aut"`
						Publisher struct {
							Text     string `xml:",chardata"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"publisher"`
						P60163 []struct {
							Text     string `xml:",chardata"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"P60163"`
						P60333 struct {
							Text     string `xml:",chardata"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"P60333"`
						P60539 struct {
							Text     string `xml:",chardata"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"P60539"`
						IsPartOf struct {
							Text     string `xml:",chardata"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"isPartOf"`
						P60470 struct {
							Text string `xml:",chardata"`
							Lang string `xml:"lang,attr"`
						} `xml:"P60470"`
						Subject []struct {
							Text     string `xml:",chardata"`
							Resource string `xml:"resource,attr"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"subject"`
						Describedby struct {
							Text        string `xml:",chardata"`
							Description struct {
								Text    string `xml:",chardata"`
								About   string `xml:"about,attr"`
								License struct {
									Text     string `xml:",chardata"`
									Resource string `xml:"resource,attr"`
								} `xml:"license"`
								Modified struct {
									Text     string `xml:",chardata"`
									Datatype string `xml:"datatype,attr"`
								} `xml:"modified"`
							} `xml:"Description"`
						} `xml:"describedby"`
						Issued struct {
							Text     string `xml:",chardata"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"issued"`
						P1053 struct {
							Text     string `xml:",chardata"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"P1053"`
						SameAs struct {
							Text     string `xml:",chardata"`
							Resource string `xml:"resource,attr"`
						} `xml:"sameAs"`
						Alternative struct {
							Text     string `xml:",chardata"`
							Datatype string `xml:"datatype,attr"`
						} `xml:"alternative"`
						Prt []struct {
							Text     string `xml:",chardata"`
							Resource string `xml:"resource,attr"`
						} `xml:"prt"`
						Pbl struct {
							Text     string `xml:",chardata"`
							Resource string `xml:"resource,attr"`
						} `xml:"pbl"`
					} `xml:"Description"`
				} `xml:"RDF"`
			} `xml:"recordData"`
			RecordPosition string `xml:"recordPosition"`
		} `xml:"record"`
	} `xml:"records"`
	EchoedSearchRetrieveRequest struct {
		Text    string `xml:",chardata"`
		Version string `xml:"version"`
		Query   string `xml:"query"`
		XQuery  struct {
			Text string `xml:",chardata"`
			Xsi  string `xml:"xsi,attr"`
			Nil  string `xml:"nil,attr"`
		} `xml:"xQuery"`
		StartRecord    string `xml:"startRecord"`
		MaximumRecords string `xml:"maximumRecords"`
	} `xml:"echoedSearchRetrieveRequest"`
}
