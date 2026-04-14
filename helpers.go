package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	warningChars = []string{"'", "--", "&"}
	escapedChars = []string{"\\'", "", ""}
)

func sanitize(input interface{}, t string) interface{} {
	m := input.(string)
	switch t {
	case "bool":
		feetFloat, _ := strconv.ParseBool(strings.TrimSpace(m))
		return feetFloat
	case "float":
		feetFloat, _ := strconv.ParseFloat(strings.TrimSpace(m), 64)
		return feetFloat
	case "string":
		outString := m
		for i := 0; i < len(warningChars); i++ {
			outString = strings.Replace(outString, warningChars[i], escapedChars[i], -1)
		}
		return outString
	case "int64":
		intVal, _ := strconv.ParseInt(strings.TrimSpace(m), 10, 64)
		return intVal
	case "int":
		intVal, _ := strconv.Atoi(strings.TrimSpace(m))
		return intVal
	default:
		panic(fmt.Sprintf("unexpected type: %T", m))
	}
}
