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
	str := ""
	for _, v := range fields {
		if i := strings.Index(v, ruleKey); i >= 0 {
			str = v[i+keyLen:]
		}
	}
	if len(str) == 0 {
		return []ruleValue{}
	}
	strArr := strings.Split(str, RuleSeparator)
	ruleValueSlice := make([]ruleValue, 0, len(strArr))
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
	return ruleValueSlice
}
