package meta

import (
	"netdisk/db"
)

// FastUploadMetaDB ...
func FastUploadMetaDB(filehash string) (db.TableFile, bool) {
	return db.QueryFastUploadMeta(filehash)
}
