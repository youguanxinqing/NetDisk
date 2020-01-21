package db

import (
	"log"
	mydb "netdisk/db/mysql"
)

// OnFileUploadFinished ...
func OnFileUploadFinished(
	filehash string,
	filename string,
	filesize int64,
	fileaddr string,
) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore tbl_file (`file_sha1`, `file_name`, `file_size`," +
			"`file_addr`, `status`) values (?,?,?,?,1)",
	)
	if err != nil {
		log.Println("failed to prepare sql, err: " + err.Error())
		return false
	}
	defer stmt.Close()
	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		log.Println(err.Error())
	}
	if rf, err := ret.RowsAffected(); err == nil {
		if rf <= 0 {
			log.Printf("file with hash: %v has been uploaded before", filehash)
		}
	}
	return true
}
