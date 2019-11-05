package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-playground/form"
)

var decoder *form.Decoder

// UserDTO get user details for signup
type UserDTO struct {
	UserName  string `form:"userName,omitempty"`
	FirstName string `form:"firstName,omitempty"`
	LastName  string `form:"lastName,omitempty"`
}

func signup(w http.ResponseWriter, r *http.Request) {
	var user UserDTO
	if r.Method == "POST" {
		r.ParseForm()

		decoder = form.NewDecoder()
		v := r.PostForm
		fmt.Println(v)
		err := decoder.Decode(&user, v)
		if err != nil {
			log.Panic(err)
		}

		fmt.Println("v:", user.UserName)
	}

	fmt.Fprintf(w, "Welcome "+user.FirstName) // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("login.html")
	t.Execute(w, nil)
}
func main() {
	http.HandleFunc("/", login) // setting router rule
	http.HandleFunc("/signup", signup)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
