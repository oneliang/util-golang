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
