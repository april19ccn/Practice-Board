// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo3_test

import (
	memo3 "example/learn/ch9/07-memo/03-memo3"
	"example/learn/ch9/07-memo/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo3.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
	m := memo3.New(httpGetBody)
	memotest.Concurrent(t, m)
}

//  go test -run=TestConcurrent -race -v .\memo3_test.go

// === RUN   TestConcurrent
// https://golang.org, 887.6744ms, 62967 bytes
// https://golang.org, 914.8352ms, 62967 bytes
// https://godoc.org, 916.9106ms, 33470 bytes
// https://godoc.org, 946.36ms, 33470 bytes
// https://play.golang.org, 1.3242273s, 30614 bytes
// https://play.golang.org, 1.3348749s, 30614 bytes
// http://gopl.io, 1.4807096s, 4154 bytes
// http://gopl.io, 1.4817487s, 4154 bytes
// http://gopl.io, 1.4817487s, 4154 bytes
// http://gopl.io, 1.4823694s, 4154 bytes
// http://gopl.io, 1.4834776s, 4154 bytes
// http://gopl.io, 1.4829643s, 4154 bytes
// --- PASS: TestConcurrent (1.48s)
// PASS
// ok      command-line-arguments  4.132s
