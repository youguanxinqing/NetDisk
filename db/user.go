package db

import (
	"log"
	"netdisk/db/mysql"
)

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

// GetToken ...
func GetToken(username string) (string, error) {
	var token string
	stmt, err := mysql.DBConn().Prepare(
		"select `user_token` from tbl_user_token where `user_name`=?",
	)
	if err != nil {
		return token, err
	}
	defer stmt.Close()

	if err := stmt.QueryRow(username).Scan(&token); err != nil {
		return token, err
	}
	return token, nil
}

// UpdateToken 更新 token
func UpdateToken(username, token string) bool {
	stmt, err := mysql.DBConn().Prepare(
		"replace into tbl_user_token(`user_name`, `user_token`) values(?,?)",
	)
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	if ret, err := stmt.Exec(username, token); err == nil {
		if num, err := ret.RowsAffected(); err == nil && num > 0 {
			return true
		} else if err != nil {
			log.Println(err)
		} else {
			log.Printf("UpdateToken does not take effect; num: %d", num)
		}
	} else {
		log.Println(err)
	}
	return false
}

// User ...
type User struct {
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}

// GetUserInfo ...
func GetUserInfo(username string) (User, error) {
	user := User{}
	stmt, err := mysql.DBConn().Prepare(
		"select user_name, signup_at " +
			"from tbl_user " +
			"where user_name=? limit 1",
	)
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	if err := stmt.QueryRow(username).Scan(
		&user.Username, &user.SignupAt,
	); err != nil {
		return user, err
	}
	return user, nil
}
