package main

import (
	"log"

	"github.com/crecentmoon/proglog/1-LetsGo/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
