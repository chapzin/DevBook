package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chapzin/DevBook/api/src/config"
	"github.com/chapzin/DevBook/api/src/router"
)

func main() {
	config.Carregar()
	fmt.Printf("Rodando API na porta %s!!!", config.Porta)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":"+config.Porta, r))
}
