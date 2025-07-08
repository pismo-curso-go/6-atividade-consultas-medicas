package dto

type RegisterPatientRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPatientRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
