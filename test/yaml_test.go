package test

import (
	"fmt"
	"github.com/oneliang/util-golang/yaml"
	"log"
	"testing"
)

type Config struct {
	ApkFile     string `yaml:"apkFile"`
	ApkMd5      string `yaml:"apkMd5"`
	VersionCode int    `yaml:"versionCode"`
	VersionName string `yaml:"versionName"`
}

func TestYaml(t *testing.T) {
	config := &Config{}
	err := yaml.LoadYamlToObject("config.yaml", config)
	if err != nil {
		log.Fatalf("yaml.LoadYamlToObject error:%v", err)
		return
	}
	fmt.Println(config)
}
