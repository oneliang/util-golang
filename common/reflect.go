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

func ConvertType[T any](instance any, superInterface any) (error, T) {
	superInterfaceType := reflect.TypeOf(superInterface)
	if superInterfaceType.Kind() != reflect.Ptr || superInterfaceType.Elem().Kind() != reflect.Interface {
		var nilValue T
		return errors.New(fmt.Sprintf("Super interface is not interface type,  interface type:%v", superInterfaceType)), nilValue
	}
	instanceType := reflect.TypeOf(instance)
	instanceInterface := superInterfaceType.Elem()
	if !instanceType.Implements(instanceInterface) {
		var nilValue T
		return errors.New(fmt.Sprintf("Must implement interface, instance type:%v, interface type:%v", instanceType, superInterfaceType)), nilValue
	}
	return nil, instance.(T)
}

func SetPointerValueByReflect(instanceValue reflect.Value, value any) error {
	if instanceValue.Kind() != reflect.Ptr {
		return errors.New("only support pointer about instance value")
	}

	valueType := reflect.TypeOf(value)
	if valueType.Kind() != reflect.Ptr {
		return errors.New("only support pointer about value")
	}
	instanceValue.Set(reflect.ValueOf(value))
	return nil
}
