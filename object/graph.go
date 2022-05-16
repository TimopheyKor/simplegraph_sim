package object

import (
	"fmt"
	"strings"
)

type Graph struct {
	Verts map[Vertex][]Vertex
	Edges []Edge
}

// NewGraph returns a new empty Graph.
func NewGraph() Graph {
	return Graph{Verts: make(map[Vertex][]Vertex)}
}

// AddVertex takes a Vertex and adds it as a key to the Graph's map of vertices.
func (g *Graph) AddVertex(v Vertex) {
	_, ok := g.Verts[v]
	if ok {
		fmt.Printf("Error: Vertex %v already exists.", v)
		return
	}
	g.Verts[v] = []Vertex{}
}

// AddEdge takes two Vertex objects, and appends each Vertex as a value to the
// other Vertex's key in the Graph's map. It then creates an Edge object
func (g *Graph) AddEdge(start, finish Vertex) {
	_, one := g.Verts[start]
	_, two := g.Verts[finish]
	if one && two {
		g.Verts[start] = append(g.Verts[start], finish)
		g.Verts[finish] = append(g.Verts[finish], start)
		g.Edges = append(g.Edges, NewEdge(start, finish))
		return
	}
	fmt.Printf("Error: vertex %v or %v not in map.", start, finish)
}

// String returns a map of the Graph's data in string format.
func (g Graph) String() string {
	return fmt.Sprint(g.Verts)
}

// PrintGraph prints the Graph in the format of a list of every edge between
// two vertices, and every disconnected vertex.
func (g Graph) PrintGraph() {
	var result string
	for key := range g.Verts {
		if len(g.Verts[key]) == 0 {
			result = fmt.Sprintf("%s\n%v", result, key)
		}
		for _, value := range g.Verts[key] {
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

// RemoveVertex takes a Vertex object and removes it from the graph.
func (g *Graph) RemoveVertex(v Vertex) {
	_, ok := g.Verts[v]
	if ok {
		delete(g.Verts, v)
		g.RemoveConnectedEdges(v)
	}
}

// TODO: Create better solution to remove single edges.
// RemoveConnectedEdges loops through the Graph's list of edges and removes
// any without vertices.
func (g *Graph) RemoveConnectedEdges(v Vertex) {
	for i, edge := range g.Edges {
		if edge.start == v || edge.finish == v {
			defer g.removeEdgeAtIndex(i)
		}
	}
}

// removeEdgeAtIndex takes an int i and removes the edge at index i of
// the Graph's slice of edges.
func (g *Graph) removeEdgeAtIndex(i int) {
	g.Edges = append(g.Edges[:i], g.Edges[i+1:]...)
}
