// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo4_test

import (
	memo4 "example/learn/ch9/07-memo/04-memo4"
	"example/learn/ch9/07-memo/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo4.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
	m := memo4.New(httpGetBody)
	memotest.Concurrent(t, m)
}

//  go test -run=TestConcurrent -race -v .\memo4_test.go
