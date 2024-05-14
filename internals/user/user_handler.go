package user

import "net/http"


type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		s,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /signup",h.createUser)
	router.HandleFunc("POST /login",h.login)
	router.HandleFunc("GET /logout",h.logout)

}

func (h *Handler) createUser(w http.ResponseWriter,r *http.Request) {

}

func (h *Handler) login(w http.ResponseWriter,r *http.Request){

}

func (h *Handler) logout(w http.ResponseWriter, r *http.Request){

}