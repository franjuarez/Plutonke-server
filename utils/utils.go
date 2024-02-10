package utils

import(
	"errors"
)

const NOT_FOUND_ERROR = "item not found" 

type IDer interface {
	GetId() string
}

// Función auxiliar para obtener el índice del elemento con el ID dado
func GetIndexById[T IDer](id string, arr []T) (int, error) {
	for i, item := range arr {
		if id == item.GetId() {
			return i, nil
		}
	}

	return -1, errors.New(NOT_FOUND_ERROR)
}

func GetItemById[T IDer](id string, arr []T) (*T, error) {
	for i, item := range arr {
		if id == item.GetId() {
			return &arr[i], nil
		}
	}
	return nil, errors.New(NOT_FOUND_ERROR)
}

