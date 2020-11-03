package main

import (
	"strconv"
	"strings"
)

func Reverse(s string) string {
	r := ""

	for _, ch := range s {
		r = string(ch) + r
	}

	return r
}

func BinarySearch(arr []int, b int) int {
	left := 0
	right := len(arr) - 1

	if len(arr) == 0 {
		return -1
	}

	index := 0
	mustIndex := -1

	for {
		if mustIndex != -1 {
			index = mustIndex
		} else {
			index = ((right - left) / 2) + left
		}

		switch {
		case b == arr[index]:
			return index
		case b > arr[index]:
			left = index
		case arr[index] > b:
			right = index
		}

		switch {
		case mustIndex != -1:
			return -1
		case right-left == 0:
			return -1
		case right-left == 1:
			mustIndex = right
		}
	}
}

func IntsToString(arr []int) string {
	sb := strings.Builder{}
	for _, v := range arr {
		if sb.Len() > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(strconv.Itoa(v))
	}

	return sb.String()
}
