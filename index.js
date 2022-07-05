const go = new Go();
let mod, inst;
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
    result => {
        mod = result.module;
        inst = result.instance;
        go.run(inst);
    }
);

// Global submit function
// Sends all input values to Go
window.submit = function () {
    handleInput(
        document.getElementById("author").value,
        document.getElementById("title").value);
}