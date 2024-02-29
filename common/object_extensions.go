package common

import (
	"github.com/oneliang/util-golang/constants"
	"reflect"
)

func ObjectInList[V interface{}](object V, list []V) bool {
	for _, item := range list {
		if reflect.DeepEqual(object, item) {
			return true
		}
	}
	return false
}

func JoinToString[V interface{}](list []V, transform func(index int, item V) string, separator string) string {
	var results = ""
	length := len(list)
	for index, item := range list {
		results += transform(index, item)
		if index < length-1 {
			results += separator + constants.STRING_SPACE
		}
	}
	return results
}
