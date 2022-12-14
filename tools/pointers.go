package tools

import "reflect"

// IsNil checks if the value is actually nil
// This function is necessary to check interfaces
func IsNil(v interface{}) bool {
	return (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}
