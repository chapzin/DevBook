package router

import (
	"github.com/chapzin/DevBook/api/src/router/rotas"
	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
