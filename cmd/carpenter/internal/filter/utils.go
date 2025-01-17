package filter

import (
	"fmt"
	"time"
)

func getTimestamp(key string, object map[string]interface{}) (time.Time, error) {
	str := getString(key, object)
	if str == "" {
		return time.Time{}, fmt.Errorf("could not find timestamp")
	}
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return time.Time{}, fmt.Errorf("bad time: %w", err)
	}
	return t, nil
}

func getString(key string, object map[string]interface{}) string {
	if name, ok := object[key]; ok {
		if str, ok := name.(string); ok {
			return str
		}
	}
	return ""
}
