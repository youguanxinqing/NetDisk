package util

import (
	"encoding/json"
	"reflect"
)

// JSON struct/map -> json
func JSON(structure interface{}) string {
	var data []byte
	dataType := reflect.TypeOf(structure).Kind()
	if dataType == reflect.Struct || dataType == reflect.Map {
		data, _ = json.Marshal(structure)
	} else {
		data, _ = json.Marshal(map[string]string{})
	}
	return string(data)
}
