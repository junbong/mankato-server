package utils


func IsNotNilOrEmpty(str string) bool {
	if str != nil && len(str) > 0 {
		return true
	}
	
	return false
}


func IsNilOrEmpty(str string) bool {
	return !IsNotNilOrEmpty(str)
}
