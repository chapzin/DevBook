package models

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoem,omitempty"`
}

func (usuario *Usuario) Preparar() error {
	if err := usuario.validar(); err != nil {
		return err
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatorio e não pode esta em branco")
	}

	if usuario.Nick == "" {
		return errors.New("o nick é obrigatorio e não pode esta em branco")
	}

	if usuario.Email == "" {
		return errors.New("o email é obrigatorio e não pode esta em branco")
	}

	if usuario.Senha == "" {
		return errors.New("a senha é obrigatorio e não pode esta em branco")
	}
	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
