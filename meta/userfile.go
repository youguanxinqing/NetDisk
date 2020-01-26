package meta

import (
	"log"
	"netdisk/db"
)

// UpdateUserFileDB 更新文件详情
func UpdateUserFileDB(
	username, filehash, filename string,
	filesize int64,
) bool {
	return db.OnUserFileUploadFinished(username, filehash, filename, filesize)
}

// QueryUserFileDetails 查询文件详情
func QueryUserFileDetails(username string, limit int) []db.UserFile {
	if ufiles, err := db.QueryUserFileMetas(username, limit); err == nil {
		return ufiles
	} else {
		log.Println(err.Error())
	}
	return []db.UserFile{}
}
