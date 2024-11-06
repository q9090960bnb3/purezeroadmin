package jsonutil

import (
	jsoniter "github.com/json-iterator/go"
)

func ToArray[T any](src string) ([]T, error) {
	var arr []T
	err := jsoniter.UnmarshalFromString(src, &arr)
	return arr, err
}

func ArrayContains[T comparable](src string, elem T) (bool, error) {
	var arr []T
	err := jsoniter.UnmarshalFromString(src, &arr)
	if err != nil {
		return false, err
	}

	for _, v := range arr {
		if v == elem {
			return true, nil
		}
	}

	return false, nil
}
