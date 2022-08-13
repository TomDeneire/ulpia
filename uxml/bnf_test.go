package uxml

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var bnfXML = []byte(`
<srw:searchRetrieveResponse xmlns:srw="http://www.loc.gov/zing/srw/" xmlns="http://catalogue.bnf.fr/namespaces/InterXMarc" xmlns:ixm="http://catalogue.bnf.fr/namespaces/InterXMarc" xmlns:mn="http://catalogue.bnf.fr/namespaces/motsnotices" xmlns:sd="http://www.loc.gov/zing/srw/diagnostic/">
<srw:version>1.2</srw:version>
<srw:echoedSearchRetrieveRequest>
<srw:version>1.2</srw:version>
<srw:query>bib.author = lipsius AND bib.title = politica</srw:query>
</srw:echoedSearchRetrieveRequest>
<srw:numberOfRecords>19</srw:numberOfRecords>
<srw:records>
<srw:record>
<srw:recordSchema>dc</srw:recordSchema>
<srw:recordPacking>xml</srw:recordPacking>
<srw:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:identifier>http://catalogue.bnf.fr/ark:/12148/cb30823499w</dc:identifier>
<dc:title>Monita et exempla politica. Libri duo, qui virtutes et vitia principum spectant.</dc:title>
<dc:creator>Lipse, Juste (1547-1606). Auteur du texte</dc:creator>
<dc:publisher>Amsterdami, apud Guilielmum Blaeuw. Anno M DC XXX.</dc:publisher>
<dc:date>1630</dc:date>
<dc:format>233 p. et index ; in-12</dc:format>
<dc:language>lat</dc:language>
<dc:language>latin</dc:language>
<dc:type xml:lang="fre">texte imprimé</dc:type>
<dc:type xml:lang="eng">printed text</dc:type>
<dc:type xml:lang="eng">text</dc:type>
<dc:rights xml:lang="fre">Catalogue en ligne de la Bibliothèque nationale de France</dc:rights>
<dc:rights xml:lang="eng">French National Library online Catalog</dc:rights>
</oai_dc:dc>
</srw:recordData>
<srw:recordIdentifier>ark:/12148/cb30823499w</srw:recordIdentifier>
<srw:recordPosition>1</srw:recordPosition>
<srw:extraRecordData>
<ixm:attr name="CreationDate">19970701</ixm:attr>
<ixm:attr name="LastModificationDate">20191220</ixm:attr>
<mn:score>11.627088</mn:score>
</srw:extraRecordData>
</srw:record>
<srw:record>
<srw:recordSchema>dc</srw:recordSchema>
<srw:recordPacking>xml</srw:recordPacking>
<srw:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:identifier>http://catalogue.bnf.fr/ark:/12148/cb33996956r</dc:identifier>
<dc:title>J. Lipsi Monita et exempla politica libri duo...</dc:title>
<dc:creator>Lipse, Juste (1547-1606). Auteur du texte</dc:creator>
<dc:date>1668</dc:date>
<dc:format>In-12</dc:format>
<dc:language>lat</dc:language>
<dc:language>latin</dc:language>
<dc:type xml:lang="fre">texte imprimé</dc:type>
<dc:type xml:lang="eng">printed text</dc:type>
<dc:type xml:lang="eng">text</dc:type>
<dc:rights xml:lang="fre">Catalogue en ligne de la Bibliothèque nationale de France</dc:rights>
<dc:rights xml:lang="eng">French National Library online Catalog</dc:rights>
</oai_dc:dc>
</srw:recordData>
<srw:recordIdentifier>ark:/12148/cb33996956r</srw:recordIdentifier>
<srw:recordPosition>2</srw:recordPosition>
<srw:extraRecordData>
<ixm:attr name="CreationDate">19970701</ixm:attr>
<ixm:attr name="LastModificationDate">19970701</ixm:attr>
<mn:score>11.481978</mn:score>
</srw:extraRecordData>
</srw:record>
<srw:record>
<srw:recordSchema>dc</srw:recordSchema>
<srw:recordPacking>xml</srw:recordPacking>
<srw:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:identifier>http://catalogue.bnf.fr/ark:/12148/cb43724046s</dc:identifier>
<dc:title>Monita et exempla politica . Libri duo, qui virtutes et vitia principum spectant.</dc:title>
<dc:creator>Lipse, Juste (1547-1606). Auteur du texte</dc:creator>
<dc:publisher>Parisiis. Excudebat Petrus Chevalier, in monte Divi Hilarii. M. DC. V.</dc:publisher>
<dc:date>1605</dc:date>
<dc:format>156 f. ; in-8</dc:format>
<dc:language>lat</dc:language>
<dc:language>latin</dc:language>
<dc:type xml:lang="fre">texte imprimé</dc:type>
<dc:type xml:lang="eng">printed text</dc:type>
<dc:type xml:lang="eng">text</dc:type>
<dc:rights xml:lang="fre">Catalogue en ligne de la Bibliothèque nationale de France</dc:rights>
<dc:rights xml:lang="eng">French National Library online Catalog</dc:rights>
</oai_dc:dc>
</srw:recordData>
<srw:recordIdentifier>ark:/12148/cb43724046s</srw:recordIdentifier>
<srw:recordPosition>3</srw:recordPosition>
<srw:extraRecordData>
<ixm:attr name="CreationDate">20131205</ixm:attr>
<ixm:attr name="LastModificationDate">20131205</ixm:attr>
<mn:score>11.381741</mn:score>
</srw:extraRecordData>
</srw:record>
<srw:record>
<srw:recordSchema>dc</srw:recordSchema>
<srw:recordPacking>xml</srw:recordPacking>
<srw:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:identifier>http://catalogue.bnf.fr/ark:/12148/cb30823496v</dc:identifier>
<dc:title>Justi Lipsii Monita et exempla politica libri duo, qui virtutes et vitia principum spectant</dc:title>
<dc:creator>Lipse, Juste (1547-1606). Auteur du texte</dc:creator>
<dc:publisher>ex officina plantiniana, apud J. Moretum (Antverpiae)</dc:publisher>
<dc:date>1606</dc:date>
<dc:format>In-8° , 282 p., index et privilège. (Van der Haeghen, II, 137.)</dc:format>
<dc:language>lat</dc:language>
<dc:language>latin</dc:language>
<dc:type xml:lang="fre">texte imprimé</dc:type>
<dc:type xml:lang="eng">printed text</dc:type>
<dc:type xml:lang="eng">text</dc:type>
<dc:rights xml:lang="fre">Catalogue en ligne de la Bibliothèque nationale de France</dc:rights>
<dc:rights xml:lang="eng">French National Library online Catalog</dc:rights>
</oai_dc:dc>
</srw:recordData>
<srw:recordIdentifier>ark:/12148/cb30823496v</srw:recordIdentifier>
<srw:recordPosition>4</srw:recordPosition>
<srw:extraRecordData>
<ixm:attr name="CreationDate">19970701</ixm:attr>
<ixm:attr name="LastModificationDate">19970701</ixm:attr>
<mn:score>10.878499</mn:score>
</srw:extraRecordData>
</srw:record>
<srw:record>
<srw:recordSchema>dc</srw:recordSchema>
<srw:recordPacking>xml</srw:recordPacking>
<srw:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:identifier>http://catalogue.bnf.fr/ark:/12148/cb323882051</dc:identifier>
<dc:title>Justi Lipsi Monita et exempla politica. Libri duo qui virtutes et vitia principum spectant</dc:title>
<dc:creator>Lipse, Juste (1547-1606). Auteur du texte</dc:creator>
<dc:date>1650</dc:date>
<dc:language>lat</dc:language>
<dc:language>latin</dc:language>
<dc:type xml:lang="fre">texte imprimé</dc:type>
<dc:type xml:lang="eng">printed text</dc:type>
<dc:type xml:lang="eng">text</dc:type>
<dc:rights xml:lang="fre">Catalogue en ligne de la Bibliothèque nationale de France</dc:rights>
<dc:rights xml:lang="eng">French National Library online Catalog</dc:rights>
</oai_dc:dc>
</srw:recordData>
<srw:recordIdentifier>ark:/12148/cb323882051</srw:recordIdentifier>
<srw:recordPosition>5</srw:recordPosition>
<srw:extraRecordData>
<ixm:attr name="CreationDate">19970701</ixm:attr>
<ixm:attr name="LastModificationDate">19970701</ixm:attr>
<mn:score>10.8782015</mn:score>
</srw:extraRecordData>
</srw:record>
<srw:record>
<srw:recordSchema>dc</srw:recordSchema>
<srw:recordPacking>xml</srw:recordPacking>
<srw:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:identifier>http://catalogue.bnf.fr/ark:/12148/cb308234976</dc:identifier>
<dc:title>Monita et exempla politica libri duo, qui virtutes et vitia principum spectant</dc:title>
<dc:creator>Lipse, Juste (1547-1606). Auteur du texte</dc:creator>
<dc:publisher>ex officina plantiniana, apud viduam et filios J. Moreti (Antverpiae)</dc:publisher>
<dc:date>1613</dc:date>
<dc:format>VIII-217 p. ; in-4</dc:format>
<dc:language>lat</dc:language>
<dc:language>latin</dc:language>
<dc:type xml:lang="fre">texte imprimé</dc:type>
<dc:type xml:lang="eng">printed text</dc:type>
<dc:type xml:lang="eng">text</dc:type>
<dc:rights xml:lang="fre">Catalogue en ligne de la Bibliothèque nationale de France</dc:rights>
<dc:rights xml:lang="eng">French National Library online Catalog</dc:rights>
</oai_dc:dc>
</srw:recordData>
<srw:recordIdentifier>ark:/12148/cb308234976</srw:recordIdentifier>
<srw:recordPosition>6</srw:recordPosition>
<srw:extraRecordData>
<ixm:attr name="CreationDate">19970701</ixm:attr>
<ixm:attr name="LastModificationDate">20191220</ixm:attr>
<mn:score>10.87768</mn:score>
</srw:extraRecordData>
</srw:record>
<srw:record>
<srw:recordSchema>dc</srw:recordSchema>
<srw:recordPacking>xml</srw:recordPacking>
<srw:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:identifier>http://catalogue.bnf.fr/ark:/12148/cb30823495h</dc:identifier>
<dc:title>Justi Lipsii Monita et exempla politica libri duo, qui virtutes et vitia principum spectant</dc:title>
<dc:creator>Lipse, Juste (1547-1606). Auteur du texte</dc:creator>
<dc:publisher>P. Chevalier (Parisiis)</dc:publisher>
<dc:date>1605</dc:date>
<dc:format>In-8° , 156 ff. (Van der Haeghen, II, 133.)</dc:format>
<dc:language>lat</dc:language>
<dc:language>latin</dc:language>
<dc:type xml:lang="fre">texte imprimé</dc:type>
<dc:type xml:lang="eng">printed text</dc:type>
<dc:type xml:lang="eng">text</dc:type>
<dc:rights xml:lang="fre">Catalogue en ligne de la Bibliothèque nationale de France</dc:rights>
<dc:rights xml:lang="eng">French National Library online Catalog</dc:rights>
</oai_dc:dc>
</srw:recordData>
<srw:recordIdentifier>ark:/12148/cb30823495h</srw:recordIdentifier>
<srw:recordPosition>7</srw:recordPosition>
<srw:extraRecordData>
<ixm:attr name="CreationDate">19970701</ixm:attr>
<ixm:attr name="LastModificationDate">20171128</ixm:attr>
<mn:score>10.877525</mn:score>
</srw:extraRecordData>
</srw:record>
<srw:record>
<srw:recordSchema>dc</srw:recordSchema>
<srw:recordPacking>xml</srw:recordPacking>
<srw:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:identifier>http://catalogue.bnf.fr/ark:/12148/cb30823498j</dc:identifier>
<dc:title>Justi Lipsii Monita et exempla politica libri duo, qui virtutes et vitia principum spectant</dc:title>
<dc:creator>Lipse, Juste (1547-1606). Auteur du texte</dc:creator>
<dc:publisher>P. Chevalier (Parisiis)</dc:publisher>
<dc:date>1618</dc:date>
<dc:format>In-8° , pièces limin., 290 p. (Van der Haeghen, II, 141.)</dc:format>
<dc:language>lat</dc:language>
<dc:language>latin</dc:language>
<dc:type xml:lang="fre">texte imprimé</dc:type>
<dc:type xml:lang="eng">printed text</dc:type>
<dc:type xml:lang="eng">text</dc:type>
<dc:rights xml:lang="fre">Catalogue en ligne de la Bibliothèque nationale de France</dc:rights>
<dc:rights xml:lang="eng">French National Library online Catalog</dc:rights>
</oai_dc:dc>
</srw:recordData>
<srw:recordIdentifier>ark:/12148/cb30823498j</srw:recordIdentifier>
<srw:recordPosition>8</srw:recordPosition>
<srw:extraRecordData>
<ixm:attr name="CreationDate">19970701</ixm:attr>
<ixm:attr name="LastModificationDate">20171203</ixm:attr>
<mn:score>10.877107</mn:score>
</srw:extraRecordData>
</srw:record>
<srw:record>
<srw:recordSchema>dc</srw:recordSchema>
<srw:recordPacking>xml</srw:recordPacking>
<srw:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:identifier>http://catalogue.bnf.fr/ark:/12148/cb33996953q</dc:identifier>
<dc:title>Les Politiques ou doctrine civile de Juste Lipse,... 3e édition</dc:title>
<dc:creator>Lipse, Juste (1547-1606). Auteur du texte</dc:creator>
<dc:publisher>C. de Monstroeil (Paris)</dc:publisher>
<dc:publisher>et J. Richer (Paris)</dc:publisher>
<dc:date>1597</dc:date>
<dc:format>In-12</dc:format>
<dc:language>frm</dc:language>
<dc:language>français moyen </dc:language>
<dc:type xml:lang="fre">texte imprimé</dc:type>
<dc:type xml:lang="eng">printed text</dc:type>
<dc:type xml:lang="eng">text</dc:type>
<dc:rights xml:lang="fre">Catalogue en ligne de la Bibliothèque nationale de France</dc:rights>
<dc:rights xml:lang="eng">French National Library online Catalog</dc:rights>
</oai_dc:dc>
</srw:recordData>
<srw:recordIdentifier>ark:/12148/cb33996953q</srw:recordIdentifier>
<srw:recordPosition>9</srw:recordPosition>
<srw:extraRecordData>
<ixm:attr name="CreationDate">19970701</ixm:attr>
<ixm:attr name="LastModificationDate">20211022</ixm:attr>
<mn:score>9.923368</mn:score>
</srw:extraRecordData>
</srw:record>
<srw:record>
<srw:recordSchema>dc</srw:recordSchema>
<srw:recordPacking>xml</srw:recordPacking>
<srw:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:identifier>http://catalogue.bnf.fr/ark:/12148/cb39312631r</dc:identifier>
<dc:title>Les Politiques ou Doctrine Civile de Juste Lipsius, où est principalement discouru de ce qui appartient à la principauté. Troisiesme Edition</dc:title>
<dc:creator>Lipse, Juste (1547-1606). Auteur du texte</dc:creator>
<dc:publisher>Claude de Monstr'oeil et Jean Richer (Paris)</dc:publisher>
<dc:date>1597</dc:date>
<dc:format>In-12, pièces limin., 259 p., table</dc:format>
<dc:language>frm</dc:language>
<dc:language>français moyen </dc:language>
<dc:type xml:lang="fre">texte imprimé</dc:type>
<dc:type xml:lang="eng">printed text</dc:type>
<dc:type xml:lang="eng">text</dc:type>
<dc:rights xml:lang="fre">Catalogue en ligne de la Bibliothèque nationale de France</dc:rights>
<dc:rights xml:lang="eng">French National Library online Catalog</dc:rights>
</oai_dc:dc>
</srw:recordData>
<srw:recordIdentifier>ark:/12148/cb39312631r</srw:recordIdentifier>
<srw:recordPosition>10</srw:recordPosition>
<srw:extraRecordData>
<ixm:attr name="CreationDate">20011211</ixm:attr>
<ixm:attr name="LastModificationDate">20211026</ixm:attr>
<mn:score>9.910714</mn:score>
</srw:extraRecordData>
</srw:record>
</srw:records>
<srw:nextRecordPosition>11</srw:nextRecordPosition>
</srw:searchRetrieveResponse>
`)

func TestBNFParse0(t *testing.T) {
	var response BNF
	err := xml.Unmarshal(bnfXML, &response)
	if err != nil {
		t.Errorf(fmt.Sprintf("\nUnmarshalling error[%s]\n", err))
	}
}

func TestBNFParse1(t *testing.T) {
	var response BNF
	_ = xml.Unmarshal(bnfXML, &response)
	result := response.NumberOfRecords
	expected := "19"
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
	}
}

// func TestBNFParse2(t *testing.T) {
// 	var response BNF
// 	_ = xml.Unmarshal(bnfXML, &response)
// 	result := response.Records.Record[1].RecordData.RDF.Description.Title.Text
// 	expected := "Ivsti LipsI Admiranda siue, de magnitvdine Romana libri qvattvor"
// 	if result != expected {
// 		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
// 	}
// }
