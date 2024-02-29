package utils

import (
	"errors"
)

const NOT_FOUND_ERROR = "item not found" 

type IDer interface {
	GetId() uint
}

func GetIndexById[T IDer](id uint, arr []T) (int, error) {
	for i, item := range arr {
		if id == item.GetId() {
			return i, nil
		}
	}

	return -1, errors.New(NOT_FOUND_ERROR)
}

func GetItemById[T IDer](id uint, arr []T) (*T, error) {
	for i, item := range arr {
		if id == item.GetId() {
			return &arr[i], nil
		}
	}
	return nil, errors.New(NOT_FOUND_ERROR)
}

