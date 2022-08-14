package ujson

import "time"

type Europeana struct {
	Apikey        string `json:"apikey"`
	Success       bool   `json:"success"`
	RequestNumber int    `json:"requestNumber"`
	ItemsCount    int    `json:"itemsCount"`
	TotalResults  int    `json:"totalResults"`
	Items         []struct {
		Completeness       int      `json:"completeness"`
		Country            []string `json:"country"`
		DataProvider       []string `json:"dataProvider"`
		DcCreator          []string `json:"dcCreator"`
		DcCreatorLangAware struct {
			Def []string `json:"def"`
		} `json:"dcCreatorLangAware"`
		DcLanguage          []string `json:"dcLanguage"`
		DcLanguageLangAware struct {
			Def []string `json:"def"`
		} `json:"dcLanguageLangAware"`
		DcTitleLangAware struct {
			Def []string `json:"def"`
		} `json:"dcTitleLangAware"`
		EdmConcept      []string `json:"edmConcept,omitempty"`
		EdmConceptLabel []struct {
			Def string `json:"def"`
		} `json:"edmConceptLabel,omitempty"`
		EdmConceptPrefLabelLangAware struct {
			De []string `json:"de"`
			Fi []string `json:"fi"`
			Ru []string `json:"ru"`
			Sv []string `json:"sv"`
			Pt []string `json:"pt"`
			El []string `json:"el"`
			Lt []string `json:"lt"`
			En []string `json:"en"`
			Hr []string `json:"hr"`
			It []string `json:"it"`
			Fr []string `json:"fr"`
			Hu []string `json:"hu"`
			Es []string `json:"es"`
			Et []string `json:"et"`
			Cs []string `json:"cs"`
			Sk []string `json:"sk"`
			Sl []string `json:"sl"`
			Pl []string `json:"pl"`
			Da []string `json:"da"`
			Ro []string `json:"ro"`
			Ca []string `json:"ca"`
			Nl []string `json:"nl"`
		} `json:"edmConceptPrefLabelLangAware,omitempty"`
		EdmDatasetName []string `json:"edmDatasetName"`
		EdmIsShownAt   []string `json:"edmIsShownAt"`
		EdmPlaceLabel  []struct {
			Def string `json:"def"`
		} `json:"edmPlaceLabel,omitempty"`
		EdmPlaceLabelLangAware struct {
			Def []string `json:"def"`
		} `json:"edmPlaceLabelLangAware,omitempty"`
		EdmPreview       []string `json:"edmPreview"`
		EdmTimespanLabel []struct {
			Def string `json:"def"`
		} `json:"edmTimespanLabel,omitempty"`
		EdmTimespanLabelLangAware struct {
			En []string `json:"en"`
			Fr []string `json:"fr"`
		} `json:"edmTimespanLabelLangAware,omitempty"`
		EuropeanaCollectionName []string  `json:"europeanaCollectionName"`
		EuropeanaCompleteness   int       `json:"europeanaCompleteness"`
		GUID                    string    `json:"guid"`
		ID                      string    `json:"id"`
		Index                   int       `json:"index"`
		Language                []string  `json:"language"`
		Link                    string    `json:"link"`
		PreviewNoDistribute     bool      `json:"previewNoDistribute"`
		Provider                []string  `json:"provider"`
		Rights                  []string  `json:"rights"`
		Score                   float64   `json:"score"`
		Timestamp               int64     `json:"timestamp"`
		TimestampCreated        time.Time `json:"timestamp_created"`
		TimestampCreatedEpoch   int64     `json:"timestamp_created_epoch"`
		TimestampUpdate         time.Time `json:"timestamp_update"`
		TimestampUpdateEpoch    int64     `json:"timestamp_update_epoch"`
		Title                   []string  `json:"title"`
		Type                    string    `json:"type"`
		Ugc                     []bool    `json:"ugc"`
		Year                    []string  `json:"year"`
		DcDescription           []string  `json:"dcDescription,omitempty"`
		DcDescriptionLangAware  struct {
			Def []string `json:"def"`
		} `json:"dcDescriptionLangAware,omitempty"`
		EdmIsShownBy     []string `json:"edmIsShownBy,omitempty"`
		EdmPlaceAltLabel []struct {
			Def string `json:"def"`
		} `json:"edmPlaceAltLabel,omitempty"`
		EdmPlaceAltLabelLangAware struct {
			Def []string `json:"def"`
		} `json:"edmPlaceAltLabelLangAware,omitempty"`
	} `json:"items"`
}
