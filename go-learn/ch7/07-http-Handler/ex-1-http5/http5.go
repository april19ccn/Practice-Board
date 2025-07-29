// 练习 7.11： 增加额外的handler让客户端可以创建，读取，更新和删除数据库记录。
// 例如，一个形如 /update?item=socks&price=6 的请求会更新库存清单里一个货品的价格并且当这个货品不存在或价格无效时返回一个错误值。
// （注意：这个修改会引入变量同时更新的问题）
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	mu.Lock()
	defer mu.Unlock()

	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

// 练习 7.11：创建，读取，更新和删除数据库记录
func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	data, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusConflict) // 409
		fmt.Fprintf(w, "item already exists: %q\n", item)
		return
	}

	db[item] = dollars(data)
	fmt.Fprintf(w, "create success! %s: %s\n", item, dollars(data))
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

// 这个写法里 涉及了死锁问题
// func (db database) update(w http.ResponseWriter, req *http.Request) {
// 	mu.Lock()
// 	item := req.URL.Query().Get("item")
// 	if _, ok := db[item]; !ok {
// 		mu.Unlock()
// 		db.create(w, req) // 共用了同一个锁，会死锁
// 		return
// 	}
// 	price := req.URL.Query().Get("price")
// 	data, err := strconv.ParseFloat(price, 32)
// 	if err != nil {
// 		w.WriteHeader(http.StatusConflict) // 409
// 		fmt.Fprintf(w, "invalid price")
// 		return
// 	}
// 	db[item] = dollars(data)
// 	mu.Unlock()
// 	fmt.Fprintf(w, "update success! %s: %s\n", item, price)
// }

func (db database) update(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	data, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}

	// 有则更新，无则创建
	db[item] = dollars(data)
	fmt.Fprintf(w, "update success! %s: %s\n", item, dollars(data))
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "delete success!")
}

func main() {
	// http4:
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// http://localhost:8000/list
// shoes: $50.00
// socks: $5.00

// http://localhost:8000/create?item=hat&price=25.5
// create success! hat: 25.5

// http://localhost:8000/list
// shoes: $50.00
// socks: $5.00
// hat: $25.50

// http://localhost:8000/read?item=hat
// hat: $25.50

// http://localhost:8000/read?item=hat1
// no such item: "hat1"

// http://localhost:8000/update?item=pants&price=200
// update success! pants: $200.00

// http://localhost:8000/update?item=hat&price=316
// update success! hat: $316.00

// http://localhost:8000/read?item=hat
// hat: $316.00

// http://localhost:8000/read?item=pants
// pants: $200.00

// http://localhost:8000/delete?item=hat
// delete success!
