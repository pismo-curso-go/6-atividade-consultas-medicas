package handlers

import (
	"net/http"
	"saudemais-api/auth"
	"saudemais-api/database"
	"saudemais-api/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Estruturas para receber os dados JSON das requisições
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterPatient lida com a requisição de cadastro de um novo paciente.
func RegisterPatient(c echo.Context) error {
	req := new(RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Dados inválidos", "code": http.StatusBadRequest})
	}

	// Validação básica para campos obrigatórios
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Nome, email e senha são obrigatórios", "code": http.StatusBadRequest})
	}

	// Criptografa a senha antes de salvar no banco
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Falha ao processar a senha", "code": http.StatusInternalServerError})
	}

	patient := models.Patient{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Tenta criar o paciente no banco de dados
	if result := database.DB.Create(&patient); result.Error != nil {
		// Retorna erro se o email já estiver cadastrado (violando a constraint 'unique')
		return c.JSON(http.StatusConflict, map[string]interface{}{"message": "Paciente já cadastrado", "code": http.StatusConflict})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Paciente cadastrado com sucesso!"})
}

// Login lida com a requisição de autenticação de um paciente.
func Login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Dados inválidos", "code": http.StatusBadRequest})
	}

	var patient models.Patient
	// Procura o paciente pelo email no banco de dados
	if result := database.DB.Where("email = ?", req.Email).First(&patient); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Email ou senha inválidos", "code": http.StatusUnauthorized})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Erro no servidor", "code": http.StatusInternalServerError})
	}

	// Compara a senha enviada com a senha criptografada armazenada no banco
	if err := bcrypt.CompareHashAndPassword([]byte(patient.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Email ou senha inválidos", "code": http.StatusUnauthorized})
	}

	// Se as credenciais estiverem corretas, gera um token JWT
	token, err := auth.GenerateToken(patient.ID, patient.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Falha ao gerar o token de autenticação", "code": http.StatusInternalServerError})
	}

	// Retorna o token para o cliente
	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
