package util

import "reflect"

// IsNumber returns true if the argument is a numeric type; false otherwise.
func IsNumber(x interface{}) bool {
	switch reflect.TypeOf(x).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

// IsString returns true if the argument is a string; false otherwise.
func IsString(x interface{}) bool {
	if reflect.TypeOf(x).Kind() == reflect.String {
		return true
	}
	return false
}
