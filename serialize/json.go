package serialize

import "encoding/json"

func JsonEncode(v interface{}) string {
	d, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(d)
}
