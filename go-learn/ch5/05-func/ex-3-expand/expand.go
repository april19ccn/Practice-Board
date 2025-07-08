package main

import "strings"

func expand(s string, f func(string) string) string {
	sArr := strings.Split(s, "foo")

	result := ""
	for i, v := range sArr {
		if i == len(sArr)-1 {
			result += v
		} else {
			result += v + f("foo")
		}
	}

	return result
}

func expandReplace(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "foo", f("foo"))
}
