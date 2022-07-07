const go = new Go();
let mod, inst;
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
    result => {
        mod = result.module;
        inst = result.instance;
        go.run(inst);
    }
);

const CORSproxy = "https://corsproxy.io/?"

const SRUprefixes = [
    "https://sru.gbv.de/hpb?version=2.0",
    "https://www.unicat.be/sru?version=1.1",
    "https://fu-berlin.alma.exlibrisgroup.com/view/sru/49KOBV_FUB?version=1.2",
    "http://catalogue.bnf.fr/api/SRU?version=1.2",
    "http://services.dnb.de/sru/dnb?version=1.1",
    // "http://gso.gbv.de/sru/DB=2.1/?version=1.1",
    // "http://gita.grainger.uiuc.edu/registry/sru/sru.asp?version=1.1",
    "http://jsru.kb.nl/sru?version=1.1",
    // "http://sru.bibsys.no/search/biblio?version=1.1",
    // "http://bvpb.mcu.es/i18n/sru/sru.cmd?version=1.1",
    // "http://lx2.loc.gov:210/LCDB?version=1.1"
    // fu-berlin werkt enkel met indices!
    // zie ook https://www.loc.gov/standards/sru/resources/listOfServers.html
    // https://www.loc.gov/standards/sru/resources/lcServers.html
    // uitgecommentarieerde geven CORS issues!
]

async function callSRU(prefix, query) {
    let url = prefix + "&operation=searchRetrieve&query=" + query
    url = url + "&startRecord=1&maximumRecords=10&recordSchema=dc"
    url = CORSproxy + url
    let response = await fetch(url, {
        method: "GET",
        headers: {
            'Content-Type': 'text/plain'
        }
    });
    let text = await response.text();
    return text
}

// Global submit function
// Perform API calls and send response to Go
window.submit = function () {
    document.getElementById("result").innerHTML = "";
    let author = document.getElementById("author").value;
    let title = document.getElementById("title").value;
    SRUprefixes.forEach(prefix => {
        let query = "";
        if (prefix.includes("fu-berlin")) {
            query = `creator%20=%20${author}%20AND%20title%20=%20${title}`
        } else {
            query = author + "%20AND%20" + title
        };
        callSRU(prefix, query).then(result => handleXML(prefix, result))
    });
}