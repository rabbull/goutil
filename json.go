package goutil

import "encoding/json"

func MustJsonize(val interface{}) string {
	bytes, err := json.Marshal(val)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func MarshalString(val interface{}) (string, error) {
	bytes, err := json.Marshal(val)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func UnmarshalString(s string, val interface{}) error {
	return json.Unmarshal([]byte(s), val)
}
