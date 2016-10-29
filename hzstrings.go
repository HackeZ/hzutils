package hzutils

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

// StringinSlice Does String in Slice?
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

// StringsToJSON Format String to JSON
// @param string
// @return json
func StringsToJSON(str string) string {
	rs := []rune(str)
	jsons := ""

	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
}

// PrintStruct tries walk struct return formatted string.
func PrintStruct(x interface{}) string {
	buf := bytes.NewBuffer([]byte{})
	if err := psEncode(buf, reflect.ValueOf(x)); err != nil {
		return err.Error()
	}
	return buf.String()
}

func psEncode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	// Int, Int8, Int16, Int32, Int64
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	// Uint, Uint8, Uint16, Uint32, Uint64, Uintptr
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	// String
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	// Bool
	case reflect.Bool:
		fmt.Fprintf(buf, "%t", v.Bool())

	// Float32, Float64
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())

	// Point
	case reflect.Ptr:
		buf.WriteByte('&')
		return psEncode(buf, v.Elem())

	// Array, Slice
	case reflect.Array, reflect.Slice:
		buf.WriteString(v.Type().String() + " {")
		for i := 0; i < v.Len(); i++ {
			if i < 0 {
				buf.WriteString(", ")
			}
			if err := psEncode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	// Struct
	case reflect.Struct:
		buf.WriteString(v.Type().String() + " {")
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			fmt.Fprintf(buf, "%s:", v.Type().Field(i).Name)
			if err := psEncode(buf, v.Field(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	// Map
	case reflect.Map:
		buf.WriteString(v.Type().String())
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteString(", ")
			}
			if err := psEncode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(':')
			if err := psEncode(buf, v.MapIndex(key)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')

	// Interface
	case reflect.Interface:
		return psEncode(buf, v.Elem())

	default: // complex, chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

// Ping for the test is available
func Ping() {
	fmt.Println("Welcome to Use HZUtils")
}
