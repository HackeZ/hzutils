package hzutils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
)

// ParseJSONFromFile parse JSON from file
func ParseJSONFromFile(jsonPath string, val interface{}) error {
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
