package path

import (
	"fmt"
	"math"
	"testing"

	"github.com/natevvv/osm-ship-routing/pkg/graph"
)

const graphFmi = `10
26
# nodes
0 0 0
1 0 1
2 0 2
3 1 0
4 1 1
5 1 2
6 2 0
7 2 1
8 2 2
9 3 3
# edges
0 1 1
0 3 1
1 0 1
1 2 1
1 4 1
2 1 1
2 5 1
3 1 1
3 4 1
3 6 1
4 1 1
4 3 1
4 5 1
4 7 1
5 2 1
5 4 1
5 8 1
6 3 1
6 7 1
7 4 1
7 6 1
7 8 1
8 5 1
8 7 1
8 9 1
9 8 1`

func TestPlainDijkstra(t *testing.T) {
	aag := graph.NewAdjacencyArrayFromFmiString(graphFmi)
	d := NewUniversalDijkstra(aag, false)
	path, length := d.GetPath(0, 9)
	fmt.Printf("length: %v\n", length)
	fmt.Printf("Path: %v\n", path)
	pathReference := []int{0, 1, 4, 5, 8, 9}
	lengthReference := 5
	if length != lengthReference {
		t.Errorf("length is %v. Should be %v\n", length, lengthReference)
	}
	if len(path) != len(pathReference) {
		t.Errorf("path has wrong length. Is %v, should be %v\n", len(path), len(pathReference))
	}
	for i, v := range pathReference {
		if path[i] != v {
			t.Errorf("path at position %v has wrong value. Is %v, should be %v\n", i, path[i], v)
		}
	}

}

func TestAStarDijkstra(t *testing.T) {
	aag := graph.NewAdjacencyArrayFromFmiString(graphFmi)
	d := NewUniversalDijkstra(aag, false)
	astar := NewUniversalDijkstra(aag, true)
	path, length := d.GetPath(0, 9)
	astarPath, astarLength := astar.GetPath(0, 9)
	fmt.Printf("astar path %v\n", astarPath)
	if length != astarLength {
		t.Errorf("Length does not match. Is %v, should be %v", astarLength, length)
	}
	if len(astarPath) != len(path) {
		t.Errorf("Path has wrong length. Is %v, should be %v", len(astarPath), len(path))
	}
	for i := 0; i < int(math.Min(float64(len(path)), float64(len(astarPath)))); i++ {
		if path[i] != astarPath[i] {
			//t.Errorf("Path does not match. At pos %v, it is %v, should be %v", i, astarPath[i], path[i])
		}
	}
}

func TestBidirecitonalDijkstra(t *testing.T) {
	aag := graph.NewAdjacencyArrayFromFmiString(graphFmi)
	d := NewUniversalDijkstra(aag, false)
	bidijkstra := NewUniversalDijkstra(aag, false)
	bidijkstra.SetBidirectional(true)
	path, length := d.GetPath(0, 9)
	astarPath, astarLength := bidijkstra.GetPath(0, 9)
	fmt.Printf("astar path %v\n", astarPath)
	if length != astarLength {
		t.Errorf("Length does not match. Is %v, should be %v", astarLength, length)
	}
	if len(astarPath) != len(path) {
		t.Errorf("Path has wrong length. Is %v, should be %v", len(astarPath), len(path))
	}
	for i := 0; i < int(math.Min(float64(len(path)), float64(len(astarPath)))); i++ {
		if path[i] != astarPath[i] {
			//t.Errorf("Path does not match. At pos %v, it is %v, should be %v", i, astarPath[i], path[i])
		}
	}
}
