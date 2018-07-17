package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mrsmuneton/platform-test/src/token"
	"github.com/mrsmuneton/platform-test/src/user"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest user.User
	body, _ := ioutil.ReadAll(r.Body)
	var _ = json.Unmarshal(body, &userRequest)

	userRecord, err := user.LoginUser(userRequest)
	fmt.Println(err)

	if err == false {
		_, token := token.CreateUserJWT(user.User{Email: userRecord.Email, Name: userRecord.Name})
		w.Write([]byte(token))
	} else {
		w.Write([]byte("sorry"))
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout Handler")

}

func UserRecordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ShowUser Handler")
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var newUser user.User
	body, err := ioutil.ReadAll(r.Body)
	var err1 = json.Unmarshal(body, &newUser)
	if err != nil || err1 != nil {
		fmt.Println(err)
		fmt.Println(err1)
		w.Write([]byte("create failed"))
		return
	}

	var createerr = user.CreateUser(newUser)

	if createerr == true {
		w.Write([]byte("create failed"))
	} else {
		w.Write([]byte("created"))
	}
}
