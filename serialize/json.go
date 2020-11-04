package serialize

import "encoding/json"

// JsonEncode json encode
func JsonEncode(v interface{}) string {
	d, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(d)
}
