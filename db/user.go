package db

import "netdisk/db/mysql"

import "log"

// UserSignUp 注册接口
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

// UserSignIn 登陆接口
func UserSignIn(username, encpasswd string) bool {
	stmt, err := mysql.DBConn().Prepare(
		"select * from tbl_user where user_name=? limit 1",
	)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		log.Println(err.Error())
		return false
	} else if rows == nil {
		log.Println("username not found: " + username)
		return false
	}

	ret := mysql.ParseRows(rows)[0]
	if pwd, ok := ret["user_pwd"]; ok {
		if pwdInt8, ok := pwd.([]uint8); ok && string(pwdInt8) == encpasswd {
			return true
		}
	}
	log.Println("password is error")
	return false
}
