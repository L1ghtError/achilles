package strutil

import (
	"fmt"
	"strings"
)

func ExtractNumber(p string) float32 {
	num := ""
	isStarted := false
	isDotted := false
	var res float32 = 0.0
	for _, v := range p {
		if !isDotted && isStarted && v == '.' {
			num += "."
			continue
		}
		if v >= '0' && v <= '9' {
			num += string(v)
			isStarted = true
			continue
		} else if isStarted {
			break
		}

	}
	fmt.Sscanf(num, "%f", &res)
	return res
}

func FindAllIndexes(str, substr string) []int {
	var indexes []int
	start := 0

	for {
		index := strings.Index(str[start:], substr)
		if index == -1 {
			break
		}
		indexes = append(indexes, start+index)
		start += index + len(substr)
	}

	return indexes
}
