package db

import (
	"netdisk/db/mysql"
	"time"

	"log"
)

// UserFile 用户文件
type UserFile struct {
	UserName   string
	FileHash   string
	FileName   string
	FileSize   int64
	UploadAt   string
	LastUpdate string
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

	ret, err := stmt.Exec(username, filehash, filename, filesize, time.Now())
	if err != nil {
		log.Println(err)
		return false
	}
	// 允许没有插入实际内容
	if num, err := ret.RowsAffected(); err != nil || num <= 0 {
		log.Println("no insert operator (OnUserFileUploadFinished)")
		return true
	}

	return true
}

// QueryUserFileMetas ...
func QueryUserFileMetas(username string, limit int) ([]UserFile, error) {
	fileMetas := []UserFile{}

	stmt, err := mysql.DBConn().Prepare(
		"select `user_name`, `file_sha1`, `file_size`, `file_name`, `last_update`, `upload_at` " +
			"from tbl_user_file " +
			"where user_name=?",
	)
	if err != nil {
		return fileMetas, err
	}
	defer stmt.Close()

	if rows, err := stmt.Query(username); err == nil {
		for rows.Next() {
			var uf UserFile
			err := rows.Scan(
				&uf.UserName, &uf.FileHash, &uf.FileSize, &uf.FileName,
				&uf.LastUpdate, &uf.UploadAt,
			)
			if err != nil {
				log.Println(err.Error())
			}
			fileMetas = append(fileMetas, uf)
		}
	} else {
		return fileMetas, err
	}

	return fileMetas, nil
}
