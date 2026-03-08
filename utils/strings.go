package utils

import (
	"fmt"
	"github.com/bytedance/sonic"
)

func MarshalString(value interface{}) string {
	marshal, err := sonic.Marshal(value)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func ValidateStrLen(text string, maxLen int) error {
	if len(text) == 0 {
		return fmt.Errorf("field can not be empty")
	}

	if len(text) > maxLen {
		return fmt.Errorf("field can not be longer than %d", maxLen)
	}

	return nil
}
