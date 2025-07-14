// 练习5.16： 编写多参数版本的strings.Join。
package main

import "strings"

func Join(sep string, elements ...string) string {
	return strings.Join(elements, sep)
}
