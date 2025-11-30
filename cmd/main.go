package main

import (
	"log"
	"os"
)

func main() {
	config := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: config,
	}

	if err := api.run(api.mount()); err != nil {
		log.Printf("Server has failed to start, err: %s", err)
		os.Exit(1)
	}
}
