// Return list of APIs with metainformation
function getAPIs() {
  return [
    {
      name: "hpb",
      url: "https://sru.k10plus.de/hpb?version=2.0&recordSchema=dc&maximumRecords=100",
      explain: "https://sru.k10plus.de/hpb?version=2.0&operation=explain",
      indices: {
        author: "dc.creator",
        title: "dc.title",
        year: "pica.jvu",
      },
      type: "sru",
      format: "xml",
    },
    {
      name: "unicat",
      url: "https://www.unicat.be/sru?version=1.1&recordSchema=dc",
      explain: "https://www.unicat.be/sru?version=1.1&operation=explain",
      indices: {
        author: "author",
        title: "title",
        year: "year",
      },
      type: "sru",
      format: "xml",
    },
    {
      name: "fu-berlin",
      url: "https://fu-berlin.alma.exlibrisgroup.com/view/sru/49KOBV_FUB?version=1.2&recordSchema=dc",
      explain:
        "https://fu-berlin.alma.exlibrisgroup.com/view/sru/49KOBV_FUB?version=1.2&operation=explain",
      indices: {
        author: "alma.creator",
        title: "alma.title",
        year: "alma.date",
      },
      type: "sru",
      format: "xml",
    },
    {
      name: "bnf",
      url: "https://catalogue.bnf.fr/api/SRU?version=1.2&recordSchema=dc",
      explain: "https://catalogue.bnf.fr/api/SRU?version=1.2&operation=explain",
      indices: {
        author: "bib.author",
        title: "bib.title",
        year: "bib.date",
      },
      type: "sru",
      format: "xml",
    },
    {
      name: "dnb",
      url: "https://services.dnb.de/sru/dnb?version=1.1",
      explain: "https://services.dnb.de/sru/dnb?version=1.1&operation=explain",
      indices: {
        author: "dc.creator",
        title: "dc.title",
        year: "dc.date",
      },
      type: "sru",
      format: "xml",
    },
    {
      name: "europeana",
      url: "https://api.europeana.eu/record/v2/search.json?wskey=apidemo",
      explain: "https://api.europeana.eu",
      type: "query",
      format: "json",
    },
    {
      name: "openlibrary",
      url: "https://openlibrary.org/search.json?",
      explain: "https://openlibrary.org",
      type: "q",
      format: "json",
    },
    {
      name: "googlebooks",
      url: "https://www.googleapis.com/books/v1/volumes?",
      explain: "https://books.google.com/",
      indices: {
        author: "inauthor",
        title: "intitle",
      },
      type: "q",
      format: "json",
    },
    {
      name: "trove",
      url: "https://api.trove.nla.gov.au/v2/result?key=r6r5ak9u0l0nkbtg&zone=book",
      explain: "https://trove.nla.gov.au/about/create-something/using-api",
      indices: {
        author: "creator",
        title: "title",
        year: "date",
      },
      type: "q",
      format: "xml",
    },
  ];
}
// "http://gso.gbv.de/sru/DB=2.1/?version=1.1"
// "http://gita.grainger.uiuc.edu/registry/sru/sru.asp?version=1.1",
// "https://jsru.kb.nl/sru/sru?version=1.1"
// "http://sru.bibsys.no/search/biblio?version=1.1",
// "http://bvpb.mcu.es/i18n/sru/sru.cmd?version=1.1",
// "http://lx2.loc.gov:210/lcdb?version=1.1"
// "https://gallica.bnf.fr/SRU?version=1.2&recordSchema=dc&operation=searchRetrieve&query=dc.title=%22admiranda%22"
// -> unexisting or blocking CORS proxy

// to do
// STCN
// <http://data.bibliotheken.nl/sparql?default-graph-uri=&query=select+*+where+%7B+%3Fs+schema%3AmainEntityOfPage%2Fschema%3AisPartOf+%3Chttp%3A%2F%2Fdata.bibliotheken.nl%2Fid%2Fdataset%2Fstcn%3E+.+%7D>
// WorldCat <https://www.oclc.org/developer/api/oclc-apis/worldcat-search-api.en.html>

async function callAPI(server, query) {
  const CORSproxy = "https://corsproxy.io/?";

  switch (server["type"]) {
    case "sru":
      url = `${server["url"]}&operation=searchRetrieve&query=${query}&startRecord=1`;
      break;

    case "query":
      url = `${server["url"]}&query=${query}`;
      break;

    case "q":
      url = `${server["url"]}&q=${query}`;
      break;

    default:
      url = "";
  }

  let rawquery = url;
  console.log(rawquery);
  url = CORSproxy + url;
  let response = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "text/plain",
    },
  });
  let text = await response.text();
  return [rawquery, text];
}
