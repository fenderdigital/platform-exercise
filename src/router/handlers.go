package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
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
	var err bool
	var valid bool
	var userRequest user.User
	userRequest, err = readUserRequestBody(r)
	fmt.Println(err)

	url_id := r.URL.Query().Get("id")

	switch r.Method {
	case "DELETE":
		err = user.DeleteUser(url_id)
		userRequest = user.User{}
	case "GET":
		valid = validToken(r)
		fmt.Println(valid)
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

func validToken(r *http.Request) bool {
	bearer := r.Header.Get("Authorization")
	if bearer == "" {
		return false
	}
	token, _ := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		return []byte("2hZ2cpjxFSz8sR2MbKqo7XLz4HS6Nx4tuBWlLpvIrXQPR5O36syvcefGZAdbZisog9LWPvDCYEJajl9X"), nil
	})
	fmt.Println(token)
	// fmt.Println(err)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		fmt.Println("fail")
	}
	return true
}
