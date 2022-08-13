package uxml

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var fuberlinXML = []byte(`
<searchRetrieveResponse xmlns="http://www.loc.gov/zing/srw/">
<version>1.2</version>
<numberOfRecords>5</numberOfRecords>
<records>
<record>
<recordSchema>dc</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<srw_dc:dc xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<dc:title>Iusti Lipsi Monita et exempla politica. [electronic resource] /</dc:title>
<dc:contributor>Lipsius, Justus, 1547-1606.</dc:contributor>
<dc:type>text</dc:type>
<dc:publisher>Vesaliae Clivorum :</dc:publisher>
<dc:publisher>Ex officina Andreae ab Hoogenhux,</dc:publisher>
<dc:date>[M] [D] CLXXI. [1671]</dc:date>
<dc:date>1671</dc:date>
<dc:language>lat</dc:language>
<dc:description>eebo-0018</dc:description>
<dc:description>Early English books tract supplement interim guide1070.m.4</dc:description>
<dc:description>Electronic reproduction.Ann Arbor, Mich. :UMI,1999-(Early English books online)Digital version of: (Early English books; Tract supplement ; E2:1[97])s1999 miun s</dc:description>
<dc:subject>Political science</dc:subject>
<dc:subject>Political ethics</dc:subject>
<dc:subject>Printers' marks</dc:subject>
<dc:coverage>Fragment; consists of t.p. only.</dc:coverage>
<dc:coverage>Printer's device on t.p.</dc:coverage>
<dc:coverage>Imprint date contains inverted C's.</dc:coverage>
<dc:coverage>Reproduction of original in: British Library.</dc:coverage>
<dc:identifier>(CKB)3360000000358616</dc:identifier>
<dc:identifier>(EEBO)2240876413</dc:identifier>
<dc:identifier>(OCoLC)ocn226372811e</dc:identifier>
<dc:identifier>(OCoLC)226372811</dc:identifier>
<dc:identifier>alma:49KOBV_FUB/bibs/9958113190402883</dc:identifier>
</srw_dc:dc>
</recordData>
<recordIdentifier>9958113190402883</recordIdentifier>
<recordPosition>1</recordPosition>
</record>
<record>
<recordSchema>dc</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<srw_dc:dc xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<dc:title>Monita et exempla politica : libri 2, qui virtutes et vitia principum spectant /</dc:title>
<dc:contributor>Lipsius, Justus 1547-1606 (DE-588)11857342X</dc:contributor>
<dc:type>text</dc:type>
<dc:date>1671</dc:date>
<dc:language>ger</dc:language>
<dc:language>ger</dc:language>
<dc:identifier>BV026207386</dc:identifier>
<dc:identifier>(DE-605)HT005054901</dc:identifier>
<dc:identifier>(DE-599)BVBBV026207386</dc:identifier>
<dc:identifier>(DE-604)BV026207386</dc:identifier>
<dc:identifier>alma:49KOBV_FUB/bibs/990007544970402883</dc:identifier>
</srw_dc:dc>
</recordData>
<recordIdentifier>990007544970402883</recordIdentifier>
<recordPosition>2</recordPosition>
</record>
<record>
<recordSchema>dc</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<srw_dc:dc xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<dc:title>Politica : six books of politics or political instruction /</dc:title>
<dc:contributor>Lipsius, Justus 1547-1606 (DE-588)11857342X</dc:contributor>
<dc:contributor>Waszink, Jan 1969- (DE-588)135743311</dc:contributor>
<dc:type>text</dc:type>
<dc:date>2004</dc:date>
<dc:language>lat</dc:language>
<dc:language>lat</dc:language>
<dc:language>eng</dc:language>
<dc:description>Zugl.: Amsterdam, Univ., Diss. von Jan Waszink</dc:description>
<dc:subject>B785.L42</dc:subject>
<dc:subject>Politieke theorie</dc:subject>
<dc:subject>Politische Wissenschaft</dc:subject>
<dc:subject>Political science</dc:subject>
<dc:subject>Staatslehre</dc:subject>
<dc:subject>Justus Lipsius (1547 - 1606)</dc:subject>
<dc:subject>Neulateinische Literatur</dc:subject>
<dc:subject>Lipsius, Justus 1547-1606</dc:subject>
<dc:subject>Lipsius, Justus 1547-1606</dc:subject>
<dc:coverage>Text lat. und engl., Kommentar engl.</dc:coverage>
<dc:identifier>BV019541825</dc:identifier>
<dc:identifier>(gbd)0826535</dc:identifier>
<dc:identifier>(gbd)0840074</dc:identifier>
<dc:identifier>(DE-599)BVBBV019541825</dc:identifier>
<dc:identifier>(DE-604)BV019541825</dc:identifier>
<dc:identifier>90-232-4038-3</dc:identifier>
<dc:identifier>alma:49KOBV_FUB/bibs/990026996230402883</dc:identifier>
</srw_dc:dc>
</recordData>
<recordIdentifier>990026996230402883</recordIdentifier>
<recordPosition>3</recordPosition>
</record>
<record>
<recordSchema>dc</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<srw_dc:dc xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<dc:title>Politicorum sive civilis doctrinae libri sex : qui ad principatum maxime spectant ; additae Notae auctiores, tum et de una religione liber.</dc:title>
<dc:title>Politica sive civilis doctrina</dc:title>
<dc:contributor>Lipsius, Justus 1547-1606 (DE-588)11857342X</dc:contributor>
<dc:type>text</dc:type>
<dc:date>1596</dc:date>
<dc:language>lat</dc:language>
<dc:language>lat</dc:language>
<dc:identifier>BV026856863</dc:identifier>
<dc:identifier>(DE-576)031517250</dc:identifier>
<dc:identifier>(DE-599)BSZ031517250</dc:identifier>
<dc:identifier>(DE-604)BV026856863</dc:identifier>
<dc:identifier>alma:49KOBV_FUB/bibs/990039913660402883</dc:identifier>
</srw_dc:dc>
</recordData>
<recordIdentifier>990039913660402883</recordIdentifier>
<recordPosition>4</recordPosition>
</record>
<record>
<recordSchema>dc</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<srw_dc:dc xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<dc:title>Politicorum sive civilis doctrinae libri sex /</dc:title>
<dc:title>Politicorvm sive civilis doctrinae libri sex</dc:title>
<dc:title>Politica sive civilis doctrina</dc:title>
<dc:contributor>Lipsius, Justus 1547-1606 (DE-588)11857342X</dc:contributor>
<dc:contributor>Weber, Wolfgang 1950- (DE-588)122629000</dc:contributor>
<dc:type>text</dc:type>
<dc:date>1998</dc:date>
<dc:language>ger</dc:language>
<dc:language>ger</dc:language>
<dc:subject>Staatslehre</dc:subject>
<dc:subject>Lipsius, Justus 1547-1606</dc:subject>
<dc:identifier>BV012189820</dc:identifier>
<dc:identifier>(DE-599)BVBBV012189820</dc:identifier>
<dc:identifier>(DE-604)BV012189820</dc:identifier>
<dc:identifier>3-487-10664-7</dc:identifier>
<dc:identifier>alma:49KOBV_FUB/bibs/990028431510402883</dc:identifier>
</srw_dc:dc>
</recordData>
<recordIdentifier>990028431510402883</recordIdentifier>
<recordPosition>5</recordPosition>
</record>
</records>
<extraResponseData xmlns:xb="http://www.exlibris.com/repository/search/xmlbeans/">
<xb:exact>true</xb:exact>
<xb:responseDate>2022-08-13T16:24:11+0200</xb:responseDate>
</extraResponseData>
</searchRetrieveResponse>
`)

func TestFUBerlinParse0(t *testing.T) {
	var response FUBerlin
	err := xml.Unmarshal(fuberlinXML, &response)
	if err != nil {
		t.Errorf(fmt.Sprintf("\nUnmarshalling error[%s]\n", err))
	}
}

func TestFUBerlinParse1(t *testing.T) {
	var response FUBerlin
	_ = xml.Unmarshal(fuberlinXML, &response)
	result := response.NumberOfRecords
	expected := "5"
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
	}
}

// func TestFUBerlinParse2(t *testing.T) {
// 	var response FUBerlin
// 	_ = xml.Unmarshal(fuberlinXML, &response)
// 	result := response.Records.Record[1].RecordData.Dc.Title[0].Text
// 	expected := "Ivsti LipsI Admiranda siue, de magnitvdine Romana libri qvattvor"
// 	if result != expected {
// 		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
// 	}
// }
