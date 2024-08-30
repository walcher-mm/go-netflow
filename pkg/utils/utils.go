package utils

import (
	"encoding/json"
	"strings"
)

func MatchStringSuffix(s string, mv []string) bool {
	for _, val := range mv {
		if strings.HasSuffix(s, val) {
			return true
		}
	}
	return false
}

func MarshalIndent(v interface{}) string {
	bs, _ := json.MarshalIndent(v, "", "    ")
	return string(bs)
}
