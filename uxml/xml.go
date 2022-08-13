package uxml

import (
	"encoding/xml"
)

type Result struct {
	Titles      []string
	Dates       []string
	Authors     []string
	Identifiers []string
	Imprints    []string
}

// Parse XML responses according to specification
func ParseXML(data []byte, source string) (Result, error) {
	var result Result

	if source == "dnb" {
		var response DNB
		_ = xml.Unmarshal(data, &response)
		for _, record := range response.Records.Record {
			result.Titles = append(result.Titles, record.RecordData.RDF.Description.Title.Text)
			identifiers := ""
			for _, identifier := range record.RecordData.RDF.Description.Identifier {
				identifiers = identifiers + "\n" + identifier.Text
			}
			result.Identifiers = append(result.Identifiers, identifiers)
			result.Dates = append(result.Dates, record.RecordData.RDF.Description.Issued.Text)
			result.Authors = append(result.Authors, record.RecordData.RDF.Description.P60327.Text)
			result.Imprints = append(result.Imprints, record.RecordData.RDF.Description.P60333.Text)
		}
		return result, nil
	} else if source == "hpb" {
		var response HPB
		_ = xml.Unmarshal(data, &response)
		for _, record := range response.Records.Record {

			identifiers := ""
			for _, identifier := range record.RecordData.Dc.Identifier {
				identifiers = identifiers + "\n" + identifier.Text
			}
			result.Identifiers = append(result.Identifiers, identifiers)

			titles := ""
			for _, title := range record.RecordData.Dc.Title {
				titles = titles + "\n" + title.Text
			}
			result.Titles = append(result.Titles, titles)

			dates := ""
			for _, date := range record.RecordData.Dc.Date {
				dates = dates + "\n" + date.Text
			}
			result.Dates = append(result.Dates, dates)

			authors := ""
			for _, author := range record.RecordData.Dc.Contributor {
				authors = authors + "\n" + author.Text
			}
			result.Authors = append(result.Authors, authors)

			result.Imprints = append(result.Imprints, record.RecordData.Dc.Publisher.Text)
		}
		return result, nil
	} else if source == "unicat" {
		var response Unicat
		_ = xml.Unmarshal(data, &response)
		for _, record := range response.Records.Record {

			identifiers := ""
			for _, identifier := range record.RecordData.Dc.Identifier {
				identifiers = identifiers + "\n" + identifier
			}
			result.Identifiers = append(result.Identifiers, identifiers)

			result.Titles = append(result.Titles, record.RecordData.Dc.Title)

			result.Dates = append(result.Dates, record.RecordData.Dc.Date)

			authors := ""
			for _, author := range record.RecordData.Dc.Creator {
				authors = authors + "\n" + author
			}
			result.Authors = append(result.Authors, authors)

			result.Imprints = append(result.Imprints, record.RecordData.Dc.Publisher)
		}
		return result, nil
	}

	return result, nil
}
