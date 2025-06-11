package main

import (
	"fmt"
	"strconv"
)

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]

	fmt.Println("Q2 length: " + strconv.Itoa(len(Q2)) + " summer length: " + strconv.Itoa(len(summer)))
	fmt.Println("Q2 cap: " + strconv.Itoa(cap(Q2)) + " summer cap: " + strconv.Itoa(cap(summer)))

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// 超过cap会引起panic
	// fmt.Println(summer[:20]) // panic: out of range // anic: runtime error: slice bounds out of range [:20] with capacity 7

	// 超过len，但小于cap会扩展slice
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"

	// 复制一个slice只是对底层的数组创建了一个新的slice别名
	fmt.Println("-------------------")
	Q2[2] = "Apple"
	fmt.Println(Q2)     // [April May Apple]
	fmt.Println(summer) // [Apple July August]
	fmt.Println(months) // [January February March April May Apple July August September October November December]
}
