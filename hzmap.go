package hzutils

import (
	"fmt"
	"reflect"
)

// MaxMapKeyLen return Max Length of Key plus ext.
// @param  interface{}, ext
// @return Max.Len() + ext
func MaxMapKeyLen(data interface{}, ext int) string {
	max := 0

	for _, k := range reflect.ValueOf(data).MapKeys() {
		if k.Len() > max {
			max = k.Len()
		}
	}

	return fmt.Sprintf("%d", max+ext)
}

// func MaxMapValue(data interface{}) interface{} {
// 	max := 0

// 	switch data.(type) {
// 	case int, int16, int32, int64, int8:
// 		for _, v := range reflect.ValueOf(data).MapKeys() {

// 		}
// 	}
// }
