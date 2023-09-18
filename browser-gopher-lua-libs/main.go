package main

import (
	httplib "github.com/vadv/gopher-lua-libs/http"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	Lua := lua.NewState()
	defer Lua.Close()

	httplib.Preload(Lua)

	err := Lua.DoString(`
local http = require('http')
local client = http.client()
local req = http.request('GET', 'https://httpbin.org/anything')
local res, err = client:do_request(req)
if err then error(err) end
print(res.body)
	`)
	if err != nil {
		panic(err)
	}
}
