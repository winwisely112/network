package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/go-playground/form"
)

var decoder *form.Decoder

func getEnvOr(key string, orValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return orValue
	}
	return val
}

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

		//fmt.Println(v)
		log.Printf("v %v:", v)
		err := decoder.Decode(&user, v)
		if err != nil {
			log.Panic(err)
		}

		log.Printf("user.UserName %s:", user.UserName)
	}

	fmt.Fprintf(w, "Welcome "+user.FirstName) // write data to response
	log.Printf(format string, v ...interface{})
	log.Printf("user.UserName %s:", user.UserName)
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
