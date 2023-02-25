package uxml

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var anetXML = []byte(`
<zs:searchRetrieveResponse xmlns:zs="http://www.loc.gov/zing/srw/">
<zs:version>1.1</zs:version>
<zs:numberOfRecords>30</zs:numberOfRecords>
<zs:records>
<zs:record>
<zs:recordSchema>dc</zs:recordSchema>
<zs:recordPacking>xml</zs:recordPacking>
<zs:recordData>
<srw_dc:dc xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<title xmlns="http://purl.org/dc/elements/1.1/">Vere admiranda, seu De magnitudine Romanae ecclesiae libri duo</title>
<creator xmlns="http://purl.org/dc/elements/1.1/">Stapletonus, Thomasaut</creator>
<type xmlns="http://purl.org/dc/elements/1.1/">text</type>
<publisher xmlns="http://purl.org/dc/elements/1.1/">Antverpiae ex Officina Plantiniana, apud Joannem Moretum</publisher>
<language xmlns="http://purl.org/dc/elements/1.1/">lat</language>
<description xmlns="http://purl.org/dc/elements/1.1/">Cockx-Indestege, E. Belgica typographica 4411</description>
<subject xmlns="http://purl.org/dc/elements/1.1/">od</subject>
<subject xmlns="http://purl.org/dc/elements/1.1/">einfo</subject>
<subject xmlns="http://purl.org/dc/elements/1.1/">zebra</subject>
<subject xmlns="http://purl.org/dc/elements/1.1/">dp-ua</subject>
<identifier xmlns="http://purl.org/dc/elements/1.1/">URN:ISBN:</identifier>
<identifier xmlns="http://purl.org/dc/elements/1.1/">URN:ISBN:</identifier>
<identifier xmlns="http://purl.org/dc/elements/1.1/">URN:ISBN:</identifier>
</srw_dc:dc>
</zs:recordData>
<zs:recordPosition>1</zs:recordPosition>
</zs:record>
<zs:record>
<zs:recordSchema>dc</zs:recordSchema>
<zs:recordPacking>xml</zs:recordPacking>
<zs:recordData>
<srw_dc:dc xmlns:srw_dc="info:srw/schema/1/dc-schema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="info:srw/schema/1/dc-schema http://www.loc.gov/standards/sru/resources/dc-schema.xsd">
<title xmlns="http://purl.org/dc/elements/1.1/">Admiranda Galliarum compendio indicata</title>
<creator xmlns="http://purl.org/dc/elements/1.1/">Frey, Janus Caeciliusaut</creator>
<type xmlns="http://purl.org/dc/elements/1.1/">text</type>
<publisher xmlns="http://purl.org/dc/elements/1.1/">Parisiis Apud Joannem Gesselin - Petrum David</publisher>
<language xmlns="http://purl.org/dc/elements/1.1/">lat</language>
<subject xmlns="http://purl.org/dc/elements/1.1/">zebra</subject>
<identifier xmlns="http://purl.org/dc/elements/1.1/">URN:ISBN:</identifier>
<identifier xmlns="http://purl.org/dc/elements/1.1/">URN:ISBN:</identifier>
</srw_dc:dc>
</zs:recordData>
<zs:recordPosition>2</zs:recordPosition>
</zs:record>
</zs:records>
<zs:echoedSearchRetrieveRequest>
<zs:version>1.1</zs:version>
<zs:query>title=admiranda</zs:query>
<zs:startRecord>1</zs:startRecord>
<zs:maximumRecords>2</zs:maximumRecords>
<zs:recordPacking>xml</zs:recordPacking>
<zs:recordSchema>dc</zs:recordSchema>
</zs:echoedSearchRetrieveRequest>
</zs:searchRetrieveResponse>
`)

func TestAnetParse0(t *testing.T) {
	var response Anet
	err := xml.Unmarshal(anetXML, &response)
	if err != nil {
		t.Errorf(fmt.Sprintf("\nUnmarshalling error[%s]\n", err))
	}
}

func TestAnetParse1(t *testing.T) {
	var response Anet
	_ = xml.Unmarshal(anetXML, &response)
	result := response.NumberOfRecords
	expected := "30"
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
	}
}

func TestAnetParse2(t *testing.T) {
	var response Anet
	_ = xml.Unmarshal(anetXML, &response)
	result := response.Records.Record[1].RecordData.Dc.Title.Text
	expected := "Admiranda Galliarum compendio indicata"
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
	}
}
