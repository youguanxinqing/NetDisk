package util

import "encoding/json"

// ExplainError 对前端解释异常
func ExplainError(info string) []byte {
	structure := map[string]string{
		"detail": info,
	}
	data, _ := json.Marshal(structure)
	return data
}
