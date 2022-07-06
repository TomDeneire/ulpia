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

const SRUprefixes = ["http://sru.gbv.de/hpb?version=2.0", "https://www.unicat.be/sru?version=1.1"]

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
    let author = document.getElementById("author").value;
    let title = document.getElementById("title").value;
    let query = author + "%20AND%20" + title;
    SRUprefixes.forEach(prefix => {
        callSRU(prefix, query).then(result => handleXML(result))
    });
}