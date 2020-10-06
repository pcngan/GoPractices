package cyclegraph

import (
	"fmt"
)

type CGraph struct{
	First, Count, N int
	VisitedVertices []bool
	VisitedEdges [][]bool
	Edges [][]bool
}

func Solution(A, B []int) bool {
	maxVertex := getMax(A, B)

	vertices := make([]bool, maxVertex+1)
	
	graph := initGraph(maxVertex)

	for i := 0; i < len(A); i++ {
		vertices[A[i]], vertices[B[i]] = true, true
		graph.Edges[A[i]][B[i]] = true
	}

	for i, v := range vertices {
		if v {
			graph.N++
			if graph.First == -1 {
				graph.First = i
			}
		}
	}

	if graph.First != -1 {
		graph.VisitedVertices[graph.First] = true
		graph.Count++
		return dfs(graph.First, &graph)
	}
	
	fmt.Println(graph.Edges)
	return false
}

func initGraph(numVertices int) CGraph{
	edges := make([][]bool, numVertices+1)
	for i := 0; i <= numVertices; i++ {
		edges[i] = make([]bool, numVertices+1)
	}

	visitedVertices := make([]bool, numVertices+1)
	visitedEdges := make([][]bool, numVertices+1)
	for i := 0; i <= numVertices; i++ {
		visitedEdges[i] = make([]bool, numVertices+1)
	}
	
	return CGraph{
			First: -1,
			Count: 0,
			N : 0,
			VisitedVertices : visitedVertices,
			VisitedEdges : visitedEdges,
			Edges : edges,
		}
}

func getMax(A, B []int) int {
	max := A[0]
	for _, v := range A {
		if v > max {
			max = v
		}
	}

	for _, v := range B {
		if v > max {
			max = v
		}
	}

	return max
}

func dfs(i int, graph *CGraph) bool {
	for j := 0; j < len(graph.Edges[i]); j++ {
		if graph.Edges[i][j] && !graph.VisitedEdges[i][j] {
			if j == graph.First && graph.Count == graph.N {
				return true
			}

			if !graph.VisitedVertices[j] {
				graph.VisitedVertices[j] = true
				graph.Count++
			}
			graph.VisitedEdges[i][j] = true
			if dfs(j, graph) {
				return true
			}
			graph.VisitedEdges[i][j] = false
		}
	}
	graph.VisitedVertices[i] = false
	graph.Count--
	return false
}

/*
var First int = -1
	var Count int = 0
	var N = 0
	var VisitedVertices []bool
	var VisitedEdges [][]bool
	var Edges [][]bool

- get max vertex num -> V
- list vertices[V+1], 
- matrix edges[V+1][V+1]
- edge(a,b) -> set vertices[i] =1, vertices[b]=1, count vertices=1 -> N
			-> edges[a][b] = 1
			
- list visitedVertices[V+1]
- matrix visitedEdges[V+1][V+1]


- get first vertex First = i, set visitedVertices[i]=1, Count++
dfs(i, )
- for Edges[i][j=0-V]{
		if (Edges[i][j] == 1 && VisitedEdges[i][j]==0){
			if(j==First && Count==N)
				return true
			if( VisitedVertices[j] == 0){
				VisitedVertices[j] = 1
				Count++
			}
			VisitedEdges[i][[j] = 1
			if(dfs(j)) return true
			VisitedEdges[i][[j] = 0
		}
	}
	
	VisitedVertices[j] = 0
	Count--
	return false

*/