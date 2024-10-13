package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tomek-skrond/stripe-tests/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("loading env file failed")
	}

	lp := ":9999"
	server := api.NewAPIServer(lp)

	server.Run()
}
