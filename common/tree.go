package common

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ListToTreeList[IdType int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string, Item interface{}](
	dataList []*Item,
	idFieldName string,
	parentIdFieldName string,
	parentIdValueList []IdType,
	childrenFieldName string,
	depthFieldName string,
) []*Item {
	var rootDataList []*Item
	var dataMap = make(map[IdType]*Item)
	for _, data := range dataList {
		dataElem := reflect.ValueOf(data).Elem()

		idField := dataElem.FieldByName(idFieldName)
		idValue := idField.Interface().(IdType)

		parentIdField := dataElem.FieldByName(parentIdFieldName)
		parentIdValue := parentIdField.Interface().(IdType)

		depthField := dataElem.FieldByName(depthFieldName)

		if SimpleObjectInList(parentIdValue, parentIdValueList) {
			if depthField.IsValid() {
				depthField.SetInt(1)
			}
			rootDataList = append(rootDataList, data)
		}
		dataMap[idValue] = data
	}

	//generate the tree
	for _, data := range dataList {
		dataElem := reflect.ValueOf(data).Elem()

		parentIdField := dataElem.FieldByName(parentIdFieldName)
		parentIdValue := parentIdField.Interface().(IdType)

		depthField := dataElem.FieldByName(depthFieldName)

		parentData := dataMap[parentIdValue]
		if parentData != nil {
			parentDataElem := reflect.ValueOf(parentData).Elem()

			parentChildrenField := parentDataElem.FieldByName(childrenFieldName)
			parentChildrenValue := parentChildrenField.Interface().([]*Item)

			parentDepthField := parentDataElem.FieldByName(depthFieldName)
			if parentDepthField.IsValid() && depthField.IsValid() {
				parentDepthValue := parentDepthField.Interface().(int32)
				depthField.SetInt(int64(parentDepthValue + 1))
			}

			parentChildrenValue = append(parentChildrenValue, data)

			parentChildrenField.Set(reflect.ValueOf(parentChildrenValue))
		}
	}

	return rootDataList
}

func PrintTreeList[V interface{}](rootDataList []*V) {
	jsonString, _ := json.Marshal(rootDataList)
	fmt.Printf("json:%v\n", string(jsonString))
}
