package router

import (
	"api/handler"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/appointments", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateAppointment(w, r)
		case http.MethodGet:
			handler.ListAppointments(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/appointments/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			handler.DeleteAppointment(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/register", handler.RegistrarUsuario)
	http.HandleFunc("/login", handler.Login)
}
