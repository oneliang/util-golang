package common

func ObjectInList[V interface{}](object *V, list []*V) bool {
	for _, item := range list {
		if item == object {
			return true
		}
	}
	return false
}
func SimpleObjectInList[V int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string](object V, list []V) bool {
	for _, item := range list {
		if item == object {
			return true
		}
	}
	return false
}
