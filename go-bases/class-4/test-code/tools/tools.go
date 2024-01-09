package tools

import "errors"

func ElementAtIndex(slice []int, index int) (result int, err error) {
	if index < 0 || index >= len(slice) {
		err = errors.New("index out of range")
		return
	}
	result = slice[index]
	return
}
