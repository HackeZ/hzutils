package hzutils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
)

// loadJSONFile load json file.
func loadJSONFile(jsonPath string, val interface{}) error {
	bs, err := ioutil.ReadFile(jsonPath)
	if err != err {
		return err
	}
	// 分隔不同行 JSON
	lines := strings.Split(string(bs), "\n")
	var bf bytes.Buffer
	// 忽略注释
	for _, line := range lines {
		lineNew := strings.TrimSpace(line)
		if (len(lineNew) > 0 && lineNew[0] == '#') || (len(lineNew) > 1 && lineNew[0:2] == "//") {
			continue
		}
		bf.WriteString(lineNew)
	}
	return json.Unmarshal(bf.Bytes(), &val)
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
