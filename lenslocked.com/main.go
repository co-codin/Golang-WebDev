package main

import (
	"github.com/gorilla/mux"
	"gofullstack/lenslocked.com/views"
	"net/http"
)

var (
	homeView *views.View
	contactView *views.View
)



func main()  {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap","views/contact.gohtml")

	

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}

func must(err error)  {
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}