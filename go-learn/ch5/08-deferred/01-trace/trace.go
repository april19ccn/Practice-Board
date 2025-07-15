// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 146.

// The trace program uses defer to add entry/exit diagnostics to a function.
package main

import (
	"log"
	"time"
)

// !+main
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

//!-main

func main() {
	bigSlowOperation()
}

/*
!+output
$ go build gopl.io/ch5/trace
$ ./trace
2015/11/18 09:53:26 enter bigSlowOperation
2015/11/18 09:53:36 exit bigSlowOperation (10.000589217s)
!-output
*/

// 2025/07/15 11:07:16 enter bigSlowOperation
// 2025/07/15 11:07:26 exit bigSlowOperation (10.0237915s)

// 2025/07/15 11:07:36 enter bigSlowOperation
// 2025/07/15 11:07:46 exit bigSlowOperation (10.0232181s)
