package yac

import "fmt"

func ErrInvalidCast(data interface{}, expected interface{}) error {
	return fmt.Errorf("can not cast '%t' to '%t'", data, expected)
}

func UnsafeCastToRoutePointer(data interface{}) (*Route) {
	val, _ := data.(*Route)
	return val
}