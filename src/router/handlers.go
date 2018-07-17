package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kennygrant/sanitize"
	"github.com/mrsmuneton/platform-test/src/db"
	"github.com/mrsmuneton/platform-test/src/token"
	"github.com/mrsmuneton/platform-test/src/user"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest user.User
	body, _ := ioutil.ReadAll(r.Body)
	var _ = json.Unmarshal(body, &userRequest)
	dbConnection, queryerr := db.DBConnect()

	var userRecord user.User
	var query = "SELECT email, name FROM users WHERE currentpassword='" + userRequest.CurrentPassword + "' AND email='" + userRequest.Email + "'"
	fmt.Println(query)
	err := dbConnection.QueryRow(query).Scan(&userRecord.Email, &userRecord.Name)
	if queryerr != nil || err != nil {
		fmt.Println(err)
		fmt.Println(queryerr)
		w.Write([]byte("DB error "))
	}

	if err == nil {
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
	dbConnection, err := db.DBConnect()
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("DB connection error"))
	}

	var newUser user.User
	body, err := ioutil.ReadAll(r.Body)
	var err1 = json.Unmarshal(body, &newUser)
	if err != nil || err1 != nil {
		fmt.Println(err)
		fmt.Println(err1)
		w.Write([]byte("Registration body parsing error"))
	}

	//The below is defective code and will be removed in the next iteration
	var password = newUser.CurrentPassword //This must be encrypted with sha256 prior to storing, and should at least implement base64 from the client request
	t := currentTime()

	_, queryerr := dbConnection.Query("INSERT INTO users(created_date, currentpassword, email, name, updated_date) VALUES($1,$2,$3,$4,$5);", t, sanitize.Name(password), sanitize.Name(newUser.Email), sanitize.Name(newUser.Name), t)

	if queryerr != nil {
		w.Write([]byte("create failed"))
	} else {
		w.Write([]byte("created"))
	}
}

func currentTime() string {
	t := time.Now()
	//This should return a timestamp in order
	//Skimming over the rabbithole of time conversion for speed
	//the line below is defective code and it will be removed in the next iteration
	var now = t.String()
	return now
}
