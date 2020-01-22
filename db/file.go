package db

import (
	"database/sql"
	"fmt"
	"log"
	mydb "netdisk/db/mysql"
	"time"
)

// OnFileUploadFinished ...
func OnFileUploadFinished(
	filehash string,
	filename string,
	filesize int64,
	fileaddr string,
) bool {
	// statement
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore tbl_file (`file_sha1`, `file_name`, `file_size`," +
			"`file_addr`, `status`) values (?,?,?,?,1)",
	)
	if err != nil {
		log.Println("failed to prepare sql, err: " + err.Error())
		return false
	}
	defer stmt.Close()
	// 插入
	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		log.Println(err.Error())
	}
	// 确定影响数
	if rf, err := ret.RowsAffected(); err == nil {
		if rf <= 0 {
			log.Printf("file with hash: %v has been uploaded before", filehash)
		}
	}
	return true
}

// TableFile ...
type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
	UpdateAt time.Time
}

// GetFileMeta 查询文件元信息
func GetFileMeta(filehash string) (*TableFile, error) {
	fmt.Println(filehash)
	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1, file_name, file_size, file_addr, update_at " +
			"from tbl_file " +
			"where file_sha1=? and status =1 limit 1",
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	tf := TableFile{}
	err = stmt.QueryRow(filehash).Scan(
		&tf.FileHash, &tf.FileName, &tf.FileSize, &tf.FileAddr, &tf.UpdateAt)
	if err != nil {
		log.Println(err.Error())
	}
	return &tf, nil
}

// DeleteFileMeta 删除文件元信息
func DeleteFileMeta(filehash string) error {
	// 逻辑删除
	stmt, err := mydb.DBConn().Prepare(
		"update tbl_file set status=0 where file_sha1=? and status=1",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash)
	if err != nil {
		return err
	}

	if num, err := ret.RowsAffected(); err == nil {
		if num <= 0 {
			return fmt.Errorf("no delete operation")
		}
	} else {
		return err
	}
	return nil
}
