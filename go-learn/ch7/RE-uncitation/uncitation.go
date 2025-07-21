package main

func test(x int) int {
	return x + 1
}

type C struct {
	x int
}

func test1(x C) int {
	return x.x + 1
}

func main() {
	// x := 1
	// p := &x
	// fmt.Println(test(p)) // cannot use p (variable of type *int) as int value in argument to test

	// c := C{1}
	// fmt.Println(test1(&c)) // cannot use &c (value of type *C) as C value in argument to test1
}
