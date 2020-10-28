package main

import (
	"log"

	"github.com/selmison/seed-desafio-cdc/pkg/api/rest"
)

func main() {
	if err := rest.InitHttpServer("127.0.0.1:3333"); err != nil {
		log.Fatalln(err)
	}
}
