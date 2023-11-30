package common

func MapToList[K interface{}, V interface{}, R interface{}](resultMap map[*K]*V, transform func(*K, *V) *R) []*R {
	var list []*R
	for key, value := range resultMap {
		item := transform(key, value)
		list = append(list, item)
	}
	return list
}
