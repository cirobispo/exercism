package flatten

func isSlice(v interface{}) bool {
	_,ok:=v.([]interface{})
	return ok
} 

func Flatten(nested interface{}) []interface{} {
	result:=make([]interface{}, 0)
	if isSlice(nested) {
		for i:=range nested.([]interface{}) {
			item:=nested.([]interface{})[i]
			if item != nil {
				if isSlice(item) {
					data:=Flatten(item)
					result = append(result, data...)
				} else {
					result = append(result, item)
				}
			}
		}
	} else {
		if value, ok:=nested.(int); ok {
			result = append(result, value)
		}
	}
	return result
}
