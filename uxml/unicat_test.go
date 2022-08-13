package uxml

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var unicatXML = []byte(`
<searchRetrieveResponse xmlns="http://www.loc.gov/zing/srw/">
<version>1.1</version>
<numberOfRecords>18</numberOfRecords>
<records>
<record>
<recordSchema>info:srw/schema/1/dc-schema</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<srw_dc:dc xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://purl.org/dc/elements/1.1/" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<title>Admiranda, sive: De magnitudine Romana libri quattuor</title>
<identifier>SYSID:93446791</identifier>
<identifier>ua:c:lvd:1039163</identifier>
<creator>Lipsius, Justus</creator>
<subject>Roman history</subject>
<language>lat</language>
<publisher>Parisiis apud Robertum Nivelle</publisher>
<date>1598</date>
</srw_dc:dc>
</recordData>
<recordPosition>1</recordPosition>
</record>
<record>
<recordSchema>info:srw/schema/1/dc-schema</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<srw_dc:dc xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://purl.org/dc/elements/1.1/" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<title>Admiranda, sive: De magnitudine Romana libri quattuor</title>
<identifier>SYSID:95178556</identifier>
<identifier>ua:c:lvd:931369</identifier>
<creator>Lipsius, Justus</creator>
<subject>Neo-Latin literature</subject>
<language>lat</language>
<publisher>Antverpiae ex Officina Plantiniana, apud Joannem Moretum</publisher>
<date>1598</date>
</srw_dc:dc>
</recordData>
<recordPosition>2</recordPosition>
</record>
<record>
<recordSchema>info:srw/schema/1/dc-schema</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<srw_dc:dc xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://purl.org/dc/elements/1.1/" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<title>Admiranda, siue, De magnitvdine Romana libri qvattvor</title>
<identifier>SYSID:95238543</identifier>
<identifier>stcv:c:stcv:12914078</identifier>
<creator>Lipsius, Justus</creator>
<language>lat</language>
<publisher>Antverpiæ ex officina Plantiniana, apud Ioannem Moretum</publisher>
<date>1598</date>
</srw_dc:dc>
</recordData>
<recordPosition>3</recordPosition>
</record>
<record>
<recordSchema>info:srw/schema/1/dc-schema</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<srw_dc:dc xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://purl.org/dc/elements/1.1/" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<title>Ivsti LipsI Admiranda, siue, De magnitvdine Romana libri qvattvor.</title>
<identifier>SYSID:91426095</identifier>
<identifier>ugent:rug01.002393933</identifier>
<creator>Lipsius, Justus</creator>
<creator>Moretus, Joannes</creator>
<language>lat</language>
<publisher>Antverpiæ [Antwerpen] : ex officina Plantiniana, apud Ioannem Moretum,</publisher>
<date>1598</date>
</srw_dc:dc>
</recordData>
<recordPosition>4</recordPosition>
</record>
<record>
<recordSchema>info:srw/schema/1/dc-schema</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<srw_dc:dc xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://purl.org/dc/elements/1.1/" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<title>Ivsti LipsI Admiranda, siue, De magnitvdine Romana libri qvattvor. ...</title>
<identifier>SYSID:96684517</identifier>
<identifier>libis:32LIBIS_ALMA_DS71252484050001471</identifier>
<creator>Lipsius, Justus</creator>
<creator>Albrecht</creator>
<creator>Moretus, Joannes I</creator>
<subject>Academic collection</subject>
<language>lat</language>
<publisher>Antverpiæ ex officina Plantiniana, apud Ioannem Moretum</publisher>
<date>1598</date>
</srw_dc:dc>
</recordData>
<recordPosition>5</recordPosition>
</record>
<record>
<recordSchema>info:srw/schema/1/dc-schema</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<srw_dc:dc xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://purl.org/dc/elements/1.1/" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<title>Justi Lipsii Admiranda, sive, De magnitudine Romana Libri Quattor. Ad Serenissimum Principem Albertum Austrium</title>
<identifier>SYSID:98717255</identifier>
<identifier>kbr:15418723</identifier>
<creator>Lipsius, Justus</creator>
<creator>Albrecht van Oostenrijk</creator>
<creator>Nivelle, Robert</creator>
<language>lat</language>
<publisher>Parisiis Apud Robertum Nivelle</publisher>
<date>1598</date>
</srw_dc:dc>
</recordData>
<recordPosition>6</recordPosition>
</record>
<record>
<recordSchema>info:srw/schema/1/dc-schema</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<srw_dc:dc xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://purl.org/dc/elements/1.1/" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<title>Ivsti LipsI Admiranda, siue, De magnitvdine romana libri quattuor. Ad serenissimum principem Ambertum Austrium.</title>
<identifier>SYSID:83873949</identifier>
<identifier>libis:32LIBIS_ALMA_DS71155374540001471</identifier>
<creator>Lipsius, Justus</creator>
<creator>Albrecht</creator>
<creator>Albert</creator>
<creator>Albertus</creator>
<creator>Albrecht and Isabella</creator>
<creator>Lipse, Juste</creator>
<creator>Lips, Joose</creator>
<creator>Lips, Josse</creator>
<creator>Lipsio, Giusto</creator>
<creator>Lipsio, Justo</creator>
<creator>Lips, Joest</creator>
<creator>Lipsii, Justi</creator>
<creator>Lipsius, I.</creator>
<creator>Lipsius, Iustus</creator>
<creator>Lipsius, J.</creator>
<creator>Nivelle, Robert</creator>
<creator>Printer old books</creator>
<subject>Academic collection</subject>
<language>lat</language>
<publisher>Parisiis apud Robertum Niuelle ...</publisher>
<date>1598</date>
</srw_dc:dc>
</recordData>
<recordPosition>7</recordPosition>
</record>
<record>
<recordSchema>info:srw/schema/1/dc-schema</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<srw_dc:dc xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://purl.org/dc/elements/1.1/" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<title>Iusti Lipsii Admiranda, sive, De magnitudine Romana libri quattuor. Ad serenissimum principem Albertum Austrium</title>
<identifier>SYSID:99126386</identifier>
<identifier>kbr:16658077</identifier>
<creator>Lipsius, Justus</creator>
<creator>Moretus, Joannes</creator>
<creator>Wesken, Iohanes</creator>
<creator>Albrecht van Oostenrijk</creator>
<language>lat</language>
<publisher>Antverpiae ex officina Plantiniana, apud Ioannem Moretum</publisher>
<date>1598</date>
</srw_dc:dc>
</recordData>
<recordPosition>8</recordPosition>
</record>
<record>
<recordSchema>info:srw/schema/1/dc-schema</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<srw_dc:dc xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://purl.org/dc/elements/1.1/" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<title>Iusti Lipsii Admiranda, sive, De magnitudine Romana libri quattuor. Ad serenissimum principem Albertum Austrium</title>
<identifier>SYSID:99152027</identifier>
<identifier>kbr:16736844</identifier>
<creator>Lipsius, Justus</creator>
<creator>Moretus, Joannes</creator>
<creator>Gijsen, Marnix</creator>
<creator>Grégoire, Henri</creator>
<language>lat</language>
<publisher>Antverpiae ex officina Plantiniana, apud Ioannem Moretum</publisher>
<date>1598</date>
</srw_dc:dc>
</recordData>
<recordPosition>9</recordPosition>
</record>
</records>
<extraResponseData/>
<echoedSearchRetrieveRequest>
<version>1.1</version>
<query>author = lipsius AND title = admiranda AND year = 1598</query>
<startRecord>1</startRecord>
<maximumRecords>10</maximumRecords>
<recordSchema>dc</recordSchema>
<xQuery>
<triple xmlns="http://www.loc.gov/zing/cql/xcql/">
<boolean>
<value>and</value>
</boolean>
<leftOperand>
<triple>
<boolean>
<value>and</value>
</boolean>
<leftOperand>
<searchClause>
<index>author</index>
<relation>
<value>=</value>
</relation>
<term>lipsius</term>
</searchClause>
</leftOperand>
<rightOperand>
<searchClause>
<index>title</index>
<relation>
<value>=</value>
</relation>
<term>admiranda</term>
</searchClause>
</rightOperand>
</triple>
</leftOperand>
<rightOperand>
<searchClause>
<index>year</index>
<relation>
<value>=</value>
</relation>
<term>1598</term>
</searchClause>
</rightOperand>
</triple>
</xQuery>
</echoedSearchRetrieveRequest>
</searchRetrieveResponse>
`)

func TestUnicatParse0(t *testing.T) {
	var response Unicat
	err := xml.Unmarshal(unicatXML, &response)
	if err != nil {
		t.Errorf(fmt.Sprintf("\nUnmarshalling error[%s]\n", err))
	}
}

func TestUnicatParse1(t *testing.T) {
	var response Unicat
	_ = xml.Unmarshal(unicatXML, &response)
	result := response.NumberOfRecords
	expected := "18"
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
	}
}

// func TestUnicatParse2(t *testing.T) {
// 	var response Unicat
// 	_ = xml.Unmarshal(unicatXML, &response)
// 	result := response.Records.Record[1].RecordData.Dc.Title[0].Text
// 	expected := "Ivsti LipsI Admiranda siue, de magnitvdine Romana libri qvattvor"
// 	if result != expected {
// 		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
// 	}
// }
