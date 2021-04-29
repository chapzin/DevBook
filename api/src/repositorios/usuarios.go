package repositorios

import (
	"database/sql"
	"fmt"

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

func (u *Usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	linhas, err := u.db.Query("select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?", nomeOuNick, nomeOuNick)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)

	}

	return usuarios, nil

}

func (u *Usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linha, err := u.db.Query("select id, nome,nick, email, criadoEm from usuarios where id = ?", ID)
	if err != nil {
		return models.Usuario{}, err
	}

	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if err = linha.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return models.Usuario{}, err
		}
	}
	return usuario, nil
}
