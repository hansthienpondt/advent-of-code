package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
)

type Node struct {
	id int64
	r  rune
}

func (n Node) ID() int64      { return n.id }
func (n Node) String() string { return string(n.r) }

type MyGraph struct {
	Graph   *simple.DirectedGraph
	mynodes map[int64]Node
}

func NewMyGraph() *MyGraph {
	return &MyGraph{
		Graph:   simple.NewDirectedGraph(),
		mynodes: make(map[int64]Node)}
}
func NewNode(id int64, r rune) Node {
	return Node{id: id, r: r}
}
func (g *MyGraph) AddNode(n Node) {
	g.Graph.AddNode(n)
	g.mynodes[n.ID()] = n
}
func (g *MyGraph) GetNode(n int64) Node {
	return g.mynodes[n]
}

func readInput(file string) []string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := string(bytes)
	lines := strings.Split(data, "\n")

	return lines
}

func day12a(input []string) int {
	grid := make([][]Node, 0)

	g := NewMyGraph()
	var start, end int64 = 0, 0
	// rune bounds, a = 97 ; z = 122
	// S = 83 ; E = 69

	var nodeCounter int = 0
	for _, l := range input {
		gr := make([]Node, 0)
		for _, c := range l {
			nodeCounter += 1
			node := NewNode(int64(nodeCounter), c)
			g.AddNode(node)
			gr = append(gr, node)
			if string(c) == "S" {
				start = int64(node.ID())
			}
			if string(c) == "E" {
				end = int64(node.ID())
			}
		}
		grid = append(grid, gr)
	}
	/*
		fmt.Println(start, end)

		for _, i := range grid {
			for _, j := range i {
				fmt.Printf("%v ", j.ID())
			}
			fmt.Println()
		}
	*/
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0])-1; j++ {
			if IsNeighbor(grid[i][j], grid[i][j+1]) {
				g.Graph.SetEdge(simple.Edge{F: grid[i][j], T: grid[i][j+1]})
			}
			if IsNeighbor(grid[i][j+1], grid[i][j]) {
				g.Graph.SetEdge(simple.Edge{F: grid[i][j+1], T: grid[i][j]})
			}
			//tmp
			if IsDescent(grid[i][j], grid[i][j+1]) {
				g.Graph.SetEdge(simple.Edge{F: grid[i][j], T: grid[i][j+1]})
			}
			//fmt.Printf("(%v %v) ", grid[i][j], grid[i][j+1])
		}
		//fmt.Println()
	}
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid)-1; j++ {
			if IsNeighbor(grid[j][i], grid[j+1][i]) {
				g.Graph.SetEdge(simple.Edge{F: grid[j][i], T: grid[j+1][i]})
			}
			if IsNeighbor(grid[j+1][i], grid[j][i]) {
				g.Graph.SetEdge(simple.Edge{F: grid[j+1][i], T: grid[j][i]})
			}
			//tmp
			if IsDescent(grid[j][i], grid[j+1][i]) {
				g.Graph.SetEdge(simple.Edge{F: grid[j][i], T: grid[j+1][i]})
			}
			//fmt.Printf("(%v %v) ", grid[j][i], grid[j+1][i])
		}
		//fmt.Println()
	}

	p := path.DijkstraFrom(g.GetNode(start), g.Graph)
	_, count := p.To(end)
	return int(count)
}

func day12b(input []string) int {
	grid := make([][]Node, 0)

	g := NewMyGraph()

	startlist := make([]int64, 0)

	var start, end int64 = 0, 0
	// rune bounds, a = 97 ; z = 122
	// S = 83 ; E = 69

	var nodeCounter int = 0
	for _, l := range input {
		gr := make([]Node, 0)
		for _, c := range l {
			nodeCounter += 1
			node := NewNode(int64(nodeCounter), c)
			g.AddNode(node)
			gr = append(gr, node)
			if string(c) == "S" {
				start = int64(node.ID())
			}
			if string(c) == "E" {
				end = int64(node.ID())
			}
			if string(c) == "a" {
				startlist = append(startlist, int64(node.ID()))
			}
		}
		grid = append(grid, gr)
	}
	/*
		fmt.Println(start, end)

		for _, i := range grid {
			for _, j := range i {
				fmt.Printf("%v ", j.ID())
			}
			fmt.Println()
		}
	*/
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0])-1; j++ {
			if IsNeighbor(grid[i][j], grid[i][j+1]) {
				g.Graph.SetEdge(simple.Edge{F: grid[i][j], T: grid[i][j+1]})
			}
			if IsNeighbor(grid[i][j+1], grid[i][j]) {
				g.Graph.SetEdge(simple.Edge{F: grid[i][j+1], T: grid[i][j]})
			}
			//tmp
			if IsDescent(grid[i][j], grid[i][j+1]) {
				g.Graph.SetEdge(simple.Edge{F: grid[i][j], T: grid[i][j+1]})
			}
			//fmt.Printf("(%v %v) ", grid[i][j], grid[i][j+1])
		}
		//fmt.Println()
	}
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid)-1; j++ {
			if IsNeighbor(grid[j][i], grid[j+1][i]) {
				g.Graph.SetEdge(simple.Edge{F: grid[j][i], T: grid[j+1][i]})
			}
			if IsNeighbor(grid[j+1][i], grid[j][i]) {
				g.Graph.SetEdge(simple.Edge{F: grid[j+1][i], T: grid[j][i]})
			}
			//tmp
			if IsDescent(grid[j][i], grid[j+1][i]) {
				g.Graph.SetEdge(simple.Edge{F: grid[j][i], T: grid[j+1][i]})
			}
			//fmt.Printf("(%v %v) ", grid[j][i], grid[j+1][i])
		}
		//fmt.Println()
	}
	var count int = 0

	p := path.DijkstraFrom(g.GetNode(start), g.Graph)
	_, c := p.To(end)

	count = int(c)

	for _, item := range startlist {
		p := path.DijkstraFrom(g.GetNode(item), g.Graph)
		_, c := p.To(end)
		if math.IsInf(c, 1) {
			continue
		}
		if int(c) < count {
			count = int(c)
		}
	}
	return count
}
func newSimpleNode(id int) graph.Node {
	return simple.Node(id)
}

func IsNeighbor(i, j Node) bool {
	if i.r == 97 && j.r == 83 {
		return true
	}
	if i.r == 83 && j.r == 97 {
		return true
	}
	if i.r == 69 && j.r == 122 {
		return true
	}
	if i.r == 122 && j.r == 69 {
		return true
	}
	if i.r == j.r {
		return true
	}
	if i.r == j.r-1 {
		return true
	}
	if i.r == j.r+1 {
		return true
	}
	return false
}
func IsDescent(i, j Node) bool {
	if i.r == 83 || j.r == 83 {
		return false
	}
	if i.r == 69 || j.r == 69 {
		return false
	}
	if j.r < i.r {
		return true
	}
	return false
}
func main() {
	filename := os.Args[1]

	fmt.Printf("Part 1: %d\n", day12a(readInput(filename)))
	fmt.Printf("Part 2: %d\n", day12b(readInput(filename)))
}
