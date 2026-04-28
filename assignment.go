package assignment

import "errors"

var ErrOverflow = errors.New("overflow")

func SafeAdd(a, b int32) (int32, error) {
	panic("implement me")
}
