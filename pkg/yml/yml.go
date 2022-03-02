package yml

import (
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
	ymlByte, err := ioutil.ReadFile(filepath)
	if err != nil {
		return files, err
	}
	var yamlObj interface{}
	if err := yaml.Unmarshal(ymlByte, &yamlObj); err != nil {
		return files, err
	}
	ext := path.Ext(filepath)
	filenamePrefix := strings.TrimSuffix(filepath, ext)
	fmt.Println(ext, filepath, filenamePrefix)
	yamlSlice := strings.Split(string(ymlByte), "---")
	for k, v := range yamlSlice {
		str := strings.TrimSpace(v)
		tag := parserTag(str)
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
	tag := ""
	fields := strings.Fields(str)
	key := "+parser:tag="
	keyLen := len(key)
	for _, v := range fields {
		if i := strings.Index(v, key); i >= 0 {
			tag = v[i+keyLen:]
		}
	}
	return tag
}
