// Return list of APIs with metainformation
function getAPIs() {
    return [
        {
            "name": "hpb",
            "url": "https://sru.gbv.de/hpb?version=2.0&recordSchema=dc",
            "explain": "https://sru.gbv.de/hpb?version=2.0&operation=explain",
            "indices":
            {
                "author": "dc.creator",
                "title": "dc.title",
                "year": "pica.jvu"
            },
            "type": "sru",
            "format": "xml"
        },
        {
            "name": "unicat",
            "url": "https://www.unicat.be/sru?version=1.1&recordSchema=dc",
            "explain": "https://www.unicat.be/sru?version=1.1&operation=explain",
            "indices":
            {
                "author": "author",
                "title": "title",
                "year": "year"
            },
            "type": "sru",
            "format": "xml"
        },
        {
            "name": "fu-berlin",
            "url": "https://fu-berlin.alma.exlibrisgroup.com/view/sru/49KOBV_FUB?version=1.2&recordSchema=dc",
            "explain": "https://fu-berlin.alma.exlibrisgroup.com/view/sru/49KOBV_FUB?version=1.2&operation=explain",
            "indices":
            {
                "author": "alma.creator",
                "title": "alma.title",
                "year": "alma.date"
            },
            "type": "sru",
            "format": "xml"
        },
        {
            "name": "bnf",
            "url": "https://catalogue.bnf.fr/api/SRU?version=1.2&recordSchema=dc",
            "explain": "https://catalogue.bnf.fr/api/SRU?version=1.2&operation=explain",
            "indices":
            {
                "author": "bib.author",
                "title": "bib.title",
                "year": "bib.date"
            },
            "type": "sru",
            "format": "xml"
        },
        {
            "name": "dnb",
            "url": "https://services.dnb.de/sru/dnb?version=1.1",
            "explain": "https://services.dnb.de/sru/dnb?version=1.1&operation=explain",
            "indices":
            {
                "author": "dc.creator",
                "title": "dc.title",
                "year": "dc.date"
            },
            "type": "sru",
            "format": "xml"
        },
        {
            "name": "europeana",
            "url": "https://api.europeana.eu/record/v2/search.json?wskey=apidemo",
            "explain": "https://api.europeana.eu",
            "type": "query",
            "format": "json"
        },
        {
            "name": "openlibrary",
            "url": "https://openlibrary.org/search.json?",
            "explain": "https://openlibrary.org",
            "type": "q",
            "format": "json"
        }
    ]
}

// "http://gso.gbv.de/sru/DB=2.1/?version=1.1"
// "http://gita.grainger.uiuc.edu/registry/sru/sru.asp?version=1.1",
// "https://jsru.kb.nl/sru/sru?version=1.1"
// "http://sru.bibsys.no/search/biblio?version=1.1",
// "http://bvpb.mcu.es/i18n/sru/sru.cmd?version=1.1",
// "http://lx2.loc.gov:210/lcdb?version=1.1"
// -> allemaal onbestaande of blokkeren CORS

async function callAPI(server, query) {
    const CORSproxy = "https://corsproxy.io/?"

    switch (server["type"]) {
        case "sru": url = `${server["url"]}&operation=searchRetrieve&query=${query}&startRecord=1&maximumRecords=10`
            break;

        case "query": url = `${server["url"]}&query=${query}`
            break;

        case "q": url = `${server["url"]}&q=${query}`
            break;


        default: url = ""
    }

    rawquery = url;
    console.log(rawquery);
    url = CORSproxy + url
    let response = await fetch(url, {
        method: "GET",
        headers: {
            'Content-Type': 'text/plain'
        }
    });
    let text = await response.text();
    return [rawquery, text]
}