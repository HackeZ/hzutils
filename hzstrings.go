package hzutils

// StringinSlice Do String in Slice?
// @param string, []string
// @return bool
func StringinSlice(s string, slice []string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
