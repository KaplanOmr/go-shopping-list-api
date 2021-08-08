package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {
	s, err := getSettings()
	if err != nil {
		panic(err)
	}

	if err = fasthttp.ListenAndServe(fmt.Sprintf(":%s", s["server"]["port"]), rootHandlers); err != nil {
		panic(err)
	}
}
