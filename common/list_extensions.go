package common

func ListToMapSimpleKey[V interface{}, K int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string](list []*V, keyTransform func(*V) K) map[K]*V {
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
