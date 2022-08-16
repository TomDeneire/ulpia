package uxml

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var troveXML = []byte(`
<response>
<query>creator:lipsius title:admiranda</query>
<zone name="book">
<records s="*" n="4" total="4">
<work id="23463054" url="/work/23463054">
<troveUrl>https://trove.nla.gov.au/work/23463054</troveUrl>
<title>Ivsti Lipsi Admiranda siue, de magnitvdine romana libri qvattvor. : Tertia editio correctior, auctióreque</title>
<contributor>Lipsius, Justus, 1547-1606</contributor>
<issued>1605</issued>
<type>Book</type>
<holdingsCount>1</holdingsCount>
<versionCount>1</versionCount>
<hasCorrections>N</hasCorrections>
<relevance score="90.279205">very relevant</relevance>
<snippet>Ivsti Lipsi <b>Admiranda</b> siue, de magnitvdine romana libri qvattvor. : Tertia editio correctior</snippet>
<snippet><b>Admiranda</b> sive</snippet>
</work>
<work id="173400596" url="/work/173400596">
<troveUrl>https://trove.nla.gov.au/work/173400596</troveUrl>
<title>Ivsti LipsI Admiranda, siue, De magnitvdine romana libri qvattvor</title>
<contributor>Lipsius, Justus, 1547-1606</contributor>
<issued>1630</issued>
<type>Book</type>
<holdingsCount>1</holdingsCount>
<versionCount>1</versionCount>
<hasCorrections>N</hasCorrections>
<relevance score="90.12953">very relevant</relevance>
<snippet><b>Admiranda</b>.</snippet>
<snippet>Iusti LipsI <b>Admiranda</b>.</snippet>
</work>
<work id="23463101" url="/work/23463101">
<troveUrl>https://trove.nla.gov.au/work/23463101</troveUrl>
<title>Ivsti Lipsi Diva sichemiensis siue aspricollis : noua eius Beneficia et Admiranda</title>
<contributor>Lipsius, Justus, 1547-1606</contributor>
<issued>1606</issued>
<type>Book</type>
<holdingsCount>1</holdingsCount>
<versionCount>1</versionCount>
<hasCorrections>N</hasCorrections>
<relevance score="78.30725">very relevant</relevance>
<snippet>Ivsti Lipsi Diva sichemiensis siue aspricollis : noua eius Beneficia et <b>Admiranda</b>.</snippet>
</work>
<work id="245452508" url="/work/245452508">
<troveUrl>https://trove.nla.gov.au/work/245452508</troveUrl>
<title>Iusti LipsI Diua Sichemiensis siue Aspricollis : noua eius beneficia et admiranda</title>
<contributor>Lipsius, Justus, 1547-1606</contributor>
<issued>1620</issued>
<type>Book</type>
<holdingsCount>0</holdingsCount>
<versionCount>1</versionCount>
<hasCorrections>N</hasCorrections>
<relevance score="73.48471">very relevant</relevance>
<snippet>Iusti LipsI Diua Sichemiensis siue Aspricollis : noua eius beneficia et <b>admiranda</b>.</snippet>
</work>
</records>
</zone>
</response>
`)

func TestTroveParse0(t *testing.T) {
	var response Trove
	err := xml.Unmarshal(troveXML, &response)
	if err != nil {
		t.Errorf(fmt.Sprintf("\nUnmarshalling error[%s]\n", err))
	}
}

func TestTroveParse1(t *testing.T) {
	var response Trove
	_ = xml.Unmarshal(troveXML, &response)
	result := response.Zone.Records.Total
	expected := "4"
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
	}
}

// func TestTroveParse2(t *testing.T) {
// 	var response Trove
// 	_ = xml.Unmarshal(troveXML, &response)
// 	result := response.Records.Record[1].RecordData.Dc.Title[0].Text
// 	expected := "Ivsti LipsI Admiranda siue, de magnitvdine Romana libri qvattvor"
// 	if result != expected {
// 		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
// 	}
// }
