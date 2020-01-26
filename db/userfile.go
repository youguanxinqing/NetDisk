package db

import (
	"netdisk/db/mysql"
	"time"

	"log"
)

// UserFile 用户文件
type UserFile struct {
	UserName    string
	FileHash    string
	FileName    string
	FileSize    int64
	UploadAt    string
	LastUpdated string
}

// OnUserFileUploadFinished 上传完成
func OnUserFileUploadFinished(
	username, filehash, filename string,
	filesize int64,
) bool {
	stmt, err := mysql.DBConn().Prepare(
		"insert ignore " +
			"tbl_user_file(`user_name`, `file_sha1`, `file_name`, `file_size`, `upload_at`)" +
			"values(?,?,?,?,?)",
	)
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filename, filehash, filename, filesize, time.Now())
	if err != nil {
		log.Println(err)
		return false
	}
	if num, err := ret.RowsAffected(); err != nil || num <= 0 {
		log.Println("no insert operator (OnUserFileUploadFinished)")
		return false
	}

	return true
}
