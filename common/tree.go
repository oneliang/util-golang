package common

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type TreeNode[IdType int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string, V interface{}] struct {
	Id       IdType
	ParentId IdType
	Value    V
	Children []*TreeNode[IdType, V]
	Depth    int
}

func ListToTreeList[IdType int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string, Item interface{}](
	dataList []*Item,
	idFieldName string,
	parentIdFieldName string,
	parentIdValueList []IdType) []*TreeNode[IdType, *Item] {
	var rootTreeNodeList []*TreeNode[IdType, *Item]
	var treeNodeList []*TreeNode[IdType, *Item]
	var treeNodeMap = make(map[IdType]*TreeNode[IdType, *Item])
	for _, data := range dataList {
		dataElem := reflect.ValueOf(data).Elem()
		idValue := dataElem.FieldByName(idFieldName).Interface().(IdType)
		parentIdValue := dataElem.FieldByName(parentIdFieldName).Interface().(IdType)

		treeNode := &TreeNode[IdType, *Item]{
			Id:       idValue,
			ParentId: parentIdValue,
			Value:    data,
			Children: []*TreeNode[IdType, *Item]{},
			Depth:    0,
		}
		treeNodeList = append(treeNodeList, treeNode)

		if SimpleObjectInList(parentIdValue, parentIdValueList) {
			treeNode.Depth = 1
			rootTreeNodeList = append(rootTreeNodeList, treeNode)
		}
		treeNodeMap[treeNode.Id] = treeNode
	}
	//fmt.Printf("len:%d\n", len(treeNodeMap))
	//fmt.Printf("len:%d\n", len(rootTreeNodeList))
	//fmt.Printf("len:%d\n", len(treeNodeList))

	//generate the tree
	for _, treeNode := range treeNodeList {
		parentTreeNode := treeNodeMap[treeNode.ParentId]
		if parentTreeNode != nil {
			treeNode.Depth = parentTreeNode.Depth + 1
			parentTreeNode.Children = append(parentTreeNode.Children, treeNode)
		}
	}

	return rootTreeNodeList
}

func PrintTreeNodeList[IdType int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string, V interface{}](treeNodeList []*TreeNode[IdType, *V]) {
	//print the root  tree node
	for _, treeNode := range treeNodeList {
		fmt.Printf("root:%v\n", treeNode.Id)
		printTreeNode[IdType, V](treeNode)
		//rootTreeNode
	}
}
func printTreeNode[IdType int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string, V interface{}](treeNode *TreeNode[IdType, *V]) {
	if len(treeNode.Children) > 0 {
		for _, childTreeNode := range treeNode.Children {
			a, _ := json.Marshal(treeNode)
			fmt.Printf("child:%v, depth:%d, json:%v\n", childTreeNode.Id, childTreeNode.Depth, string(a))
			printTreeNode(childTreeNode)
		}
	}
}
