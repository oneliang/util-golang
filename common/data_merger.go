package common

import (
	"github.com/oneliang/util-golang/constants"
	"strconv"
)

type MergerConfig struct {
	MasterDataList []*map[string]any
	SlaveDataList  []*SlaveData
	StaticData     *map[string]any //no merge key, merge data to every data map when output
}

type SlaveData struct {
	DataList  []*map[string]any
	MergeKeys []string
}

type slaveMapData struct {
	dataMap   *map[string]*map[string]any
	mergeKeys []string
}

func generateKey(data map[string]any, mergeKeys []string) string {
	return ListJoinToString[string](mergeKeys, func(index int, item string) string {
		value, exist := data[item]
		if !exist {
			value = constants.STRING_BLANK
		}
		switch value.(type) {
		case int:
			return strconv.FormatInt(int64(value.(int)), 10)
		case int8:
			return strconv.FormatInt(int64(value.(int8)), 10)
		case int16:
			return strconv.FormatInt(int64(value.(int16)), 10)
		case int32:
			return strconv.FormatInt(int64(value.(int32)), 10)
		case int64:
			return strconv.FormatInt(value.(int64), 10)
		}
		return value.(string)
	}, constants.SYMBOL_COMMA)
}

func Merge(mergerConfig *MergerConfig) []*map[string]any {
	if mergerConfig == nil {
		return nil
	}

	masterDataList := mergerConfig.MasterDataList
	slaveDataList := mergerConfig.SlaveDataList
	staticDataPointer := mergerConfig.StaticData

	dataList := make([]*map[string]any, len(masterDataList))

	var slaveMapDataList []*slaveMapData = nil
	if slaveDataList != nil {
		slaveMapDataList = ListToNewList[*SlaveData, *slaveMapData](slaveDataList, func(index int, item *SlaveData) *slaveMapData {
			if item == nil {
				return nil
			}
			mergeKeys := item.MergeKeys
			mapData := ListToMap[*map[string]any, string](item.DataList, func(index int, innerItemPointer *map[string]any) string {
				if innerItemPointer == nil {
					return constants.STRING_BLANK
				}
				innerItem := *innerItemPointer
				return generateKey(innerItem, mergeKeys)
			})
			return &slaveMapData{
				dataMap:   &mapData,
				mergeKeys: mergeKeys,
			}
		})
	}

	for i := 0; i < len(masterDataList); i++ {
		itemPointer := masterDataList[i]
		item := *itemPointer
		masterData := MapToNewMap[string, any, string, any](item, func(key string, value any) (string, any) {
			return key, value
		})
		dataList[i] = &masterData

		//find static data map and append to master data
		if staticDataPointer != nil {
			staticData := *staticDataPointer
			for staticDataKey, staticDataValue := range staticData {
				masterData[staticDataKey] = staticDataValue
			}
		}

		if slaveMapDataList == nil {
			continue
		}
		//find slave data map and append to master data
		for _, slaveMapDataPointer := range slaveMapDataList {
			mergeKeys := slaveMapDataPointer.mergeKeys //need to merge keys
			mergeKeysValue := generateKey(masterData, mergeKeys)

			slaveDataMap := *(slaveMapDataPointer.dataMap)
			slaveDataPointer, exist := slaveDataMap[mergeKeysValue]
			if !exist {
				continue //next data
			}
			slaveData := *slaveDataPointer
			for slaveDataKey, slaveDataValue := range slaveData {
				masterData[slaveDataKey] = slaveDataValue
			}
		}
	}

	return dataList
}
