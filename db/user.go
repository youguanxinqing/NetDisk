package db

import "netdisk/db/mysql"

import "log"

// UserSignUp ...
func UserSignUp(username, passwd string) bool {
	stmt, err := mysql.DBConn().Prepare(
		"insert ignore tbl_user(`user_name`, `user_pwd`) values(?,?)",
	)
	if err != nil {
		log.Println("Failed to signup, err: " + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, passwd)
	if err != nil {
		log.Println("Failed to signup, err: " + err.Error())
		return false
	}

	if num, err := ret.RowsAffected(); err == nil && num > 0 {
		return true
	} else if err != nil {
		log.Println("Failed to signup, err: " + err.Error())
	} else {
		log.Println("user existed")
	}
	return false
}
