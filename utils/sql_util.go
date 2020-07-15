package utils

import (
	"bytes"
)

func GetInCondition(args []int64) string {
	len := len(args)
	var buffer bytes.Buffer
	for i := range args {
		if i == 0 {
			buffer.WriteString("(")
		}

		buffer.WriteString("?")

		if i == len-1 {
			buffer.WriteString(")")
		} else {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func GetInParams(args []interface{}) string {
	len := len(args)
	var buffer bytes.Buffer
	for i := range args {
		if i == 0 {
			buffer.WriteString("(")
		}

		buffer.WriteString("?")

		if i == len-1 {
			buffer.WriteString(")")
		} else {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}
