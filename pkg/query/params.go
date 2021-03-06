package query

import (
	"strconv"
	"strings"
)

// GetBool converts query params to boolean
func GetBool(name string) bool {
	//params, err := url.ParseQuery(r.URL.RawQuery)
	//if err != nil {
	//	return false
	//}

	//value, ok := params[name]
	//if !ok {
	//	return false
	//}

	strValue := strings.Join([]string{"", ""}, "")
	if strValue == "" {
		return true
	}

	boolValue, err := strconv.ParseBool(strValue)
	if err != nil {
		return false
	}

	return boolValue
}
