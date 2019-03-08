package main

import (
	"zero-Chan/blueworld/detector/router"
	"fmt"
)

func main() {
	err := router.HttpServerInit()
	if err != nil {
		fmt.Printf("http server init fail: %s", err)
		return
	}
}
