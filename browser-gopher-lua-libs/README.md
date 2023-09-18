# Use GopherLua's gopher-lua-libs library to send HTTP(s) requests from inside WASM

## Instructions for this devcontainer

### Preparation

1. Open this repo in devcontainer, e.g. using Github Codespaces.
   Type or copy/paste following commands to devcontainer's terminal.

### Building

1. `cd` into the folder of this example:

```sh
cd browser-gopher-lua-libs
```

2. Install GopherLua and the `gopher-lua-libs` library:

```sh
go get github.com/yuin/gopher-lua
go get github.com/vadv/gopher-lua-libs
```

3. Workaround compilation error (for WASM target) by updating `pq` library:

```sh
go get -u github.com/lib/pq@master
```

4. Compile the example:

```sh
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

5. Copy the glue JS from Golang distribution to example's folder:

```sh
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

### Test with browser

1. Run simple HTTP server to temporarily publish project to Web:

```sh
python3 -m http.server
```

Codespace will show you "Open in Browser" button. Just click that button or
obtain web address from "Forwarded Ports" tab.

2. As `index.html` and a **30M**-sized `main.wasm` are loaded into browser, refer to browser developer console
   to see the results.

### Finish

Perform your own experiments if desired.
