package main

import (
	"log"
	"net/http"

	"test/internal/server/handler"
	"test/internal/server/router"
)

func main() {

	handler := handler.New()

	router := router.NewRouter(handler)

	log.Println(http.ListenAndServe("127.0.0.1:8080", router))
}
