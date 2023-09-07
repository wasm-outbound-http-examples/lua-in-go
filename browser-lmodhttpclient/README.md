# Use ofunc/lua's lmodhttpclient library to send HTTP(s) requests from inside WASM

## Instructions for this devcontainer

### Preparation

1. Open this repo in devcontainer, e.g. using Github Codespaces.
   Type or copy/paste following commands to devcontainer's terminal.

As of [erdian718/lua](https://github.com/erdian718/lua) (aka ofunc/lua) received no updates since Jul 2019,
it can not be compiled by recent versions of Golang compiler due to enforcement of module mode (say `GO111MODULE=on`).
So let's build that Lua from sources, applying necessary patches:

2. Clone erdian718/lua's repo to get sources:

```sh
git clone --depth=1 https://github.com/erdian718/lua.git 
```

3. `cd` into the folder of erdian718/lua's sources:

```sh
cd lua
```

4. Obtain erdian718/lmodhttpclient's sources by cloning its repo next to built-in Lua modules:

```sh
git clone --depth=1 https://github.com/erdian718/lmodhttpclient.git
````

5. Init Golang module:

```sh
go mod init github.com/ofunc/lua
```

6. Patch sources to be compatible with module mod by running a Sed command:

```sh
grep -rl '"ofunc/lua' * | xargs sed -i.bak 's|"ofunc/lua|"github.com/ofunc/lua|g'
```

7. Additionally patch lmodjs's sources by running a Sed command:

```sh
sed -i.bak 's@if x := v.Get(LuaKey); x != undefined && x != null {@if x := v.Get(LuaKey); !(x.IsUndefined() || x.IsNull()) {@' lmodjs/mod.go
```

### Building

1. Ensure that your shell terminal is still in `lua` folder after completing all preparation steps.

2. Install lmodhttpclient's dependency library `publicsuffix`:

```sh
go get golang.org/x/net/publicsuffix
```

3. Make a folder for WASM module, copy there the sources of this example, and `cd` into it:

```sh
mkdir main
cd main
cp ../../browser-lmodhttpclient/main.go ./
cp ../../browser-lmodhttpclient/index.html ./
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

2. As `index.html` and a 10M-sized `main.wasm` are loaded into browser, refer to browser developer console
   to see the results.

### Finish

Perform your own experiments if desired.
