package utils

import "github.com/bytedance/sonic"

func MarshalString(value interface{}) string {
	marshal, err := sonic.Marshal(value)
	if err != nil {
		return ""
	}
	return string(marshal)
}
