package yml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

func SplitByFiles(filepath []string) error {
	for _, path := range filepath {
		if files, err := Split(path); err != nil {
			fmt.Printf("[%s] Error: %s\n", path, err)
		} else {
			fmt.Printf("[%s] %s\n", path, strings.Join(files, ","))
		}
	}
	return nil
}

func Split(path string) ([]string, error) {
	files := make([]string, 0)
	ymlByte, err := ioutil.ReadFile(path)
	if err != nil {
		return files, err
	}
	var yamlObj interface{}
	if err := yaml.Unmarshal(ymlByte, &yamlObj); err != nil {
		return files, err
	}

	return files, nil
}
