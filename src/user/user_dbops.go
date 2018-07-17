package user

import (
	"fmt"

	"github.com/mrsmuneton/platform-test/src/db"
	"github.com/mrsmuneton/platform-test/src/utils"
)

func CreateUser(newUser User) bool {
	var error_result = false
	dbConnection, err := db.DBConnect()
	if err != nil {
		fmt.Println(err)
		error_result = true
	}

	//The below is defective code and will be removed in the next iteration
	var password = newUser.CurrentPassword //This must be encrypted with sha256 prior to storing, and should at least implement base64 from the client request
	t := utils.CurrentTime()

	_, queryerr := dbConnection.Query("INSERT INTO users(created_date, currentpassword, email, name, updated_date) VALUES($1,$2,$3,$4,$5);", t, password, newUser.Email, newUser.Name, t)
	if queryerr != nil {
		fmt.Println(queryerr)
		error_result = true
	}

	return error_result
}

func LoginUser(userRequest User) (User, bool) {
	var errorBool = false
	dbConnection, queryerr := db.DBConnect()
	var userRecord User
	var query = "SELECT email, name FROM users WHERE currentpassword='" + userRequest.CurrentPassword + "' AND email='" + userRequest.Email + "'"
	err := dbConnection.QueryRow(query).Scan(&userRecord.Email, &userRecord.Name)
	if queryerr != nil || err != nil {
		fmt.Println(err)
		fmt.Println(queryerr)
		errorBool = true
	}
	return userRecord, errorBool
}
