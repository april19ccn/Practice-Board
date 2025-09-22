// 仅理解（匿名导入的应用原理）通过init注册到全局变量的原理
package core

import "sync"

type decodeFunc = func(path string)

var registerList = make(map[string]decodeFunc)

var lock sync.Mutex

func Register(name string, decode decodeFunc) {
	lock.Lock()
	defer lock.Unlock()
	registerList[name] = decode
}

func Decode(format string, path string) {
	lock.Lock()
	defer lock.Unlock()
	registerList[format](path)
}
