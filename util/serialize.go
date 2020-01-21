package util

import (
	"encoding/json"
	"reflect"
)

// StructToJSON struct -> json
func StructToJSON(structure interface{}) string {
	var data []byte
	dataType := reflect.TypeOf(structure).Kind()
	if dataType == reflect.Struct {
		data, _ = json.Marshal(structure)
	} else {
		data, _ = json.Marshal(map[string]string{})
	}
	return string(data)
}
