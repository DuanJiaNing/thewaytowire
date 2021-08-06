package main

import (
	"context"
	"log"
)

func main() {
	server, cleanup, err := setup(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	log.Println("start server and listen: ", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
