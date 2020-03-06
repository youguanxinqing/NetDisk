package meta

import (
	"log"
	"netdisk/db1"
)

// UpdateUserFileDB 更新文件详情
func UpdateUserFileDB(
	username, filehash, filename string,
	filesize int64,
) bool {
	return db1.OnUserFileUploadFinished(username, filehash, filename, filesize)
}

// QueryUserFileDetails 查询文件详情
func QueryUserFileDetails(username string, limit int) []db1.UserFile {
	if ufiles, err := db1.QueryUserFileMetas(username, limit); err == nil {
		return ufiles
	} else {
		log.Println(err.Error())
	}
	return []db1.UserFile{}
}
