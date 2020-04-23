package controllers

import (
	"fmt"
	"gofullstack/lenslocked.com/models"
	"gofullstack/lenslocked.com/views"
	"net/http"
)

func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us: us,
	}
}

type Users struct {
	NewView *views.View
	LoginView *views.View
	us *models.UserService
}

type SignupForm struct {
	Name string `schema:"name"`
	Email string `schema:"email"`
	Password string `schema:"password"`
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.User{
		Name: form.Name,
		Email: form.Email,
		Password: form.Password,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, form)
}

type LoginForm struct {
	Email string `schema:"email"`
	Password string `schema:"password"`
}

func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			fmt.Fprint(w, "Invalid email address.")
		case models.ErrInvalidPassword:
			fmt.Fprint(w, "Invalid password provided.")
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err != nil {
			return
		}

		cookie := http.Cookie{
			Name: "email",
			Value: user.Email,
		}
		http.SetCookie(w, &cookie)
	}
	fmt.Fprintln(w, form)
}