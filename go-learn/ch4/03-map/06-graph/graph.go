package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("a", "b")
	addEdge("a", "c")
	addEdge("b", "d")
	addEdge("b", "e")
	addEdge("c", "f")
	addEdge("c", "g")
	addEdge("d", "h")
	addEdge("d", "i")
	addEdge("e", "j")
	addEdge("e", "k")
	addEdge("f", "l")
	addEdge("f", "m")
	addEdge("g", "n")
	addEdge("g", "o")

	fmt.Println(graph)
	fmt.Println(hasEdge("a", "e"))
}
