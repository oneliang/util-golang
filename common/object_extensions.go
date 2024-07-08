package common

import (
	"errors"
	"fmt"
	"reflect"
)

func ObjectInList[V any](object V, list []V) bool {
	for _, item := range list {
		if reflect.DeepEqual(object, item) {
			return true
		}
	}
	return false
}

func ObjectNotInList[V any](object V, list []V) bool {
	return !ObjectInList(object, list)
}

func CheckNotNil(object any) error {
	if object == nil {
		return errors.New(fmt.Sprintf("object %+v is nil", object))
	}
	return nil
}
