package models

import (
	"database/sql"
	"fmt"

	"github.com/dani2159/echo-rest/databases"
	"github.com/dani2159/echo-rest/helpers"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := databases.CreateCon()

	sqlStatement := "SELECT * FROM tb_users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username Not FOund")
		return false, err
	}

	if err != nil {
		fmt.Println("Query Error")
		return false, err
	}

	match, err := helpers.ChechkPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Passsword doesn't match")
		return false, err
	}
	return true, nil
}
