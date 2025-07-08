package repository

import (
	"api/banco"
	"api/model"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UsuarioRepositorio struct {
	db *sqlx.DB
}

// NovoRepositorioDeUsuarios cria uma nova instância no repositório
func NovoRepositorioDeUsuarios(dq *sqlx.DB) *UsuarioRepositorio {
	return &UsuarioRepositorio{db: banco.DB}
}

// CriarUsuario insere um novo usuário no banco de dados
func (repo *UsuarioRepositorio) CriarUsuario(usuario model.Usuario) (uint, error) {

	// Verifica se o email já está cadastrado
	var emailExistente string
	err := repo.db.QueryRow(
		`SELECT email FROM usuarios WHERE email = $1`,
		usuario.Email,
	).Scan(&emailExistente)

	if err == nil {
		return 0, fmt.Errorf("email já cadastrado")
	} else if err != sql.ErrNoRows {
		return 0, fmt.Errorf("erro ao verificar email: %w", err)
	}

	// Validar campos
	if err := usuario.Validar(); err != nil {
		return 0, fmt.Errorf("erro de validação: %w", err)
	}

	// Gerar hash da senha
	hash, err := GerarHashSenha(usuario.Senha)
	if err != nil {
		return 0, fmt.Errorf("erro ao gerar hash da senha: %w", err)

	}
	usuario.Senha = hash

	// Cria um novo usuário
	var id uint
	err = repo.db.QueryRow(
		`INSERT INTO usuarios (nome, email, senha)
		VALUES($1, $2, $3) RETURNING id`,
		usuario.Nome,
		usuario.Email,
		usuario.Senha,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("erro ao inserir usuário: %w", err)
	}
	return id, nil
}

// BuscarUsuarioPorID busca um usuário pelo email
func (repo *UsuarioRepositorio) BuscarUsuarioPorEmail(email string) (*model.Usuario, error) {
	// Busca um usuário pelo email
	query := `SELECT id, nome, email, senha FROM usuarios WHERE email = $1`
	row := repo.db.QueryRow(query, email)

	var usuario model.Usuario
	err := row.Scan(&usuario.ID, &usuario.Nome, &usuario.Email, &usuario.Senha)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuário não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return &usuario, nil
}

func GerarHashSenha(senha string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return string(hash), err
}

func CompararSenha(hash, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
}
