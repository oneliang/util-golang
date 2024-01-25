package common

import "reflect"

func MapToList[K SimpleTypeAndStruct, V interface{}, R interface{}](inputMap map[K]V, transform func(K, V) R) []R {
	var list []R
	for key, value := range inputMap {
		item := transform(key, value)
		list = append(list, item)
	}
	return list
}

func MapToNewMap[K SimpleTypeAndStruct, V interface{}, NK SimpleTypeAndStruct, NV interface{}](inputMap map[K]V, transform func(K, V) (NK, NV)) map[NK]NV {
	var newMap = make(map[NK]NV)
	for key, value := range inputMap {
		newKey, newValue := transform(key, value)
		newMap[newKey] = newValue
	}
	return newMap
}

func MapDiffersDefault[K SimpleTypeAndStruct, V interface{}](inputMap map[K]V, otherMap map[K]V) []K {
	return MapDiffers[K, V](inputMap, otherMap, func(_ K, inputValue V, otherValue V) bool {
		return reflect.DeepEqual(inputValue, otherValue)
	})
}

func MapDiffers[K SimpleTypeAndStruct, V interface{}](inputMap map[K]V, otherMap map[K]V, valueComparator func(inputKey K, inputValue V, otherValue V) bool) []K {
	var list []K
	for key, inputValue := range inputMap {
		otherValue, ok := otherMap[key]
		if !ok { //key not exists
			list = append(list, key)
		} else { //key exists
			if !valueComparator(key, inputValue, otherValue) {
				list = append(list, key)
			}
		}
	}
	return list
}

func MapDiffersAccurate[K SimpleTypeAndStruct, V interface{}](inputMap map[K]V, otherMap map[K]V, valueComparator func(inputKey K, inputValue V, otherValue V) bool) ([]K, []K) {
	var list []K
	var valueCompareKeyList []K
	for key, inputValue := range inputMap {
		otherValue, ok := otherMap[key]
		if !ok { //key not exists
			list = append(list, key)
		} else { //key exists
			if !valueComparator(key, inputValue, otherValue) {
				valueCompareKeyList = append(valueCompareKeyList, key)
			}
		}
	}
	return list, valueCompareKeyList
}
