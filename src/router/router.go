package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrsmuneton/platform-test/src/token"
	"github.com/mrsmuneton/platform-test/src/user"
)

func Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/logout", LogoutHandler)
	r.HandleFunc("/register", RegistrationHandler).Methods("POST")
	r.HandleFunc("/user/{id:[0-9]+}", UserRecordHandler).Methods("GET")        //if READ right from jwt
	r.HandleFunc("/user/{id:[0-9]+}", UserRecordHandler).Methods("POST")       //if WRITE right from jwt
	r.HandleFunc("/user/{id:[0-9]+}", UserRecordHandler).Methods("PUT, PATCH") //if WRITE right from jwt
	r.HandleFunc("/user/{id:[0-9]+}", UserRecordHandler).Methods("DELETE")     //if DELETE right from jwt
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

func UserRecordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ShowUser Handler")
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	var newUser user.User
	body, err1 := ioutil.ReadAll(r.Body)
	var err2 = json.Unmarshal(body, &newUser)

	fmt.Println(err1)
	fmt.Println(err2)

	fmt.Println(newUser)
	fmt.Println(newUser.Name)
	w.Write([]byte("created"))
}
