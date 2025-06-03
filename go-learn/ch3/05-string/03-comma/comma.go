package main

// “12345”处理后成为“12,345”
func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	return comma(s[:len(s)-3]) + "," + s[len(s)-3:]
}
