package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


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

	if formDataErr := r.ParseForm(); formDataErr != nil {
		http.Error(w,formDataErr.Error(),http.StatusBadRequest)
		return
	}

	username:=r.FormValue("username")
	email:=r.FormValue("email")
	password:=r.FormValue("password")

	fmt.Println("This is the form data: ",username,email,password)

	if username == "" || email == "" || password == "" {
		http.Error(w,"username,password and email are required",http.StatusBadRequest)
		return
	}

	user,createUserError := h.service.createUser(r.Context(),&CreateUserReq{Username: username,Email: email,Password: password})

	if createUserError != nil {
		http.Error(w,createUserError.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func (h *Handler) login(w http.ResponseWriter,r *http.Request){
	if formError := r.ParseForm();formError!=nil{
		http.Error(w,formError.Error(),http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	

	if email == "" || password == "" {
		http.Error(w,"email and password both are required",http.StatusBadRequest)
		return
	}

	u,loginErr := h.service.login(r.Context(), LoginUserReq{Email: email,Password: password})

	if loginErr != nil {
		log.Println(loginErr.Error())
		http.Error(w, loginErr.Error(), http.StatusBadRequest)
		return
	}

	cookie := http.Cookie{
		Name: "jwt",
		Value: u.accessToken,
		Path: "/",
		Domain: "localhost",
		MaxAge: 60*60*24,
	}

	http.SetCookie(w,&cookie)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login Successfull"))
}

func (h *Handler) logout(w http.ResponseWriter, r *http.Request){

	cookie := http.Cookie{
		Name: "jwt",
		Value: "",
		Path: "/",
		MaxAge: -1,
	}

	http.SetCookie(w,&cookie)

	json.NewEncoder(w).Encode("logout successfull")

}