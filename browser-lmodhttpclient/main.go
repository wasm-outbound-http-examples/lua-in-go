package main

import (
	"fmt"
	"strings"

	"github.com/ofunc/lua/lmodhttpclient"
	"github.com/ofunc/lua/util"
)

func main() {
	l := util.NewState()
	l.Preload("http/client", lmodhttpclient.Open)

	err := l.LoadText(strings.NewReader(`
local io = require 'io'
local httpclient = require 'http/client'

local res = httpclient.get('https://httpbin.org/anything')
io.copy(res, io.stdout)
res:close()
	`), "", 0)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	errmsg := l.PCall(0, 0, true)
	if errmsg != nil {
		fmt.Printf("%v", errmsg)
	}
}
