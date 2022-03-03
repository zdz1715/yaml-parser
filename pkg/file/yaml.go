package file

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
)

func SplitByFiles(filepath []string) error {
	for _, path := range filepath {
		if files, err := Split(path); err != nil {
			fmt.Printf("[%s] Error: %s\n", path, err)
		} else {
			fmt.Printf("[%s] %s\n", path, strings.Join(files, " "))
		}
	}
	return nil
}

func Split(filepath string) ([]string, error) {
	files := make([]string, 0)
	yamlContent, err := loadYamlToString(filepath)
	if err != nil {
		return files, err
	}
	ext := path.Ext(filepath)
	filenamePrefix := strings.TrimSuffix(filepath, ext)
	yamlSlice := strings.Split(yamlContent, "---")
	for k, v := range yamlSlice {
		str := strings.TrimSpace(v)
		tag := parserTag(v)
		if tag == "" {
			tag = strconv.Itoa(k)
		}
		filename := fmt.Sprintf("%s_%s%s", filenamePrefix, tag, ext)
		if err := ioutil.WriteFile(filename, []byte(str), 0755); err != nil {
			return files, err
		}
		files = append(files, filename)
	}
	return files, nil
}

func parserTag(str string) string {
	strSlice := Parse(str, "tag=")
	if len(strSlice) > 0 && strSlice[0].Value != "" {
		return strSlice[0].Value
	}
	return ""
}

func ParseParam(filepath string, key string) error {
	yamlContent, err := loadYamlToString(filepath)
	if err != nil {
		return err
	}
	params := Parse(yamlContent, "param:")
	if len(key) == 0 {
		jsonByte, err := json.Marshal(params)
		if err != nil {
			return err
		}
		fmt.Println(string(jsonByte))
		return nil
	} else {
		value := ""
		for _, v := range params {
			if v.Key == key {
				value = v.Value
			}
		}
		fmt.Println(value)
		return nil
	}
}

func loadYaml(filename string) ([]byte, error) {
	ymlByte, err := LoadFile(filename)
	if err != nil {
		return ymlByte, err
	}
	var yamlObj interface{}
	if err := yaml.Unmarshal(ymlByte, &yamlObj); err != nil {
		return ymlByte, err
	}
	return ymlByte, nil
}

func loadYamlToString(filename string) (string, error) {
	yamlByte, err := loadYaml(filename)
	if err != nil {
		return "", err
	}
	return string(yamlByte), nil
}
