package main

import "encoding/json"

func Stringify[T any](s T) string {
	data, err := json.Marshal(s)
	if err != nil {
		return "{}"
	}
	return string(data)
}
