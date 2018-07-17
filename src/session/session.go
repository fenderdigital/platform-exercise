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
	var error_result = false
	dbConnection, err := db.DBConnect()
	_, queryerr := dbConnection.Query("INSERT INTO sessions(email, token, updated_date) VALUES($1,$2,$3);", CurrentSession.Email, CurrentSession.Token, CurrentSession.UpdatedDate)
	if err != nil || queryerr != nil {
		fmt.Println(err)
		fmt.Println(queryerr)
		error_result = true
	}
	return error_result
}
