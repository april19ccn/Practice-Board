package main

import (
	"bufio"
	"example/learn/ch2/06-package/ex-2/converter"
	"fmt"
	"os"
	"strconv"
)

func GToMGAndKG(g converter.Gram) {
	mg := converter.GToMG(converter.Gram(g))
	kg := converter.GToKG(converter.Gram(g))
	fmt.Printf("%s = %s, %s = %s\n", g, mg, g, kg)
}

func main() {
	if len(os.Args) != 0 {
		for _, arg := range os.Args[1:] {
			g, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}

			GToMGAndKG(converter.Gram(g))
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() { // 对于 test.txt 也是一行一行读取
			g, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}

			GToMGAndKG(converter.Gram(g))
		}
	}
}
