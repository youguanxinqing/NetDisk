package meta

import (
	"netdisk/db"
)

// UpdateUserFileDB 更新文件详情
func UpdateUserFileDB(
	username, filehash, filename string,
	filesize int64,
) bool {
	return db.OnUserFileUploadFinished(username, filehash, filename, filesize)
}
