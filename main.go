package main

import (
	"fmt"
	"time"
)

// Node ...
type Node struct {
	visited    bool
	value      string
	neighbours []*Node
	depth      int
	parent     *Node
}

// NewNode ...
func NewNode(value string) *Node {
	return &Node{
		value:      value,
		neighbours: nil,
	}
}

func (n *Node) connect(node ...*Node) {
	n.neighbours = append(n.neighbours, node...)
	n.depth = n.depth + len(n.neighbours)
	for _, v := range n.neighbours {
		v.parent = n
	}
}

// Graph ...
type Graph struct{}

var visited map[*Node]bool

// Recursive DFS
func (g *Graph) recusrsiveDFS(node *Node) {
	if node.visited {
		return
	}
	node.visited = true
	fmt.Println(node.value, node.depth)
	for _, v := range node.neighbours {
		g.recusrsiveDFS(v)
	}
}

// Stack DFS
func (g *Graph) traverse(root *Node) {
	stack := NewStack()
	visited = make(map[*Node]bool, 0)
	stack.push(root)
	for stack.len() > 0 {
		temp, _ := stack.pop()
		x := temp.(*Node)
		if _, ok := visited[x]; !ok {
			visited[x] = true
			fmt.Println(x.value, x.depth)
			if x.neighbours == nil {
				fmt.Println("Path to parent: ")
				getPath(x)
			}
			for _, v := range x.neighbours {
				if _, ok := visited[v]; !ok {
					stack.push(v)
				}
			}
		}

	}
}

// Print path by backtracking
func getPath(x *Node) {
	// Parent is nil for super parent or an orphan
	if x.parent == nil {
		fmt.Print(x.value)
		fmt.Println("")
		return
	}
	fmt.Print(x.value, "->")
	getPath(x.parent)

}
func main() {
	v1 := NewNode("A")
	v1.parent = nil
	v2 := NewNode("B")
	v3 := NewNode("C")
	v4 := NewNode("D")
	v5 := NewNode("E")
	v6 := NewNode("F")
	v7 := NewNode("G")
	v8 := NewNode("H")
	v9 := NewNode("I")
	v10 := NewNode("J")
	v11 := NewNode("K")
	v12 := NewNode("L")

	g := Graph{}
	v1.connect(v2)
	v1.connect(v3)
	v1.connect(v4)
	v2.connect(v5)
	v2.connect(v6)
	v6.connect(v7)
	v6.connect(v8)
	v8.connect(v9)
	v5.connect(v10)
	v5.connect(v11)
	v8.connect(v12)
	fmt.Println("Recursive DFS =================")
	g.recusrsiveDFS(v1)
	fmt.Println("DFS using Stack implementation =================")
	start := time.Now()
	g.traverse(v1)
	fmt.Println(time.Since(start))
}
