package main

import (
	"net/http"

	"github.com/cjoudrey/gluahttp"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	Lua := lua.NewState()
	defer Lua.Close()

	Lua.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)

	if err := Lua.DoString(`
        local http = require("http")

        local res, err = http.request('GET', 'https://httpbin.org/anything', {
            headers={
                ['User-Agent']='gluahttp/wasm'
            }
        })

        print(res.body)
    `); err != nil {
		panic(err)
	}
}
