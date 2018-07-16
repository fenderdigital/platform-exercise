package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mrsmuneton/platform-test/src/db"
	"github.com/mrsmuneton/platform-test/src/token"
	"github.com/mrsmuneton/platform-test/src/user"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest user.User
	body, _ := ioutil.ReadAll(r.Body)
	var _ = json.Unmarshal(body, &userRequest)

	database := db.DBConnect()
	err := database.QueryRow("SELECT * FROM users WHERE currentpassword = $1 email = $2;", &userRequest.CurrentPassword, &userRequest.Email).Scan(&userRequest.Id, &userRequest.CurrentPassword, &userRequest.Email, &userRequest.Name)
	fmt.Println(err)

	_, token := token.CreateUserJWT(user.User{Email: userRequest.Email, Name: userRequest.Name})

	w.Write([]byte(token))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout Handler")

}

func UserRecordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ShowUser Handler")
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	database := db.DBConnect()
	var newUser user.User
	body, err1 := ioutil.ReadAll(r.Body)
	var err2 = json.Unmarshal(body, &newUser)

	fmt.Println(err1)
	fmt.Println(err2)
	u, err := database.Query("INSERT INTO users(currentpassword, email, name) VALUES($1,$2,$3);", newUser.CurrentPassword, newUser.Email, newUser.Name)
	fmt.Println(err)
	fmt.Println(u)

	w.Write([]byte("created"))
}
