package path

import "github.com/natevvv/osm-ship-routing/pkg/graph"

func FindShortestPath(g graph.Graph, origin, destination int) ([]int, int) {
	navigator := GetNavigator(g)
	length := navigator.ComputeShortestPath(origin, destination)
	path := navigator.GetPath(origin, destination)
	return path, length
}

func GetNavigator(g graph.Graph) Navigator {
	//return NewBidirectionalDijkstra(g)
	//return NewAStar(g)
	//return NewDijkstra(g)
	//return NewUniversalDijkstra(g) // Dijkstra
	// astar := NewUniversalDijkstra(g)
	// astar.SetUseHeuristic(true)
	// return astar
	bidijkstra := NewUniversalDijkstra(g)
	bidijkstra.SetBidirectional(true)
	return bidijkstra
}