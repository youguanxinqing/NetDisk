package meta

import (
	"log"
	"netdisk/db1"
)

// FileMeta 文件元信息
type FileMeta struct {
	FileSha1 string `json:"filehash"`
	FileName string `json:"filename"`
	FileSize int64  `json:"filesize"`
	Location string `json:"location"`
	UploadAt string `json:"update_at"`
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta 更新或新增
func UpdateFileMeta(fileMeta FileMeta) {
	fileMetas[fileMeta.FileSha1] = fileMeta
}

// UpdateFileMetaDB ...
func UpdateFileMetaDB(fileMeta FileMeta) bool {
	return db1.OnFileUploadFinished(fileMeta.FileSha1,
		fileMeta.FileName, fileMeta.FileSize, fileMeta.Location)
}

// GetFileMeta 获取文件元信息
func GetFileMeta(fileSha1 string) (FileMeta, bool) {
	if data, ok := fileMetas[fileSha1]; ok {
		return data, true
	}
	return FileMeta{}, false
}

// GetFileMetaDB 数据库查询
// TODO: 修改所有调用该函数的地方
func GetFileMetaDB(fileSha1 string) (*FileMeta, bool) {
	fileMeta, err := db1.GetFileMeta(fileSha1)
	if fileMeta == nil {
		log.Println("(GetFileMetaDB) " + err.Error())
		// 如果没有查询到内容
		return nil, false
	}

	return &FileMeta{
		FileSha1: fileMeta.FileHash,
		FileName: fileMeta.FileName.String,
		FileSize: fileMeta.FileSize.Int64,
		Location: fileMeta.FileAddr.String,
		UploadAt: fileMeta.UpdateAt.Local().String(),
	}, true
}

// GetLastFileMetas 获取批量文件元信息
func GetLastFileMetas(count int) []FileMeta {
	fileMetaArr := make([]FileMeta, len(fileMetas))
	for _, v := range fileMetas {
		fileMetaArr = append(fileMetaArr, v)
	}
	return fileMetaArr[:count]
}

// RemoveFileMeta 移除文件元信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}

// RemoveFileMetaDB 移除文件元信息(from db1)
func RemoveFileMetaDB(fileSha1 string) {
	err := db1.DeleteFileMeta(fileSha1)
	if err != nil {
		log.Println("occur ygerr while delete filemeta, err: " + err.Error())
	}
}
