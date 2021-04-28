package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = ""
)

// Inicia variaveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta = os.Getenv("API_PORT")

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORTA"))
	if err != nil {
		dbPort = 5432
	}

	StringConexaoBanco = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), dbPort, os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_DATABASE"))

}
