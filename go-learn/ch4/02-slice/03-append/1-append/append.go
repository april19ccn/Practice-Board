package main

func main() {
	var runes []rune
	for _, r := range "hello, 世界" {
		runes = append(runes, r)
	}
	println(string(runes))
}
