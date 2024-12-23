package yaml

import (
	"fmt"
	yamlV3 "gopkg.in/yaml.v3"
	"os"
)

func LoadYamlToObject(path string, object any) error {
	// 加载配置
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to load file, path:%s", path))
		return nil
	}
	return yamlV3.Unmarshal(data, object)
}
