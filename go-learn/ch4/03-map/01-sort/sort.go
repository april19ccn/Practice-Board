package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"charlie": 34,
		"emily":   28,
		"alice":   31,
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	fmt.Println(names)
	sort.Strings(names)
	fmt.Println(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
