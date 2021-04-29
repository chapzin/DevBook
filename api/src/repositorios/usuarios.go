package repositorios

import (
	"database/sql"

	"github.com/chapzin/DevBook/api/src/models"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (u *Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, err := u.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) values (?,?,?,?)")
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	ultimoIdInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(ultimoIdInserido), nil
}
