const go = new Go();
let mod, inst;
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
    result => {
        mod = result.module;
        inst = result.instance;
        go.run(inst);
    }
);

// to do: useful for being able to use this as API?
// function download(filename, text) {
//     // https://ourcodeworld.com/articles/read/189/how-to-create-a-file-and-generate-a-download-with-javascript-in-the-browser-without-a-server
//     var element = document.createElement('a');
//     element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
//     element.setAttribute('download', filename);
//     element.style.display = 'none';
//     document.body.appendChild(element);
//     element.click();
//     document.body.removeChild(element);
// }

async function callSRU(url, query) {
    const CORSproxy = "https://corsproxy.io/?"
    url = `${url}&operation=searchRetrieve&query=${query}&startRecord=1&maximumRecords=10`
    console.log(url);
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
    let year = document.getElementById("year").value;
    getSRUservers().forEach(server => {
        let query = "";
        if (author != "") {
            query = `${server["indices"]["author"]}%20=%20${author}`
        };
        if (title != "") {
            query = `${query}%20AND%20${server["indices"]["title"]}%20=%20${title}`
        };
        if (year != "") {
            query = `${query}%20AND%20${server["indices"]["year"]}%20=%20${year}`
        };
        callSRU(server["url"], query).then(result => handleXML(server["name"], result));
    })
}