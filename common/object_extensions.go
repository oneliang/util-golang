package common

import "reflect"

func ObjectInList[V interface{}](object V, list []V) bool {
	for _, item := range list {
		if reflect.DeepEqual(object, item) {
			return true
		}
	}
	return false
}
