package uxml

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var dnbXML = []byte(`
<searchRetrieveResponse xmlns="http://www.loc.gov/zing/srw/">
<version>1.1</version>
<numberOfRecords>2</numberOfRecords>
<records>
<record>
<recordSchema>RDFxml</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<rdf:RDF xmlns:schema="http://schema.org/" xmlns:gndo="https://d-nb.info/standards/elementset/gnd#" xmlns:lib="http://purl.org/library/" xmlns:owl="http://www.w3.org/2002/07/owl#" xmlns:xsd="http://www.w3.org/2001/XMLSchema#" xmlns:skos="http://www.w3.org/2004/02/skos/core#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:editeur="https://ns.editeur.org/thema/" xmlns:geo="http://www.opengis.net/ont/geosparql#" xmlns:umbel="http://umbel.org/umbel#" xmlns:rdau="http://rdaregistry.info/Elements/u/" xmlns:sf="http://www.opengis.net/ont/sf#" xmlns:bflc="http://id.loc.gov/ontologies/bflc/" xmlns:dcterms="http://purl.org/dc/terms/" xmlns:isbd="http://iflastandards.info/ns/isbd/elements/" xmlns:mesh="http://id.nlm.nih.gov/mesh/vocab#" xmlns:foaf="http://xmlns.com/foaf/0.1/" xmlns:mo="http://purl.org/ontology/mo/" xmlns:marcRole="http://id.loc.gov/vocabulary/relators/" xmlns:agrelon="https://d-nb.info/standards/elementset/agrelon#" xmlns:dcmitype="http://purl.org/dc/dcmitype/" xmlns:dbp="http://dbpedia.org/property/" xmlns:dnbt="https://d-nb.info/standards/elementset/dnb#" xmlns:madsrdf="http://www.loc.gov/mads/rdf/v1#" xmlns:dnb_intern="http://dnb.de/" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:v="http://www.w3.org/2006/vcard/ns#" xmlns:wdrs="http://www.w3.org/2007/05/powder-s#" xmlns:ebu="http://www.ebu.ch/metadata/ontologies/ebucore/ebucore#" xmlns:bibo="http://purl.org/ontology/bibo/" xmlns:gbv="http://purl.org/ontology/gbv/" xmlns:dc="http://purl.org/dc/elements/1.1/">
<rdf:Description rdf:about="https://d-nb.info/98577178X">
<rdf:type rdf:resource="http://purl.org/ontology/bibo/Document"/>
<dcterms:medium rdf:resource="http://rdaregistry.info/termList/RDACarrierType/1044"/>
<rdau:P60049 rdf:resource="http://rdaregistry.info/termList/RDAContentType/1020"/>
<rdau:P60050 rdf:resource="http://rdaregistry.info/termList/RDAMediaType/1007"/>
<rdau:P60048 rdf:resource="http://rdaregistry.info/termList/RDACarrierType/1049"/>
<dc:identifier rdf:datatype="http://www.w3.org/2001/XMLSchema#string">(DE-101)98577178X</dc:identifier>
<bibo:isbn13 rdf:datatype="http://www.w3.org/2001/XMLSchema#string">9783487134161</bibo:isbn13>
<rdau:P60521 rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Gewebe : EUR 118.00</rdau:P60521>
<bibo:gtin14 rdf:datatype="http://www.w3.org/2001/XMLSchema#string">9783487134161</bibo:gtin14>
<dc:identifier rdf:datatype="http://www.w3.org/2001/XMLSchema#string">(OCoLC)254177603</dc:identifier>
<dcterms:language rdf:resource="http://id.loc.gov/vocabulary/iso639-2/ger"/>
<rdau:P60049 rdf:resource="https://d-nb.info/gnd/4135952-5"/>
<rdau:P60327 rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Justus Lipsius</rdau:P60327>
<dc:title rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Admiranda oder Wundergeschichten von der vnaußsprächlichen Macht, Herrlich- vnd Großmächtigkeit der Statt Rom</dc:title>
<dcterms:creator rdf:resource="https://d-nb.info/gnd/11857342X"/>
<marcRole:aut rdf:resource="https://d-nb.info/gnd/11857342X"/>
<dc:publisher rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Olms-Weidmann</dc:publisher>
<rdau:P60163 rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Hildesheim</rdau:P60163>
<rdau:P60163 rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Zürich</rdau:P60163>
<rdau:P60163 rdf:datatype="http://www.w3.org/2001/XMLSchema#string">New York, NY</rdau:P60163>
<rdau:P60333 rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Hildesheim ; Zürich ; New York, NY : Olms-Weidmann</rdau:P60333>
<rdau:P60539 rdf:datatype="http://www.w3.org/2001/XMLSchema#string">21 cm</rdau:P60539>
<dcterms:isPartOf rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Historia scientiarum : Fachgebiet Geschichte und Politik</dcterms:isPartOf>
<rdau:P60470 xml:lang="de">In Fraktur</rdau:P60470>
<dcterms:subject rdf:resource="https://d-nb.info/gnd/4050471-2"/>
<dcterms:subject rdf:resource="https://d-nb.info/gnd/4076769-3"/>
<dcterms:subject rdf:resource="https://d-nb.info/gnd/4076778-4"/>
<dc:subject rdf:datatype="https://d-nb.info/standards/elementset/dnb#ddc-subject-category">930</dc:subject>
<dcterms:subject rdf:resource="http://dewey.info/class/937/e22/"/>
<wdrs:describedby>
<rdf:Description rdf:about="https://d-nb.info/98577178X/about">
<dcterms:license rdf:resource="http://creativecommons.org/publicdomain/zero/1.0/"/>
<dcterms:modified rdf:datatype="http://www.w3.org/2001/XMLSchema#dateTime">2017-12-02T04:43:00.000</dcterms:modified>
</rdf:Description>
</wdrs:describedby>
<dcterms:issued rdf:datatype="http://www.w3.org/2001/XMLSchema#string">2007</dcterms:issued>
<isbd:P1053 rdf:datatype="http://www.w3.org/2001/XMLSchema#string">XCVII, 365, [17] S.</isbd:P1053>
<owl:sameAs rdf:resource="http://hub.culturegraph.org/resource/DNB-98577178X"/>
</rdf:Description>
</rdf:RDF>
</recordData>
<recordPosition>1</recordPosition>
</record>
<record>
<recordSchema>RDFxml</recordSchema>
<recordPacking>xml</recordPacking>
<recordData>
<rdf:RDF xmlns:schema="http://schema.org/" xmlns:gndo="https://d-nb.info/standards/elementset/gnd#" xmlns:lib="http://purl.org/library/" xmlns:owl="http://www.w3.org/2002/07/owl#" xmlns:xsd="http://www.w3.org/2001/XMLSchema#" xmlns:skos="http://www.w3.org/2004/02/skos/core#" xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#" xmlns:editeur="https://ns.editeur.org/thema/" xmlns:geo="http://www.opengis.net/ont/geosparql#" xmlns:umbel="http://umbel.org/umbel#" xmlns:rdau="http://rdaregistry.info/Elements/u/" xmlns:sf="http://www.opengis.net/ont/sf#" xmlns:bflc="http://id.loc.gov/ontologies/bflc/" xmlns:dcterms="http://purl.org/dc/terms/" xmlns:isbd="http://iflastandards.info/ns/isbd/elements/" xmlns:mesh="http://id.nlm.nih.gov/mesh/vocab#" xmlns:foaf="http://xmlns.com/foaf/0.1/" xmlns:mo="http://purl.org/ontology/mo/" xmlns:marcRole="http://id.loc.gov/vocabulary/relators/" xmlns:agrelon="https://d-nb.info/standards/elementset/agrelon#" xmlns:dcmitype="http://purl.org/dc/dcmitype/" xmlns:dbp="http://dbpedia.org/property/" xmlns:dnbt="https://d-nb.info/standards/elementset/dnb#" xmlns:madsrdf="http://www.loc.gov/mads/rdf/v1#" xmlns:dnb_intern="http://dnb.de/" xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#" xmlns:v="http://www.w3.org/2006/vcard/ns#" xmlns:wdrs="http://www.w3.org/2007/05/powder-s#" xmlns:ebu="http://www.ebu.ch/metadata/ontologies/ebucore/ebucore#" xmlns:bibo="http://purl.org/ontology/bibo/" xmlns:gbv="http://purl.org/ontology/gbv/" xmlns:dc="http://purl.org/dc/elements/1.1/">
<rdf:Description rdf:about="https://d-nb.info/998710113">
<rdf:type rdf:resource="http://purl.org/ontology/bibo/Document"/>
<dcterms:medium rdf:resource="http://rdaregistry.info/termList/RDACarrierType/1044"/>
<rdau:P60049 rdf:resource="http://rdaregistry.info/termList/RDAContentType/1020"/>
<rdau:P60050 rdf:resource="http://rdaregistry.info/termList/RDAMediaType/1007"/>
<rdau:P60048 rdf:resource="http://rdaregistry.info/termList/RDACarrierType/1049"/>
<dc:identifier rdf:datatype="http://www.w3.org/2001/XMLSchema#string">(DE-101)998710113</dc:identifier>
<dc:identifier rdf:datatype="http://www.w3.org/2001/XMLSchema#string">(OCoLC)723978183</dc:identifier>
<dcterms:language rdf:resource="http://id.loc.gov/vocabulary/iso639-2/lat"/>
<dc:title rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Ivsti LipsI Admiranda siue, de magnitvdine Romana libri qvattvor</dc:title>
<dcterms:alternative rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Admiranda sive de magnitudine Romana libri 4</dcterms:alternative>
<dcterms:creator rdf:resource="https://d-nb.info/gnd/11857342X"/>
<marcRole:aut rdf:resource="https://d-nb.info/gnd/11857342X"/>
<marcRole:prt rdf:resource="https://d-nb.info/gnd/119333910"/>
<marcRole:pbl rdf:resource="https://d-nb.info/gnd/10105548-1"/>
<marcRole:prt rdf:resource="https://d-nb.info/gnd/10105548-1"/>
<dc:publisher rdf:datatype="http://www.w3.org/2001/XMLSchema#string">ex Officina Plantiniana, apud Ioannem Moretum</dc:publisher>
<rdau:P60163 rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Antverpiae</rdau:P60163>
<rdau:P60333 rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Antverpiae : ex Officina Plantiniana, apud Ioannem Moretum</rdau:P60333>
<rdau:P60539 rdf:datatype="http://www.w3.org/2001/XMLSchema#string">4°</rdau:P60539>
<wdrs:describedby>
<rdf:Description rdf:about="https://d-nb.info/998710113/about">
<dcterms:license rdf:resource="http://creativecommons.org/publicdomain/zero/1.0/"/>
<dcterms:modified rdf:datatype="http://www.w3.org/2001/XMLSchema#dateTime">2021-12-13T16:35:17.000</dcterms:modified>
</rdf:Description>
</wdrs:describedby>
<dcterms:issued rdf:datatype="http://www.w3.org/2001/XMLSchema#string">1598</dcterms:issued>
<isbd:P1053 rdf:datatype="http://www.w3.org/2001/XMLSchema#string">Tit. m. gest. Sign., [5] Bl., 255 S., [6] Bl.</isbd:P1053>
<owl:sameAs rdf:resource="http://hub.culturegraph.org/resource/DNB-998710113"/>
</rdf:Description>
</rdf:RDF>
</recordData>
<recordPosition>2</recordPosition>
</record>
</records>
<echoedSearchRetrieveRequest>
<version>1.1</version>
<query>dc.creator = lipsius AND dc.title = admiranda</query>
<xQuery xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:nil="true"/>
<startRecord>1</startRecord>
<maximumRecords>10</maximumRecords>
</echoedSearchRetrieveRequest>
</searchRetrieveResponse>
`)

func TestDNBParse0(t *testing.T) {
	var response DNB
	err := xml.Unmarshal(dnbXML, &response)
	if err != nil {
		t.Errorf(fmt.Sprintf("\nUnmarshalling error[%s]\n", err))
	}
}

func TestDNBParse1(t *testing.T) {
	var response DNB
	_ = xml.Unmarshal(dnbXML, &response)
	result := response.Version
	expected := "1.1"
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
	}
}

func TestDNBParse2(t *testing.T) {
	var response DNB
	_ = xml.Unmarshal(dnbXML, &response)
	result := response.Records.Record[1].RecordData.RDF.Description.Title.Text
	expected := "Ivsti LipsI Admiranda siue, de magnitvdine Romana libri qvattvor"
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
	}
}
