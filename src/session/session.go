package session

import (
	"fmt"

	"github.com/mrsmuneton/platform-test/src/db"
)

type Session struct {
	Email       string `email`
	Token       string `token`
	UpdatedDate string `updated_date`
}

func WriteSession(CurrentSession Session) bool {
	dbConnection, err := db.DBConnect()
	_, queryerr := dbConnection.Query("INSERT INTO sessions(email, token, updated_date) VALUES($1,$2,$3);", CurrentSession.Email, CurrentSession.Token, CurrentSession.UpdatedDate)
	error_result := handleErrors(err, queryerr)
	return error_result
}

func DeleteSession(jwToken string) bool {
	dbConnection, err := db.DBConnect()
	_, queryerr := dbConnection.Query("DELETE FROM sessions WHERE token=$1", jwToken)
	error_result := handleErrors(err, queryerr)
	return error_result
}

func handleErrors(err error, queryerr error) bool {
	if err != nil || queryerr != nil {
		fmt.Println(err)
		fmt.Println(queryerr)
		return true
	} else {
		return false
	}
}
