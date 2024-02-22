package common

func ListToMap[V interface{}, K comparable](list []V, keyTransform func(index int, item V) K) map[K]V {
	resultMap := make(map[K]V)
	for index, item := range list {
		key := keyTransform(index, item)
		resultMap[key] = item
	}
	return resultMap
}

func ListToNewMap[V interface{}, K comparable, NV interface{}](list []V, transform func(index int, item V) (K, NV)) map[K]NV {
	resultMap := make(map[K]NV)
	for index, item := range list {
		key, value := transform(index, item)
		resultMap[key] = value
	}
	return resultMap
}

func ListToNewList[V interface{}, NV interface{}](list []V, valueTransform func(index int, item V) NV) []NV {
	var newList []NV
	for index, item := range list {
		newItem := valueTransform(index, item)
		newList = append(newList, newItem)
	}
	return newList
}

func ListFilter[V interface{}](list []V, filter func(index int, item V) bool) []V {
	var newList []V
	for index, item := range list {
		if filter(index, item) {
			newList = append(newList, item)
		}
	}
	return newList
}

func ListFilterToNewList[V interface{}, NV interface{}](list []V, filter func(index int, item V) bool, valueTransform func(index int, item V) NV) []NV {
	var newList []NV
	for index, item := range list {
		if filter(index, item) {
			newItem := valueTransform(index, item)
			newList = append(newList, newItem)
		}
	}
	return newList
}

func ListMinOf[V interface{}, R NumberType](list []V, selector func(index int, item V) R) R {
	var minResult R
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

func ListSumOf[V interface{}, R NumberType](list []V, selector func(index int, item V) R) R {
	var sumResult R = 0
	for index, item := range list {
		sumResult += selector(index, item)
	}
	return sumResult
}

func ListGroupBy[V interface{}, K comparable](list []V, keySelector func(index int, item V) K) map[K][]V {
	var groupByMap = make(map[K][]V)
	for index, item := range list {
		key := keySelector(index, item)
		existItemList, exist := groupByMap[key]
		if exist {
			existItemList = append(existItemList, item)
			groupByMap[key] = existItemList
		} else { //not exist
			var itemList []V
			groupByMap[key] = itemList
		}
	}
	return groupByMap
}
