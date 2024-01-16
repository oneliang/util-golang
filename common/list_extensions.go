package common

func ListToMapSimpleKey[V interface{}, K int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string](list []*V, keyTransform func(*V) K) map[K]*V {
	resultMap := make(map[K]*V)
	for _, item := range list {
		key := keyTransform(item)
		resultMap[key] = item
	}
	return resultMap
}

func ListToMap[V interface{}, K interface{}](list []*V, keyTransform func(*V) *K) map[*K]*V {
	resultMap := make(map[*K]*V)
	for _, item := range list {
		key := keyTransform(item)
		resultMap[key] = item
	}
	return resultMap
}

func ListToNewList[V interface{}, NV interface{}](list []*V, valueTransform func(*V) *NV) []*NV {
	var newList []*NV
	for _, item := range list {
		newItem := valueTransform(item)
		newList = append(newList, newItem)
	}
	return newList
}

func ListFilter[V interface{}](list []*V, filter func(*V) bool) []*V {
	var newList []*V
	for _, item := range list {
		if filter(item) {
			newList = append(newList, item)
		}
	}
	return newList
}

func ListFilterToNewList[V interface{}, NV interface{}](list []*V, filter func(*V) bool, valueTransform func(*V) *NV) []*NV {
	var newList []*NV
	for _, item := range list {
		if filter(item) {
			newItem := valueTransform(item)
			newList = append(newList, newItem)
		}
	}
	return newList
}

func ListMinOf[V interface{}, R int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](list []*V, selector func(item *V) R) R {
	var minResult R
	for index, item := range list {
		value := selector(item)
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

func ListSumOf[V interface{}, R int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](list []*V, selector func(item *V) R) R {
	var sumResult R = 0
	for _, item := range list {
		sumResult += selector(item)
	}
	return sumResult
}
