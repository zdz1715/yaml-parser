package file

import (
	"io/ioutil"
	"strings"
)

const RulePrefix = "+parser:"
const RuleSeparator = ","
const RuleValueSeparator = "="

type ruleValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func LoadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func Parse(fileContent, key string) []ruleValue {
	ruleKey := RulePrefix + key
	fields := strings.Fields(fileContent)
	keyLen := len(ruleKey)
	lineSlice := make([]string, 0)
	ruleValueSlice := make([]ruleValue, 0)
	for _, v := range fields {
		if i := strings.Index(v, ruleKey); i >= 0 {
			lineSlice = append(lineSlice, v[i+keyLen:])
		}
	}
	if len(lineSlice) == 0 {
		return ruleValueSlice
	}
	for _, line := range lineSlice {
		strArr := strings.Split(line, RuleSeparator)
		for _, v := range strArr {
			strTemp := strings.Split(v, RuleValueSeparator)
			if len(strTemp) >= 2 {
				ruleValueSlice = append(ruleValueSlice, ruleValue{
					Key:   strTemp[0],
					Value: strTemp[1],
				})
			} else {
				ruleValueSlice = append(ruleValueSlice, ruleValue{
					Key:   "",
					Value: strTemp[0],
				})
			}
		}
	}

	return ruleValueSlice
}
