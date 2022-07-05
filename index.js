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
window.submit = function () {
    let v1 = document.getElementById("val1").value;
    let v2 = document.getElementById("val2").value;
    handle(v1, v2);
}