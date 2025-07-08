package model

import (
	"errors"
	"strings"
)

type Usuario struct {
	ID    int64  `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type UsuarioResponse struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
	Nome  string `json:"nome"`
}

type CredenciaisLogin struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func (u *Usuario) Validar() error {
	errs := make([]string, 0)

	// Validação do nome
	if strings.TrimSpace(u.Nome) == "" {
		errs = append(errs, "Nome é obrigatório")
	}

	// Validação do email
	if u.Email == "" {
		errs = append(errs, "E-mail é obrigatório")
	} else if !validarEmail(u.Email) {
		errs = append(errs, "E-mail inválido")
	}

	// Validação da sennha
	if u.Senha == "" {
		errs = append(errs, "Senha é obrigatória")
	}

	// Verifica se há erros
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}

	return nil
}

func validarEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}
