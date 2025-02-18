package strutil

import (
	"fmt"
	"math"
	"strconv"
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

func JoinInts[T int | int32 | uint32 | uint64](elems []T, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return strconv.Itoa(int(elems[0]))
	}
	var n int
	if len(sep) > 0 {
		if len(sep) >= math.MaxInt64/(len(elems)-1) {
			panic("strings: Join output length overflow")
		}
		n += len(sep) * (len(elems) - 1)
	}
	for _, elem := range elems {
		n += int(math.Log10(float64(elem)) + 1)
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.Itoa(int(elems[0])))
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(strconv.Itoa(int(s)))
	}
	return b.String()
}
