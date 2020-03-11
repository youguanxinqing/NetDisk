package settings

func _if(status bool, a interface{}, b interface{}) interface{} {
	if status {
		return a
	} else {
		return b
	}
}
