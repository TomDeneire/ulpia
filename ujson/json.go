package ujson

import (
	"encoding/json"
	"strconv"

	tools "tomdeneire.github.io/ulpia/tools"
	uhtml "tomdeneire.github.io/ulpia/uhtml"
)

// Parse JSON responses according to specification
func ParseJSON(data []byte, source string) (uhtml.Result, error) {
	var result uhtml.Result

	if source == "europeana" {
		var response Europeana
		err := json.Unmarshal(data, &response)
		if err != nil {
			return result, err
		}
		for _, record := range response.Items {
			result.Titles = append(result.Titles,
				tools.LoopConCat(record.Title))
			result.Identifiers = append(result.Identifiers,
				record.ID)

			result.Dates = append(result.Dates,
				tools.LoopConCat(record.Year))

			result.Authors = append(result.Authors,
				tools.LoopConCat(record.DcCreator))

			result.Imprints = append(result.Imprints,
				"")
		}
		return result, nil
	} else if source == "openlibrary" {
		var response OpenLibrary
		err := json.Unmarshal(data, &response)
		if err != nil {
			return result, err
		}
		for _, record := range response.Docs {
			result.Titles = append(result.Titles,
				record.Title)
			result.Identifiers = append(result.Identifiers,
				record.Key)

			result.Dates = append(result.Dates,
				strconv.Itoa(record.FirstPublishYear))

			result.Authors = append(result.Authors,
				tools.LoopConCat(record.AuthorName))

			result.Imprints = append(result.Imprints,
				tools.LoopConCat(record.PublishPlace)+
					tools.LoopConCat(record.Publisher))
		}
		return result, nil
	} else if source == "googlebooks" {
		var response GoogleBooks
		err := json.Unmarshal(data, &response)
		if err != nil {
			return result, err
		}
		for _, record := range response.Items {
			result.Titles = append(result.Titles,
				record.VolumeInfo.Title)
			result.Identifiers = append(result.Identifiers,
				record.ID)

			result.Dates = append(result.Dates,
				record.VolumeInfo.PublishedDate)

			result.Authors = append(result.Authors,
				tools.LoopConCat(record.VolumeInfo.Authors))

			result.Imprints = append(result.Imprints,
				record.VolumeInfo.Publisher)
		}
		return result, nil
	}

	return result, nil
}
