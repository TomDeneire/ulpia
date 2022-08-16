package ujson

import (
	"encoding/json"
	"fmt"
	"testing"
)

var openlibraryJSON = []byte(`
{
    "numFound": 2,
    "start": 0,
    "numFoundExact": true,
    "docs": [
        {
            "key": "/works/OL658813W",
            "type": "work",
            "seed": [
                "/books/OL19277441M",
                "/books/OL18873772M",
                "/works/OL658813W",
                "/subjects/antiquities",
                "/subjects/place:rome",
                "/subjects/time:early_works_to_1800",
                "/authors/OL4700580A"
            ],
            "title": "Iusti Lipsi Admiranda, siue, De magnitudine Romana libri quattuor",
            "title_suggest": "Iusti Lipsi Admiranda, siue, De magnitudine Romana libri quattuor",
            "subtitle": "ad Serenissimum Principem Albertum Austrium.",
            "edition_count": 2,
            "edition_key": [
                "OL19277441M",
                "OL18873772M"
            ],
            "publish_date": [
                "1598"
            ],
            "publish_year": [
                1598
            ],
            "first_publish_year": 1598,
            "number_of_pages_median": 315,
            "publish_place": [
                "Parisiis",
                "Antuerpiae"
            ],
            "contributor": [
                "Nivelle, Robert, 1558-1598"
            ],
            "last_modified_i": 1373932718,
            "ebook_count_i": 0,
            "ebook_access": "no_ebook",
            "has_fulltext": false,
            "public_scan_b": false,
            "publisher": [
                "Apud Robertum Niuelle ...",
                "Ex Officina Plantiniana apud Ioannem Moretum"
            ],
            "language": [
                "lat"
            ],
            "author_key": [
                "OL4700580A"
            ],
            "author_name": [
                "Justus Lipsius"
            ],
            "author_alternative_name": [
                "La Mara"
            ],
            "place": [
                "Rome"
            ],
            "subject": [
                "Antiquities"
            ],
            "time": [
                "Early works to 1800"
            ],
            "publisher_facet": [
                "Apud Robertum Niuelle ...",
                "Ex Officina Plantiniana apud Ioannem Moretum"
            ],
            "place_key": [
                "rome"
            ],
            "time_facet": [
                "Early works to 1800"
            ],
            "subject_facet": [
                "Antiquities"
            ],
            "_version_": 1735674887355236352,
            "place_facet": [
                "Rome"
            ],
            "author_facet": [
                "OL4700580A Justus Lipsius"
            ],
            "subject_key": [
                "antiquities"
            ],
            "time_key": [
                "early_works_to_1800"
            ]
        },
        {
            "key": "/works/OL16828355M",
            "type": "work",
            "seed": [
                "/books/OL16828355M",
                "/works/OL16828355M",
                "/subjects/catholic_church_--_apologetic_works.",
                "/subjects/rome_(italy)_--_antiquities."
            ],
            "title": "Admiranda et vere admiranda, siue, De magnitudine et vrbis et Ecclesiae Romanae",
            "title_suggest": "Admiranda et vere admiranda, siue, De magnitudine et vrbis et Ecclesiae Romanae",
            "edition_count": 1,
            "edition_key": [
                "OL16828355M"
            ],
            "publish_date": [
                "1600"
            ],
            "publish_year": [
                1600
            ],
            "first_publish_year": 1600,
            "number_of_pages_median": 236,
            "publish_place": [
                "Romae"
            ],
            "contributor": [
                "Lipsius, Justus, 1547-1606.",
                "Stapleton, Thomas, 1535-1598.",
                "Schoppe, Kaspar, 1576-1649."
            ],
            "last_modified_i": 1655273261,
            "ebook_count_i": 0,
            "ebook_access": "no_ebook",
            "has_fulltext": false,
            "public_scan_b": false,
            "publisher": [
                "Ex bibliotheca Bartholomaei Grassi apud Nicolaum Mutium"
            ],
            "language": [
                "lat"
            ],
            "subject": [
                "Catholic Church -- Apologetic works.",
                "Rome (Italy) -- Antiquities."
            ],
            "publisher_facet": [
                "Ex bibliotheca Bartholomaei Grassi apud Nicolaum Mutium"
            ],
            "subject_facet": [
                "Catholic Church -- Apologetic works.",
                "Rome (Italy) -- Antiquities."
            ],
            "_version_": 1735679815643561984,
            "subject_key": [
                "catholic_church_--_apologetic_works.",
                "rome_(italy)_--_antiquities."
            ]
        }
    ],
    "num_found": 2,
    "q": "lipsius admiranda 1598",
    "offset": null
}
`)

func TestOpenLibraryParse0(t *testing.T) {
	var response OpenLibrary
	err := json.Unmarshal(openlibraryJSON, &response)
	if err != nil {
		t.Errorf(fmt.Sprintf("\nUnmarshalling error[%s]\n", err))
	}
}

func TestOpenLibraryParse1(t *testing.T) {
	var response OpenLibrary
	_ = json.Unmarshal(openlibraryJSON, &response)
	result := response.NumFound
	expected := 2
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n[%v]\nExpected: \n[%v]\n", result, expected))
	}
}
