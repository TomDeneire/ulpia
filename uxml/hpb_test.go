package uxml

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var hpbXML = []byte(`
<zs:searchRetrieveResponse xmlns:zs="http://docs.oasis-open.org/ns/search-ws/sruResponse">
<zs:numberOfRecords>25</zs:numberOfRecords>
<zs:records>
<zs:record>
<zs:recordSchema>dc</zs:recordSchema>
<zs:recordXMLEscaping>xml</zs:recordXMLEscaping>
<zs:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">IVSTI LIPSI ADMIRANDA, siue, DE MAGNITVDINE ROMANA LIBRI QVATTVOR: Ad Serenissimvm Principem Albertvm Avstrivm</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Admiranda, sive de magnitudine Romana</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Iusti Lipsi Admiranda, sive, de magnitudine Romana libri quattuor</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Admiranda, sive de magnitudine Romana</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Iusti Lipsi Admiranda, sive, de magnitudine Romana libri quattuor</dc:title>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Lipsius, Justus , 1547-1606</dc:contributor>
<dc:type xmlns:dc="http://purl.org/dc/elements/1.1/">Text</dc:type>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:language xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">lat</dc:language>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">print</dc:format>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">[6] Bl., 255 S., [6] Bl. Ill. 4° [i.e. 2°]</dc:format>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Vorlageform des Erscheinungsvermerks: ANTVERPIÆ, EX OFFICINA PLANTINIANA, Apud Ioannem Moretum. M. D. XCVIII.</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Hochschul- und Landesbibliothek Fulda 4o Gesch K 12/50</dc:description>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">(DE-599)HBZHT005025369</dc:identifier>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">(DE-605)HT005025369</dc:identifier>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">(DE-603)1048555313</dc:identifier>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/">ppn: (DE-601)151620598</dc:identifier>
</oai_dc:dc>
</zs:recordData>
<zs:recordPosition>1</zs:recordPosition>
</zs:record>
<zs:record>
<zs:recordSchema>dc</zs:recordSchema>
<zs:recordXMLEscaping>xml</zs:recordXMLEscaping>
<zs:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Iusti Lipsi Admiranda, siue, de magnitudine romana libri quattuor</dc:title>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Lipsius, Justus</dc:contributor>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Nivelle, Robert 1587-1598</dc:contributor>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">FR Parigi</dc:contributor>
<dc:type xmlns:dc="http://purl.org/dc/elements/1.1/">Text</dc:type>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:publisher xmlns:dc="http://purl.org/dc/elements/1.1/">apud Robertum Niuelle, sub signo Colomnarum, via Iacobea</dc:publisher>
<dc:language xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">lat</dc:language>
<dc:language xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">grc</dc:language>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">print</dc:format>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">[16], 374, [14] p. 8</dc:format>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Bianca l'ultima c.</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Cors. ; rom</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Iniziali, finalini e fregi xilogr</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Marca n.c. sul front</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Segn.: [ast]8 a-2A8 2B2</dc:description>
<dc:relation xmlns:dc="http://purl.org/dc/elements/1.1/">http://id.sbn.it/bid/VEAE139572</dc:relation>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/">ppn: (DE-601)144103036</dc:identifier>
</oai_dc:dc>
</zs:recordData>
<zs:recordPosition>2</zs:recordPosition>
</zs:record>
<zs:record>
<zs:recordSchema>dc</zs:recordSchema>
<zs:recordXMLEscaping>xml</zs:recordXMLEscaping>
<zs:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Iusti Lipsi admiranda, sive, de magnitudine romana. Libri quattuor</dc:title>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Lipsius, Justus (*aut)</dc:contributor>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Orry, Marc (*pbl)</dc:contributor>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">FR Parigi</dc:contributor>
<dc:type xmlns:dc="http://purl.org/dc/elements/1.1/">Text</dc:type>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:publisher xmlns:dc="http://purl.org/dc/elements/1.1/">Apud Marcum Orry, via Iacobaea, sub insigni Leonis Salientis</dc:publisher>
<dc:language xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">lat</dc:language>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">print</dc:format>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">14, 374, 12, 8</dc:format>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Segn.: *4 A-Z4 Aa4 Bb1</dc:description>
<dc:relation xmlns:dc="http://purl.org/dc/elements/1.1/">http://id.sbn.it/bid/RMSE005843</dc:relation>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/">ppn: (DE-601)14077291X</dc:identifier>
</oai_dc:dc>
</zs:recordData>
<zs:recordPosition>3</zs:recordPosition>
</zs:record>
<zs:record>
<zs:recordSchema>dc</zs:recordSchema>
<zs:recordXMLEscaping>xml</zs:recordXMLEscaping>
<zs:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Iusti Lipsi Admiranda, siue, De magnitudine Romana libri quattuor</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">De magnitudine Romana libri quattuor</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">De magnitudine Romana libri quattuor</dc:title>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Lipsius, Justus (*aut)</dc:contributor>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Moretus, Iohannes 1 (*pbl)</dc:contributor>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Officina Plantiniana Anversa (*pbl)</dc:contributor>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">BE Anversa</dc:contributor>
<dc:type xmlns:dc="http://purl.org/dc/elements/1.1/">Text</dc:type>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:publisher xmlns:dc="http://purl.org/dc/elements/1.1/">ex officina Plantiniana, apud Ioannem Moretum</dc:publisher>
<dc:language xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">lat</dc:language>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">print</dc:format>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">\12! , 255 , \13! p. 4o</dc:format>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Cors. ; gr. ; rom</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Iniziali xil.</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Marca calcogr. sul front., marca alla c.i4v</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Segn.: \ast!6A-Z4a-i4k6</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Manufactured: Antuerpiae : ex officina Plantiniana apud Ioannem Moretum, 1598</dc:description>
<dc:relation xmlns:dc="http://purl.org/dc/elements/1.1/">http://id.sbn.it/bid/BVEE005021</dc:relation>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/">ppn: (DE-601)134967666</dc:identifier>
</oai_dc:dc>
</zs:recordData>
<zs:recordPosition>4</zs:recordPosition>
</zs:record>
<zs:record>
<zs:recordSchema>dc</zs:recordSchema>
<zs:recordXMLEscaping>xml</zs:recordXMLEscaping>
<zs:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">IVSTI LIPSI[I] ADMIRANDA, siue, DE MAGNITVDINE ROMANA LIBRI QVATTVOR: AD SERENISSIMVM PRINCIPEM ALBERTVM AVSTRIVM</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Iusti Lipsi admiranda, sive, de magnitudine Romana libri quattuor</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Admiranda, sive, de magnitudine Romana libri quattuor</dc:title>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Lipsius, Justus , 1547-1606 (*aut)</dc:contributor>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Moretus, Joannes , 1543-1610 (*prt)</dc:contributor>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Plantijnsche Drukkerij (*prt)</dc:contributor>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Belgie Antverpy</dc:contributor>
<dc:type xmlns:dc="http://purl.org/dc/elements/1.1/">Text</dc:type>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:publisher xmlns:dc="http://purl.org/dc/elements/1.1/">Ex officina Plantiniana, Apud Ioannem Moretum</dc:publisher>
<dc:language xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">lat</dc:language>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">print</dc:format>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">[12], 255, [12] s. 4º (25 cm)</dc:format>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">CZ-PrNK.STT.stt20140120323</dc:identifier>
<dc:relation xmlns:dc="http://purl.org/dc/elements/1.1/">http://books.google.cz/books?vid=NKP:1003118376</dc:relation>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/">ppn: (DE-601)123758106</dc:identifier>
</oai_dc:dc>
</zs:recordData>
<zs:recordPosition>5</zs:recordPosition>
</zs:record>
<zs:record>
<zs:recordSchema>dc</zs:recordSchema>
<zs:recordXMLEscaping>xml</zs:recordXMLEscaping>
<zs:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Ivsti LipsI Admiranda: siue, De magnitvdine romana libri qvattvor</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Admiranda</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Admiranda</dc:title>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Lipsius, Justus , 1547-1606</dc:contributor>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Wing, John M. , 1845-1917 (former ownerpJohn Mansir)</dc:contributor>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Plantijnsche Drukkerij (printer)</dc:contributor>
<dc:type xmlns:dc="http://purl.org/dc/elements/1.1/">Text</dc:type>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">[1598]</dc:date>
<dc:publisher xmlns:dc="http://purl.org/dc/elements/1.1/">ex officina Plantiniana, apud Ioannem Moretum</dc:publisher>
<dc:language xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">fre</dc:language>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">print</dc:format>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">[12], 255, [12] p. 26 cm.</dc:format>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Printer's mark on t.p.; another form on verso of p. 235.</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">IBVA</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Provenance: Wing, John M., 1845-1917 / former owner [ICN., John Mansir]</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Bookplate of J.M. Wing. [ICN]</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Wing ZP 5465.P707</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">The Newberry, Chicago Wing ZP 5465.P707</dc:description>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">(NBYdb)505850</dc:identifier>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/">ppn: (DE-601)122129105</dc:identifier>
</oai_dc:dc>
</zs:recordData>
<zs:recordPosition>6</zs:recordPosition>
</zs:record>
<zs:record>
<zs:recordSchema>dc</zs:recordSchema>
<zs:recordXMLEscaping>xml</zs:recordXMLEscaping>
<zs:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Jvsti Lipsi Admiranda, siue, De magnitvdine romana libri quattuor</dc:title>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Lipsius, Justus , 1547-1606</dc:contributor>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Drouart, Ambroise , 1548-1608 (printer)</dc:contributor>
<dc:type xmlns:dc="http://purl.org/dc/elements/1.1/">Text</dc:type>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:publisher xmlns:dc="http://purl.org/dc/elements/1.1/">apud Ambrosivm Drovard</dc:publisher>
<dc:language xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">lat</dc:language>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">print</dc:format>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">3 p. l., 374 p., 6 l. 17 cm.</dc:format>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">IBVA</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Case F 0236.507</dc:description>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">The Newberry, Chicago Case F 0236.507</dc:description>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">(NBYdb)505849</dc:identifier>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/">ppn: (DE-601)122129091</dc:identifier>
</oai_dc:dc>
</zs:recordData>
<zs:recordPosition>7</zs:recordPosition>
</zs:record>
<zs:record>
<zs:recordSchema>dc</zs:recordSchema>
<zs:recordXMLEscaping>xml</zs:recordXMLEscaping>
<zs:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Ivsti LipsI Admiranda, siue, De Magnitvdine Romana Libri Qvattvor: Ad Serenissimvm Principem Albertvm Avstrivm</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Admiranda, sive, De Magnitudine Romana</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Iusti</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">9Admiranda, sive, De Magnitudine Romana</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Iusti</dc:title>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Lipsius, Justus , 1547-1606 (Verfasser, *aut)</dc:contributor>
<dc:type xmlns:dc="http://purl.org/dc/elements/1.1/">Text</dc:type>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:publisher xmlns:dc="http://purl.org/dc/elements/1.1/">Moretus</dc:publisher>
<dc:language xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">und</dc:language>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">print</dc:format>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">[6] Bl., 255 S., [6] Bl.</dc:format>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Erstausg. in 1350 Exemplaren. - Vorlageform d. Ersch.vermerks: Antverpiae, Ex Officina Plantiniana, Apud Ioannem Moretum. M.D.XCVIII. - Bibliogr. Nachweis: Bibliographie Lipsienne I.3</dc:description>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">(DE-599)BVBBV027395275</dc:identifier>
<dc:relation xmlns:dc="http://purl.org/dc/elements/1.1/">http://bvbd2.bib-bvb.de/opt-cgi/order/bestand.pl?ID=BV027395275</dc:relation>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/">ppn: (DE-601)103567348</dc:identifier>
</oai_dc:dc>
</zs:recordData>
<zs:recordPosition>8</zs:recordPosition>
</zs:record>
<zs:record>
<zs:recordSchema>dc</zs:recordSchema>
<zs:recordXMLEscaping>xml</zs:recordXMLEscaping>
<zs:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Iusti Lipsi Admiranda, sive, De magnitudine Romana: libri quattuor</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Admiranda, sive, De magnitudine Romana</dc:title>
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">9Admiranda, sive, De magnitudine Romana</dc:title>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Lipsius, Justus , 1547-1606 (Verfasser, *aut)</dc:contributor>
<dc:type xmlns:dc="http://purl.org/dc/elements/1.1/">Text</dc:type>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:publisher xmlns:dc="http://purl.org/dc/elements/1.1/">Sonnius</dc:publisher>
<dc:language xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">lat</dc:language>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">print</dc:format>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">8 Bl., 374 S., 6 Bl.</dc:format>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Nicht bei Schudt</dc:description>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">(DE-599)BVBBV037760803</dc:identifier>
<dc:relation xmlns:dc="http://purl.org/dc/elements/1.1/">http://bvbd2.bib-bvb.de/opt-cgi/order/bestand.pl?ID=BV037760803</dc:relation>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/">ppn: (DE-601)102956332</dc:identifier>
</oai_dc:dc>
</zs:recordData>
<zs:recordPosition>9</zs:recordPosition>
</zs:record>
<zs:record>
<zs:recordSchema>dc</zs:recordSchema>
<zs:recordXMLEscaping>xml</zs:recordXMLEscaping>
<zs:recordData>
<oai_dc:dc xmlns:oai_dc="http://www.openarchives.org/OAI/2.0/oai_dc/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai_dc/ http://www.openarchives.org/OAI/2.0/oai_dc.xsd">
<dc:title xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Admiranda sive de magnitudine Romana: libri quatuor</dc:title>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Lipsius, Justus , 1547-1606 (Verfasser, *aut)</dc:contributor>
<dc:contributor xmlns:dc="http://purl.org/dc/elements/1.1/">Antwerpen</dc:contributor>
<dc:type xmlns:dc="http://purl.org/dc/elements/1.1/">Text</dc:type>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:date xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">1598</dc:date>
<dc:publisher xmlns:dc="http://purl.org/dc/elements/1.1/">Plantin</dc:publisher>
<dc:language xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">lat</dc:language>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">print</dc:format>
<dc:format xmlns:dc="http://purl.org/dc/elements/1.1/">[6] Bl., 255 S., [6] Bl. 4 ̊</dc:format>
<dc:description xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">Justus Lipsius*</dc:description>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:srw_dc="info:srw/schema/1/dc-schema">(DE-599)BVBBV012400530</dc:identifier>
<dc:relation xmlns:dc="http://purl.org/dc/elements/1.1/">http://bvbd2.bib-bvb.de/opt-cgi/order/bestand.pl?ID=BV012400530</dc:relation>
<dc:identifier xmlns:dc="http://purl.org/dc/elements/1.1/">ppn: (DE-601)098118676</dc:identifier>
</oai_dc:dc>
</zs:recordData>
<zs:recordPosition>10</zs:recordPosition>
</zs:record>
</zs:records>
<zs:echoedSearchRetrieveRequest>
<zs:version>2.0</zs:version>
<zs:query>dc.creator = lipsius AND dc.title = admiranda AND pica.jvu = 1598</zs:query>
<zs:startRecord>1</zs:startRecord>
<zs:maximumRecords>10</zs:maximumRecords>
<zs:recordXMLEscaping>xml</zs:recordXMLEscaping>
<zs:recordSchema>dc</zs:recordSchema>
</zs:echoedSearchRetrieveRequest>
<zs:resultCountPrecision>exact</zs:resultCountPrecision>
</zs:searchRetrieveResponse>
`)

func TestHPBParse0(t *testing.T) {
	var response HPB
	err := xml.Unmarshal(hpbXML, &response)
	if err != nil {
		t.Errorf(fmt.Sprintf("\nUnmarshalling error[%s]\n", err))
	}
}

func TestHPBParse1(t *testing.T) {
	var response HPB
	_ = xml.Unmarshal(hpbXML, &response)
	result := response.NumberOfRecords
	expected := "25"
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
	}
}

// func TestHPBParse2(t *testing.T) {
// 	var response HPB
// 	_ = xml.Unmarshal(hpbXML, &response)
// 	result := response.Records.Record[1].RecordData.Dc.Title[0].Text
// 	expected := "Ivsti LipsI Admiranda siue, de magnitvdine Romana libri qvattvor"
// 	if result != expected {
// 		t.Errorf(fmt.Sprintf("\nResult: \n[%s]\nExpected: \n[%s]\n", result, expected))
// 	}
// }
