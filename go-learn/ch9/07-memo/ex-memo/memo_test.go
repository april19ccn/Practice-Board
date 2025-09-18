// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	memo "example/learn/ch9/07-memo/ex-memo"
	"testing"
)

var httpGetBody = memo.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memo.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memo.Concurrent(t, m)
}

func TestCancel(t *testing.T) {
	m := memo.New(httpGetBody)
	memo.CancelRequest(t, m)
}

//  go test -run=TestCancel -race -v .\memo_test.go
