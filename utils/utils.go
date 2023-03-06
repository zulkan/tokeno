package utils

import (
	"encoding/json"
	"strconv"
)

func InSlice[T comparable](wanted []T, list []T) bool {
	mapList := map[T]bool{}
	for _, val := range list {
		mapList[val] = true
	}
	for _, val := range wanted {
		if _, exist := mapList[val]; !exist {
			return false
		}
	}
	return true
}

func MapSlice[T any, U any](list []T, convFunc func(data T) U) []U {
	res := []U{}
	for _, val := range list {
		res = append(res, convFunc(val))
	}
	return res
}

func ToJson[T any](data T) string {
	res, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(res)
}

func FromJson[T any](data string) T {
	var res T
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res
	}
	return res
}

// AllIsInt return all string array is int
func AllIsInt(list []string) bool {
	for _, val := range list {
		_, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return false
		}
	}
	return true
}
