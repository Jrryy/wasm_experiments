# Experiments with Go and WASM.

Just some small stuff to test the capabilities of Go and WebAssembly. Done with [go-app](https://github.com/maxence-charriere/go-app) (requires Go 1.18+).

To build:
```sh
GOOS=js GOARCH=wasm go build -o web/app.wasm  # To build the client
go build  # To build the server
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .  # Necessary for your browser to run wasm
./wasm_experiments
```
You will be able to see the result in `localhost:8000/`
