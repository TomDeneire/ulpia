GOOS=js GOARCH=wasm go build -o main.wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
go test -json tomdeneire.github.io/ulpia/uxml > ./test/uxml_test.json
go test -json tomdeneire.github.io/ulpia/ujson > ./test/ujson_test.json