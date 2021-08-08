package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {
	if err := fasthttp.ListenAndServe(fmt.Sprintf(":%s", settings["server"]["port"]), rootHandlers); err != nil {
		panic(err)
	}
}
