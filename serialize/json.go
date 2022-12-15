package serialize

import "encoding/json"

// JsonEncode json encode
func JsonEncode(v interface{}) string {
	return string(JsonEncodeByte(v))
}

// JsonEncodeByte json encode returns byte
func JsonEncodeByte(v interface{}) []byte {
	d, _ := json.Marshal(v)
	return d
}
