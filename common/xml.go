package common

import (
	"encoding/xml"
	"fmt"
	"os"
)

func LoadXmlToObject(path string, object any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to load file, path:%s", path))
		return nil
	}
	return xml.Unmarshal(data, object)
}
