package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "server: ", log.LstdFlags)
	logger.Println("старт")
	srv := server.ReturnRouter(logger)
	err := srv.HTTP.ListenAndServe()
	if err != nil {
		srv.Logger.Fatal(err)
	}
}
