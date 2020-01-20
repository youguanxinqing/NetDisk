package meta

// FileMeta 文件元信息
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta 更新或新增
func UpdateFileMeta(fileMeta FileMeta) {
	fileMetas[fileMeta.FileSha1] = fileMeta
}

// GetFileMeta 获取文件元信息
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// GetLastFileMetas 获取批量文件元信息
func GetLastFileMetas(count int) []FileMeta {
	fileMetaArr := make([]FileMeta, len(fileMetas))
	for _, v := range fileMetas {
		fileMetaArr = append(fileMetaArr, v)
	}
	return fileMetaArr[:count]
}

// RemoveFileMeta ...
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}