package uxml

import (
	"encoding/xml"

	tools "tomdeneire.github.io/ulpia/tools"
	uhtml "tomdeneire.github.io/ulpia/uhtml"
)

// Parse XML responses according to specification
func ParseXML(data []byte, source string) (uhtml.Result, error) {
	var result uhtml.Result

	if source == "dnb" {
		var response DNB
		err := xml.Unmarshal(data, &response)
		if err != nil {
			return result, err
		}
		for _, record := range response.Records.Record {
			result.Titles = append(result.Titles,
				record.RecordData.RDF.Description.Title.Text)

			result.Identifiers = append(result.Identifiers,
				loopConCatS3(record.RecordData.RDF.Description.Identifier))

			result.Dates = append(result.Dates,
				record.RecordData.RDF.Description.Issued.Text)

			result.Authors = append(result.Authors,
				record.RecordData.RDF.Description.P60327.Text)

			result.Imprints = append(result.Imprints,
				record.RecordData.RDF.Description.P60333.Text)
		}
		return result, nil
	} else if source == "hpb" {
		var response HPB
		err := xml.Unmarshal(data, &response)
		if err != nil {
			return result, err
		}
		for _, record := range response.Records.Record {

			result.Identifiers = append(result.Identifiers,
				loopConCatS(record.RecordData.Dc.Identifier))

			result.Titles = append(result.Titles,
				loopConCatS(record.RecordData.Dc.Title))

			result.Dates = append(result.Dates,
				loopConCatS(record.RecordData.Dc.Date))

			result.Authors = append(result.Authors,
				loopConCatS2(record.RecordData.Dc.Contributor))

			result.Imprints = append(result.Imprints,
				record.RecordData.Dc.Publisher.Text)
		}
		return result, nil
	} else if source == "unicat" {
		var response Unicat
		err := xml.Unmarshal(data, &response)
		if err != nil {
			return result, err
		}
		for _, record := range response.Records.Record {

			result.Identifiers = append(result.Identifiers,
				tools.LoopConCat(record.RecordData.Dc.Identifier))

			result.Titles = append(result.Titles,
				record.RecordData.Dc.Title)

			result.Dates = append(result.Dates,
				record.RecordData.Dc.Date)

			result.Authors = append(result.Authors,
				tools.LoopConCat(record.RecordData.Dc.Creator))

			result.Imprints = append(result.Imprints,
				record.RecordData.Dc.Publisher)
		}
		return result, nil
	} else if source == "fu-berlin" {
		var response FUBerlin
		err := xml.Unmarshal(data, &response)
		if err != nil {
			return result, err
		}
		for _, record := range response.Records.Record {
			result.Identifiers = append(result.Identifiers,
				tools.LoopConCat(record.RecordData.Dc.Identifier))

			result.Titles = append(result.Titles,
				tools.LoopConCat(record.RecordData.Dc.Title))

			result.Dates = append(result.Dates,
				tools.LoopConCat(record.RecordData.Dc.Date))

			result.Authors = append(result.Authors,
				tools.LoopConCat(record.RecordData.Dc.Contributor))

			result.Imprints = append(result.Imprints,
				tools.LoopConCat(record.RecordData.Dc.Publisher))

		}
		return result, nil
	} else if source == "bnf" {
		var response BNF
		err := xml.Unmarshal(data, &response)
		if err != nil {
			return result, err
		}
		for _, record := range response.Records.Record {
			result.Identifiers = append(result.Identifiers,
				record.RecordData.Dc.Identifier)

			result.Titles = append(result.Titles,
				record.RecordData.Dc.Title)

			result.Dates = append(result.Dates,
				record.RecordData.Dc.Date)

			result.Authors = append(result.Authors,
				record.RecordData.Dc.Creator)

			result.Imprints = append(result.Imprints,
				tools.LoopConCat(record.RecordData.Dc.Publisher))

		}
		return result, nil
	}
	return result, nil
}

func loopConCatS(data []struct {
	Text  string "xml:\",chardata\""
	Dc    string "xml:\"dc,attr\""
	SrwDc string "xml:\"srw_dc,attr\""
}) string {
	var result string
	for _, item := range data {
		result += "\n" + item.Text
	}
	return result
}

func loopConCatS2(data []struct {
	Text string "xml:\",chardata\""
	Dc   string "xml:\"dc,attr\""
}) string {
	var result string
	for _, item := range data {
		result += "\n" + item.Text
	}
	return result
}

func loopConCatS3(data []struct {
	Text     string "xml:\",chardata\""
	Datatype string "xml:\"datatype,attr\""
}) string {
	var result string
	for _, item := range data {
		result += "\n" + item.Text
	}
	return result
}
