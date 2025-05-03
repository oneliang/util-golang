package common

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadJsonToObject(path string, object any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to load file, path:%s", path))
		return nil
	}
	return json.Unmarshal(data, object)
}
