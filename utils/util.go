package utils


func IsNotNilOrEmpty(obj interface{}) bool {
	if obj != nil {
		switch obj.(type) {
		case string:
			return obj != ""
		
		case []byte:
			return len(obj.([]byte)) != 0
		
		default:
			return true
		}
	}
	
	return false
}


func IsNilOrEmpty(obj interface{}) bool {
	return !IsNotNilOrEmpty(obj)
}
