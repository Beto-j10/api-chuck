package main

import (
	"fmt"
	"os"
	"prueba/server"

	"github.com/go-chi/chi/v5"
)

type Chuck struct {
	Id string `json:"id"`
}

type Chucks struct {
	Chucks []Chuck `json:"chucks"`
}

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "startup error: %v\n", err)
		os.Exit(1)
	}

}

func run() error {
	// New router
	router := chi.NewRouter()

	server := server.NewServer(router)
	err := server.Run()
	if err != nil {
		return err
	}
	return nil
}
