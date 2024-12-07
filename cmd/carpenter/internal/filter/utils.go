package filter

import (
	"strings"
)

func getString(key string, object map[string]interface{}) string {
	if name, ok := object[key]; ok {
		if str, ok := name.(string); ok {
			return str
		}
	}
	return ""
}

func nameContains(nameSubstring string, object map[string]interface{}) bool {
	return strings.Contains(getString("logger", object), nameSubstring)
}
