GOOS=js GOARCH=wasm go build -o main.wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
GOOS=js GOARCH=wasm go test -v tomdeneire.github.io/ulpia/uxml
GOOS=js GOARCH=wasm go test -v tomdeneire.github.io/ulpia/ujson