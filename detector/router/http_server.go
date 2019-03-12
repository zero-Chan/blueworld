package router

import (
	"log"
	"net/http"
	"zero-Chan/blueworld/detector/handler/geocoding"
)

func HttpServerInit() (err error) {
	server := http.NewServeMux()
	server.Handle("/geocoding/reverse", geocoding.GetModule())
	log.Printf("http Listen: %s", "127.0.0.1:8080")
	err = http.ListenAndServe("127.0.0.1:8080", server)
	return
}
