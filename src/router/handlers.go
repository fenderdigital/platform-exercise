package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

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
		token, _ := token.CreateUserJWT(userRecord)
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
	var err bool
	var err1 bool
	var request_id string
	var sessionToken string
	var user_id string
	var userRequest user.User
	var valid bool
	userRequest, err = readUserRequestBody(r)
	userRequest, err1 = user.GetUserRecordByEmail(userRequest.Email)
	if err != false || err1 != false {
		fmt.Println(err)
		fmt.Println(err1)
		w.Write([]byte("error"))
		return
	}
	request_id = strconv.Itoa(userRequest.Id)

	sessionToken, user_id, valid = isAuthorizedToken(r)
	fmt.Println(user_id)
	fmt.Println(request_id)
	fmt.Println(valid)
	if user_id != request_id || valid == false {
		w.Write([]byte("Invalid credentials"))
		return
	}

	session.DeleteSession(sessionToken)

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
	var err bool
	var user_id string
	var userRequest user.User
	var valid bool

	userRequest, err = readUserRequestBody(r)
	fmt.Println(err)

	url_id := r.URL.Query().Get("id")

	_, user_id, valid = isAuthorizedToken(r)
	if user_id != url_id || valid == false {
		w.Write([]byte("Invalid credentials"))
		return
	}

	switch r.Method {
	case "DELETE":
		err = user.DeleteUser(url_id)
		userRequest = user.User{}
	case "GET":
		userRequest, err = user.GetUserRecordById(url_id)
	case "PATCH":
		// could implement new method in order to handle partial updates of the record
		userRequest, err = user.UpdateUserFields(url_id, userRequest)
	case "POST":
		// could implement new method that only accepts new records
		userRequest, err = user.UpdateUserFields(url_id, userRequest)
	case "PUT":
		userRequest, err = user.UpdateUserFields(url_id, userRequest)
	default:
		err = true
	}
	if err != true {
		b, marshalerr := json.Marshal(userRequest)
		fmt.Println(marshalerr)
		w.Write(b)
	} else {
		w.Write([]byte("error, please verify all required fields (email,name, password) are present"))
	}
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

func isAuthorizedToken(r *http.Request) (string, string, bool) {
	var user_id string
	var valid bool
	header := r.Header.Get("Authorization")
	jwToken := strings.Trim(strings.Replace(header, "Bearer", "", 1), " ")
	user_id, valid = token.ParseJWT(jwToken)
	return jwToken, user_id, valid
}
