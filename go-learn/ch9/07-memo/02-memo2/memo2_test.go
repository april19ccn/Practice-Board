// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo2_test

import (
	memo2 "example/learn/ch9/07-memo/02-memo2"
	"example/learn/ch9/07-memo/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo2.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
	m := memo2.New(httpGetBody)
	memotest.Concurrent(t, m)
}

//  go test -run=TestConcurrent -race -v .\memo2_test.go

// === RUN   TestConcurrent
// http://gopl.io, 1.2051831s, 4154 bytes
// http://gopl.io, 1.2051831s, 4154 bytes
// http://gopl.io, 1.2051831s, 4154 bytes
// http://gopl.io, 1.2051831s, 4154 bytes
// https://godoc.org, 2.013595s, 33470 bytes
// https://golang.org, 2.8170698s, 62967 bytes
// http://gopl.io, 2.8170698s, 4154 bytes
// https://play.golang.org, 3.9967995s, 30614 bytes
// https://golang.org, 3.9962946s, 62967 bytes
// https://play.golang.org, 3.9962946s, 30614 bytes
// http://gopl.io, 3.9962946s, 4154 bytes
// https://godoc.org, 3.9962946s, 33470 bytes
// --- PASS: TestConcurrent (4.00s)
// PASS
// ok      command-line-arguments  6.684s
