package common

import (
	"github.com/oneliang/util-golang/base"
	"strings"
)

func ListToMap[V any, K comparable](list []V, keyTransform func(index int, item V) K) map[K]V {
	resultMap := make(map[K]V)

	if list == nil || keyTransform == nil {
		return resultMap

	}
	for index, item := range list {
		key := keyTransform(index, item)
		resultMap[key] = item
	}
	return resultMap
}

func ListToNewMap[V any, K comparable, NV any](list []V, transform func(index int, item V) (K, NV)) map[K]NV {
	resultMap := make(map[K]NV)

	if list == nil || transform == nil {
		return resultMap

	}
	for index, item := range list {
		key, value := transform(index, item)
		resultMap[key] = value
	}
	return resultMap
}

func ListToNewList[V any, NV any](list []V, valueTransform func(index int, item V) NV) []NV {
	var newList = make([]NV, 0)

	if list == nil || valueTransform == nil {
		return newList
	}

	for index, item := range list {
		newItem := valueTransform(index, item)
		newList = append(newList, newItem)
	}
	return newList
}

func ListFilter[V any](list []V, filter func(index int, item V) bool) []V {
	var newList = make([]V, 0)

	if list == nil {
		return newList
	}

	for index, item := range list {
		if filter != nil && filter(index, item) {
			newList = append(newList, item)
		}
	}
	return newList
}

func ListFilterToNewList[V any, NV any](list []V, filter func(index int, item V) bool, valueTransform func(index int, item V) NV) []NV {
	var newList = make([]NV, 0)

	if list == nil || valueTransform == nil {
		return newList
	}

	for index, item := range list {
		if filter != nil && filter(index, item) {
			newItem := valueTransform(index, item)
			newList = append(newList, newItem)
		}
	}
	return newList
}

func ListMinOf[V any, R base.NumberType](list []V, selector func(index int, item V) R) R {
	var minResult R

	if list == nil || selector == nil {
		return minResult
	}

	for index, item := range list {
		value := selector(index, item)
		if index == 0 {
			minResult = value
			continue
		}
		if value < minResult {
			minResult = value
		}
	}
	return minResult
}

func ListSumOf[V any, R base.NumberType](list []V, selector func(index int, item V) R) R {
	var sumResult R = 0

	if list == nil || selector == nil {
		return sumResult
	}

	for index, item := range list {
		sumResult += selector(index, item)
	}
	return sumResult
}

func ListGroupBy[V any, K comparable](list []V, keySelector func(index int, item V) K) map[K][]V {
	var groupByMap = make(map[K][]V)

	if list == nil || keySelector == nil {
		return groupByMap
	}

	for index, item := range list {
		key := keySelector(index, item)
		existItemList, exist := groupByMap[key]
		if !exist {
			//value not exist in map
			existItemList = []V{}
		}
		existItemList = append(existItemList, item)
		groupByMap[key] = existItemList
	}
	return groupByMap
}

func ListJoinToString[V any](list []V, transform func(index int, item V) string, separator string) string {
	var results strings.Builder

	if list == nil || transform == nil {
		return results.String()
	}

	length := len(list)
	for index, item := range list {
		results.WriteString(transform(index, item))
		if index < length-1 {
			results.WriteString(separator)
		}
	}
	return results.String()
}

func ListJoinToStringWithMaxCount[V any](list []V, transform func(index int, item V) string, maxCount int, separator string) string {
	var results strings.Builder

	if list == nil || transform == nil {
		return results.String()
	}

	length := len(list)
	if maxCount <= 0 || maxCount >= length {
		maxCount = length
	}
	for index, item := range list {
		if index >= maxCount {
			break
		}
		results.WriteString(transform(index, item))
		if index < maxCount-1 {
			results.WriteString(separator)
		}
	}
	return results.String()
}
