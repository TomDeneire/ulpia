GOOS=js GOARCH=wasm go build -o main.wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
go test -v tomdeneire.github.io/ulpia/uxml
go test -v tomdeneire.github.io/ulpia/ujson