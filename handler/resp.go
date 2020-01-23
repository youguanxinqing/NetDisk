package handler

import (
	"encoding/json"
	"log"
)

// ResponseJSON ...
func ResponseJSON(data interface{}) string {
	respData := map[string]interface{}{}
	respData["data"] = data
	if dataBytes, err := json.Marshal(respData); err == nil {
		return string(dataBytes)
	} else {
		log.Println(err.Error())
		return "{}"
	}
}
