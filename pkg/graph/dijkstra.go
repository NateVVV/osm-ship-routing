package graph

import (
	"container/heap"
)

type Dijkstra struct{}

func (d Dijkstra) ShortestPath(g Graph, origin, destination int) ([]int, int) {
	dijkstraItems := make([]*PriorityQueueItem, g.NodeCount(), g.NodeCount())
	originItem := PriorityQueueItem{itemId: origin, priority: 0, predecessor: -1, index: -1}
	dijkstraItems[origin] = &originItem

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, dijkstraItems[origin])

	for len(pq) > 0 {
		currentPqItem := heap.Pop(&pq).(*PriorityQueueItem)
		currentNodeId := currentPqItem.itemId

		for _, edge := range g.GetEdgesFrom(currentNodeId) {
			successor := edge.To

			if dijkstraItems[successor] == nil {
				newPriority := dijkstraItems[currentNodeId].priority + edge.Distance
				pqItem := PriorityQueueItem{itemId: successor, priority: newPriority, predecessor: currentNodeId, index: -1}
				dijkstraItems[successor] = &pqItem
				heap.Push(&pq, &pqItem)
			} else {
				if updatedDistance := dijkstraItems[currentNodeId].priority + edge.Distance; updatedDistance < dijkstraItems[successor].priority {
					pq.update(dijkstraItems[successor], updatedDistance)
					dijkstraItems[successor].predecessor = currentNodeId
				}
			}
		}

		if currentNodeId == destination {
			break
		}
	}

	length := -1           // by default a non-existing path has length -1
	path := make([]int, 0) // by default, a non-existing path is an empty slice
	if dijkstraItems[destination] != nil {
		length = dijkstraItems[destination].priority
		for nodeId := destination; nodeId != -1; nodeId = dijkstraItems[nodeId].predecessor {
			path = append([]int{nodeId}, path...)
		}
	}
	return path, length
}
