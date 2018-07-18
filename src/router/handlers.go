package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mrsmuneton/platform-test/src/session"
	"github.com/mrsmuneton/platform-test/src/token"
	"github.com/mrsmuneton/platform-test/src/user"
	"github.com/mrsmuneton/platform-test/src/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest user.User
	var err bool
	userRequest, err = readUserRequestBody(r)
	if err == true {
		w.Write([]byte("unexpected login request decoding result"))
	}

	userRecord, err := user.LoginUser(userRequest)
	if err == true {
		w.Write([]byte("login error"))
	}

	if err == false {
		_, token := token.CreateUserJWT(user.User{Email: userRecord.Email, Name: userRecord.Name})
		t := utils.CurrentTime()
		newSession := session.Session{Email: userRecord.Email, Token: token, UpdatedDate: t}
		err := session.WriteSession(newSession)
		if err != false {
			w.Write([]byte("session not saved"))
		} else {
			w.Write([]byte(token))
		}
	} else {
		w.Write([]byte("sorry"))
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest user.User
	var err bool
	userRequest, err = readUserRequestBody(r)
	fmt.Println(err)

	session.DeleteSession(userRequest.Email)

	w.Write([]byte("Session deleted"))
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var newUser user.User
	var err bool
	newUser, err = readUserRequestBody(r)
	fmt.Println(err)

	var createerr = user.CreateUser(newUser)

	if createerr == true {
		w.Write([]byte("create failed"))
	} else {
		w.Write([]byte("created"))
	}
}

func UserRecordHandler(w http.ResponseWriter, r *http.Request) {
	// var userRequest user.User
	// var err bool
	// userRequest, err = readUserRequestBody(r)

	w.Write([]byte("wip"))
}

func readUserRequestBody(r *http.Request) (user.User, bool) {
	var userRequest user.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return user.User{}, true
	}

	var unmarshalErr = json.Unmarshal(body, &userRequest)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr.Error())
		return userRequest, true
	}

	return userRequest, false
}
