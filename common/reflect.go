package common

import (
	"errors"
	"fmt"
	"reflect"
)

// CopyDataFromMap only support simple data field in instance
func CopyDataFromMap(instance any, dataMap map[string]any, ignoreNotExistField bool) error {
	instanceValue := reflect.ValueOf(instance)
	if instanceValue.Kind() == reflect.Ptr {
		instanceValue = instanceValue.Elem()
	}

	if instanceValue.Kind() != reflect.Struct {
		return errors.New("instance must be a struct or a pointer to a struct")
	}

	for instanceKey, value := range dataMap {
		instanceField := instanceValue.FieldByName(instanceKey)
		if !instanceField.IsValid() {
			if ignoreNotExistField {
				continue
			} else {
				return fmt.Errorf("key: %s was not found in instance: %v", instanceKey, instance)
			}
		}
		if !instanceField.CanSet() {
			return fmt.Errorf("key: %s cannot be set in instance: %v", instanceKey, instance)
		}

		newValue := reflect.ValueOf(value)

		instanceFieldType := instanceField.Type()
		if instanceFieldType.Kind() == reflect.Ptr {
			//create a pointer and point to newValue
			newValuePointer := reflect.New(instanceFieldType.Elem())
			if newValuePointer.Elem().CanSet() {
				newValuePointer.Elem().Set(newValue)
			}
			newValue = newValuePointer //replace newValue
		} else if !newValue.Type().AssignableTo(instanceField.Type()) {
			return fmt.Errorf("value type for key: %s is not assignable to field type in instance: %v", instanceKey, instance)
		}

		instanceField.Set(newValue)
	}
	return nil
}
