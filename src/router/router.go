package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/logout", LogoutHandler)
	r.HandleFunc("/register", RegistrationHandler).Methods("POST")
	r.HandleFunc("/user", UserRecordHandler).Methods("GET")        //if READ right from jwt
	r.HandleFunc("/user", UserRecordHandler).Methods("POST")       //if WRITE right from jwt
	r.HandleFunc("/user", UserRecordHandler).Methods("PUT, PATCH") //if WRITE right from jwt
	r.HandleFunc("/user", UserRecordHandler).Methods("DELETE")     //if DELETE right from jwt
	http.Handle("/", r)

	return r
}
