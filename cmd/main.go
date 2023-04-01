package main

import (
	"log"
	"net/http"

	"test/internal/server/handlers"
	"test/internal/server/routers"
)

func main() {

	hdl := handlers.New()

	rtr := routers.New(hdl)

	log.Println(http.ListenAndServe("127.0.0.1:8080", rtr))
}
