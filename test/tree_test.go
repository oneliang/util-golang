package test

import (
	"github.com/oneliang/util-golang/common"
	"testing"
)

type Data struct {
	Id       string
	ParentId string
	Name     string
	Children []*Data
	Depth    int32
}

func TestTree(t *testing.T) {
	var dataList []*Data
	dataList = append(dataList, &Data{
		Id:       "1",
		ParentId: "",
		Name:     "Name_1(1)",
	}, &Data{
		Id:       "2",
		ParentId: "",
		Name:     "Name_2(2)",
	}, &Data{
		Id:       "3",
		ParentId: "1",
		Name:     "Name_1_1(3)",
	}, &Data{
		Id:       "4",
		ParentId: "1",
		Name:     "Name_1_2(4)",
	}, &Data{
		Id:       "5",
		ParentId: "3",
		Name:     "Name_1_1_1(5)",
	}, &Data{
		Id:       "6",
		ParentId: "3",
		Name:     "Name_1_1_2(6)",
	})
	rootDataList := common.ListToTreeList[string, Data](dataList, "Id", "ParentId", []string{""}, "Children", "Depth")
	common.PrintTreeList(rootDataList)
}
