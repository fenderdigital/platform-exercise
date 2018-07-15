package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrsmuneton/platform-test/src/token"
	"github.com/mrsmuneton/platform-test/src/user"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/logout", LogoutHandler)
	r.HandleFunc("/user/{id:[0-9]+}", ShowUserHandler).Methods("GET")        //if READ right from jwt
	r.HandleFunc("/user/{id:[0-9]+}", ShowUserHandler).Methods("POST")       //if WRITE right from jwt
	r.HandleFunc("/user/{id:[0-9]+}", ShowUserHandler).Methods("PUT, PATCH") //if WRITE right from jwt
	r.HandleFunc("/user/{id:[0-9]+}", ShowUserHandler).Methods("DELETE")     //if DELETE right from jwt
	http.Handle("/", r)

	return r
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	_, token := token.CreateUserJWT(user.User{CurrentPassword: "P0werpuff", Email: "cinnamon@nice.com", Name: "Ray May"})
	w.Write([]byte(token))
}
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout Handler")

}
func ShowUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ShowUser Handler")
}
