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

// Global submit function
// Perform API calls and send response to Go
window.submit = function () {
    document.getElementById("result").innerHTML = "";
    let author = document.getElementById("author").value;
    let title = document.getElementById("title").value;
    let year = document.getElementById("year").value;
    if (author + title + year == "") {
        return
    }
    getAPIs().forEach(server => {
        let query = "";
        if (server["type"] == "sru") {
            if (author != "") {
                query = `${server["indices"]["author"]}%20=%20${author}`
            };
            if (title != "") {
                if (query != "") {
                    query = `${query}%20AND%20`
                }
                query = `${query}${server["indices"]["title"]}%20=%20${title}`
            };
            if (year != "") {
                if (query != "") {
                    query = `${query}%20AND%20`
                }
                query = `${query}${server["indices"]["year"]}%20=%20${year}`
            }
        } else if ((server["type"] == "query" || (server["type"] == "q"))) {
            if (author != "") {
                query = author
                if (server["name"] == "googlebooks") {
                    query = `${server["indices"]["author"]}:${author}`
                };
                if (server["name"] == "trove") {
                    query = `${server["indices"]["author"]}%3A${author}`
                };
            };
            if (title != "") {
                if (query != "") {
                    query = `${query}+`
                }

                if (server["name"] == "googlebooks") {
                    query = `${query}${server["indices"]["title"]}:${title}`
                } else if (server["name"] == "trove") {
                    query = `${query}${server["indices"]["title"]}%3A${title}`
                } else {
                    query = `${query}${title}`
                };
            };
            if (year != "") {
                if (query != "") {
                    query = `${query}+`
                }
                if (server["name"] == "trove") {
                    query = `${query}${server["indices"]["year"]}%3A${year}`
                } else {
                    query = `${query}${year}`
                }
            }
        };
        callAPI(server, query).then(result => handleResponse(
            server["name"],
            server["explain"],
            server["format"],
            result[0],
            result[1]));
    })
}