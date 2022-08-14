package ujson

type OpenLibrary struct {
	NumFound      int  `json:"numFound"`
	Start         int  `json:"start"`
	NumFoundExact bool `json:"numFoundExact"`
	Docs          []struct {
		Key                   string   `json:"key"`
		Type                  string   `json:"type"`
		Seed                  []string `json:"seed"`
		Title                 string   `json:"title"`
		TitleSuggest          string   `json:"title_suggest"`
		Subtitle              string   `json:"subtitle,omitempty"`
		EditionCount          int      `json:"edition_count"`
		EditionKey            []string `json:"edition_key"`
		PublishDate           []string `json:"publish_date"`
		PublishYear           []int    `json:"publish_year"`
		FirstPublishYear      int      `json:"first_publish_year"`
		NumberOfPagesMedian   int      `json:"number_of_pages_median"`
		PublishPlace          []string `json:"publish_place"`
		Contributor           []string `json:"contributor"`
		LastModifiedI         int      `json:"last_modified_i"`
		EbookCountI           int      `json:"ebook_count_i"`
		EbookAccess           string   `json:"ebook_access"`
		HasFulltext           bool     `json:"has_fulltext"`
		PublicScanB           bool     `json:"public_scan_b"`
		Publisher             []string `json:"publisher"`
		Language              []string `json:"language"`
		AuthorKey             []string `json:"author_key,omitempty"`
		AuthorName            []string `json:"author_name,omitempty"`
		AuthorAlternativeName []string `json:"author_alternative_name,omitempty"`
		Place                 []string `json:"place,omitempty"`
		Subject               []string `json:"subject"`
		Time                  []string `json:"time,omitempty"`
		PublisherFacet        []string `json:"publisher_facet"`
		PlaceKey              []string `json:"place_key,omitempty"`
		TimeFacet             []string `json:"time_facet,omitempty"`
		SubjectFacet          []string `json:"subject_facet"`
		Version               int64    `json:"_version_"`
		PlaceFacet            []string `json:"place_facet,omitempty"`
		AuthorFacet           []string `json:"author_facet,omitempty"`
		SubjectKey            []string `json:"subject_key"`
		TimeKey               []string `json:"time_key,omitempty"`
	} `json:"docs"`
	Q      string      `json:"q"`
	Offset interface{} `json:"offset"`
}
