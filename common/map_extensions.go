package common

import "reflect"

func MapToList[K comparable, V any, R any](inputMap map[K]V, transform func(key K, value V) R) []R {
	var list = make([]R, 0)

	if inputMap == nil {
		return list
	}

	for key, value := range inputMap {
		item := transform(key, value)
		list = append(list, item)
	}
	return list
}

func MapToNewMap[K comparable, V any, NK comparable, NV any](inputMap map[K]V, transform func(key K, value V) (NK, NV)) map[NK]NV {
	var newMap = make(map[NK]NV)

	if inputMap == nil || transform == nil {
		return newMap
	}

	for key, value := range inputMap {
		newKey, newValue := transform(key, value)
		newMap[newKey] = newValue
	}
	return newMap
}

func MapDiffersDefault[K comparable, V any](inputMap map[K]V, otherMap map[K]V) []K {
	return MapDiffers[K, V](inputMap, otherMap, func(_ K, inputValue V, otherValue V) bool {
		return reflect.DeepEqual(inputValue, otherValue)
	})
}

func MapDiffers[K comparable, V any](inputMap map[K]V, otherMap map[K]V, valueComparator func(inputKey K, inputValue V, otherValue V) bool) []K {
	var list []K
	for key, inputValue := range inputMap {
		otherValue, ok := otherMap[key]
		if !ok { //key not exists
			list = append(list, key)
		} else { //key exists
			if valueComparator != nil && !valueComparator(key, inputValue, otherValue) {
				list = append(list, key)
			}
		}
	}
	return list
}

func MapDiffersAccurate[K comparable, V any](inputMap map[K]V, otherMap map[K]V, valueComparator func(inputKey K, inputValue V, otherValue V) bool) ([]K, []K) {
	var list []K
	var valueCompareKeyList []K
	for key, inputValue := range inputMap {
		otherValue, ok := otherMap[key]
		if !ok { //key not exists
			list = append(list, key)
		} else { //key exists
			if valueComparator != nil && !valueComparator(key, inputValue, otherValue) {
				valueCompareKeyList = append(valueCompareKeyList, key)
			}
		}
	}
	return list, valueCompareKeyList
}

// MapKeys .
func MapKeys[K comparable, V any](inputMap map[K]V) []K {
	var keys []K
	for key, _ := range inputMap {
		keys = append(keys, key)
	}
	return keys
}

// MapValues .
func MapValues[K comparable, V any](inputMap map[K]V) []V {
	var values []V
	for _, value := range inputMap {
		values = append(values, value)
	}
	return values
}
