// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo5_test

import (
	memo5 "example/learn/ch9/07-memo/05-memo5"
	"example/learn/ch9/07-memo/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo5.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
	m := memo5.New(httpGetBody)
	memotest.Concurrent(t, m)
}

//  go test -run=TestConcurrent -race -v .\memo4_test.go
