package main

import (
	"log"
	"net/http"

	"test/internal/server/handler"
	"test/internal/server/router"
)

func main() {

	hdl := handler.New()

	rtr := router.New(hdl)

	log.Println(http.ListenAndServe("127.0.0.1:8080", rtr))
}
