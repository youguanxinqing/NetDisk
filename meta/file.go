package meta

import (
	"netdisk/db1"
)

// FastUploadMetaDB ...
func FastUploadMetaDB(filehash string) (db1.TableFile, bool) {
	return db1.QueryFastUploadMeta(filehash)
}
