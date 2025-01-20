package util

import "strings"

// IndexOfNum returns the index of the num-th occurrence of substr in s
func IndexOfNum(s, substr string, num int) (index int) {
	index = strings.Index(s, substr)
	// end point
	if num == 1 {
		return
	}
	// if the length of s is less than the length of substr, return -1
	if len(s) < len(substr) {
		return -1
	}
	// recursive call
	nextIndex := IndexOfNum(s[strings.Index(s, substr)+len(substr):], substr, num-1)
	// if the next index is -1, return -1
	if nextIndex == -1 {
		return -1
	}
	// add the length of substr and the next index
	index += (len(substr) + nextIndex)

	return
}
