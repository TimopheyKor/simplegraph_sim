package object

import (
	"fmt"
	"strings"
)

type Graph struct {
	Data map[Vertex][]Vertex
}

// NewGraph returns a new empty Graph.
func NewGraph() Graph {
	return Graph{Data: make(map[Vertex][]Vertex)}
}

// AddVertex takes a Vertex and adds it as a key to the Graph's map of vertices.
func (g *Graph) AddVertex(v Vertex) {
	_, ok := g.Data[v]
	if ok {
		fmt.Printf("Error: Vertex %v already exists.", v)
		return
	}
	g.Data[v] = []Vertex{}
}

// AddEdge takes two Vertex objects, and appends each Vertex as a value to the
// other Vertex's key in the Graph's map.
func (g *Graph) AddEdge(start, end Vertex) {
	_, one := g.Data[start]
	_, two := g.Data[end]
	if one && two {
		g.Data[start] = append(g.Data[start], end)
		g.Data[end] = append(g.Data[end], start)
		return
	}
	fmt.Printf("Error: vertex %v or %v not in map.", start, end)
}

// String returns a map of the Graph's data in string format.
func (g Graph) String() string {
	return fmt.Sprint(g.Data)
}

// PrintGraph prints the Graph in the format of a list of every edge between
// two vertices, and every disconnected vertex.
func (g Graph) PrintGraph() {
	var result string
	for key := range g.Data {
		if len(g.Data[key]) == 0 {
			result = fmt.Sprintf("%s\n%v", result, key)
		}
		for _, value := range g.Data[key] {
			eOne := fmt.Sprintf("%v<->%v", key, value)
			eTwo := fmt.Sprintf("%v<->%v", value, key)
			if strings.Contains(result, eOne) || strings.Contains(result, eTwo) {
				continue
			}
			result = fmt.Sprintf("%s\n%s", result, eOne)
		}
	}
	fmt.Println(result)
}
