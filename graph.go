package main

import "fmt"

type VisitedNodes struct {
	stack []string
}

func (v *VisitedNodes) Nil() bool {
	if len(v.stack) > 0 {
		return true
	}
	return false
}
func (v *VisitedNodes) Append(item string) {
	v.stack = append(v.stack, item)
}

func (v *VisitedNodes) Print() {
	for _, item := range v.stack {
		fmt.Println(item)
	}
}

func (v *VisitedNodes) HasAny(item string) bool {
	for _, ite := range v.stack {
		if item == ite {
			return true
		}
	}
	return false
}

type Queue struct {
	stack []string
}

func (q *Queue) Nil() bool {
	if len(q.stack) > 0 {
		return true
	}
	return false
}
func (q *Queue) Append(item string) {
	q.stack = append(q.stack, item)
}

func (q *Queue) Print() {
	for _, item := range q.stack {
		fmt.Println(item)
	}
}

func (q *Queue) Dequeue() (string, bool) {
	if len(q.stack) == 0 {
		return "", false
	} else if len(q.stack) > 1 {
		item := q.stack[0]
		q.stack = q.stack[1:len(q.stack)]
		return item, true
	}
	item := q.stack[0]
	q.stack = q.stack[:0]
	return item, true
}

func NewQueue() *Queue {
	return &Queue{}
}

func NewVisitedNodes() *VisitedNodes {
	return &VisitedNodes{}
}

func BFS(graph map[string][]string, root string) {
	q := NewQueue()
	visited := NewVisitedNodes()
	q.Append(root)

	for q.Nil() {
		node, _ := q.Dequeue()

		if !visited.HasAny(node) {
			visited.Append(node)

		}
		for _, item := range graph[node] {
			if !visited.HasAny(item) {
				q.Append(item)
			}

		}

	}
	visited.Print()

}

func main() {
	graph := make(map[string][]string)
	graph["a"] = []string{"b", "c"}
	graph["b"] = []string{"a", "d"}
	graph["c"] = []string{"a", "d"}
	graph["d"] = []string{"b", "c", "e"}
	graph["e"] = []string{"d"}
	BFS(graph, "e")

}
