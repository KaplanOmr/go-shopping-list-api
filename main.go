package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {
	settings, err := getSettings()
	if err != nil {
		panic(err)
	}

	if err = fasthttp.ListenAndServe(fmt.Sprintf(":%s", settings["server"]["port"]), rootHandlers); err != nil {
		panic(err)
	}
}
