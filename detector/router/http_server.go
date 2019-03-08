package router

import (
	"net/http"
	"zero-Chan/blueworld/detector/handler/geocoding"
)

func HttpServerInit() (err error) {
	server := http.NewServeMux()
	server.Handle("geocoding/reverse", geocoding.GetModule())
	err = http.ListenAndServe("127.0.0.1:8080", server)
	return
}
