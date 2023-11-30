package common

func ListToMap[V interface{}, K interface{}](list []*V, keyTransform func(*V) *K) map[*K]*V {
	resultMap := make(map[*K]*V)
	for _, item := range list {
		key := keyTransform(item)
		resultMap[key] = item
	}
	return resultMap
}
