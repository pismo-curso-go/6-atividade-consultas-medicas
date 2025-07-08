package handler

import (
	"api/banco"
	"api/model"
	"api/repository"
	"api/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func extractToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if !strings.HasPrefix(token, "Bearer ") {
		return "", http.ErrNoCookie
	}
	return strings.TrimPrefix(token, "Bearer "), nil
}

func CreateAppointment(w http.ResponseWriter, r *http.Request) {
	token, err := extractToken(r)
	if err != nil {
		utils.ResponseWithError(w, "Não autorizado", http.StatusUnauthorized)
	}

	patientID, err := utils.ParseJwt(token)
	if err != nil {
		utils.ResponseWithError(w, "Token expirado ou inválido", http.StatusUnauthorized)
		return
	}

	var body struct {
		Datetime string `json:"datetime"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.ResponseWithError(w, "JSON Inválido", http.StatusBadRequest)
	}

	datetime, err := time.Parse("2006-01-02T15:04:05", body.Datetime)
	if err != nil {
		utils.ResponseWithError(w, "Formato de data inválido", http.StatusBadRequest)
		return
	}

	appointment, err := repository.CreateAppointment(patientID, datetime)
	if err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(appointment); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func ListAppointments(w http.ResponseWriter, r *http.Request) {
	token, err := extractToken(r)
	if err != nil {
		utils.ResponseWithError(w, "Acesso não autorizado", http.StatusUnauthorized)
		return
	}

	patientID, err := utils.ParseJwt(token)
	if err != nil {
		utils.ResponseWithError(w, "Token expirado ou inválido", http.StatusUnauthorized)
		return
	}

	appointments, err := repository.GetAppointmentsByPatientID(patientID)
	if err != nil {
		utils.ResponseWithError(w, "Nenhum agendamento encontrado", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(appointments); err != nil {
		http.Error(w, "Erro ao codificar JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	token, err := extractToken(r)
	if err != nil {
		utils.ResponseWithError(w, "Acesso não autorizado", http.StatusUnauthorized)
		return
	}

	patientID, err := utils.ParseJwt(token)
	if err != nil {
		utils.ResponseWithError(w, "Token expirado ou inválido", http.StatusUnauthorized)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/appointments/")
	if err := repository.DeleteAppointment(patientID, id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func RegistrarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var usuario *model.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := usuario.Validar(); err != nil {
		http.Error(w, "Erro de validação: "+err.Error(), http.StatusBadRequest)
		return
	}

	repo := repository.NovoRepositorioDeUsuarios(banco.DB)
	id, err := repo.CriarUsuario(*usuario)
	if err != nil {
		http.Error(w, "Erro ao criar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Usuário criado com sucesso",
		"id":      id,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var creds model.CredenciaisLogin
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Erro ao decodificar JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if creds.Email == "" || creds.Senha == "" {
		http.Error(w, "E-mail e senha são obrigatórios", http.StatusBadRequest)
		return
	}

	repo := repository.NovoRepositorioDeUsuarios(banco.DB)
	usuario, err := repo.BuscarUsuarioPorEmail(creds.Email)
	if err != nil {
		http.Error(w, "Erro ao buscar usuário: "+err.Error(), http.StatusUnauthorized)
		return
	}

	if usuario == nil {
		http.Error(w, "Usuário não encontrado ou senha incorreta", http.StatusUnauthorized)
		return
	}

	if err := repository.CompararSenha(usuario.Senha, creds.Senha); err != nil {
		http.Error(w, "Usuário não encontrado ou senha incorreta", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJwt(strconv.FormatInt(usuario.ID, 10))
	if err != nil {
		log.Println("Erro ao gerar token JWT:", err)
		http.Error(w, "Erro ao gerar token JWT:"+err.Error(), http.StatusInternalServerError)
		return
	}

	// Resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login realizado com sucesso",
		"token":   token,
	})
}
