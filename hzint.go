package hzutils

import "strconv"

// Atio64 return int64 from string
// @param string
// @return int64,error
func Atio64(s string) (int64, error) {
	i64, err := strconv.ParseInt(s, 10, 0)
	return i64, err
}
