package path

import (
	"container/heap"
	"math"

	"github.com/natevvv/osm-ship-routing/pkg/graph"
	"github.com/natevvv/osm-ship-routing/pkg/queue"
)

type BidirectionalDijkstra struct{}

func (bd BidirectionalDijkstra) ShortestPath(g graph.Graph, origin, destination int) ([]int, int) {
	forwardDijkstraItems := make([]*queue.PriorityQueueItem, g.NodeCount(), g.NodeCount())
	originItem := queue.PriorityQueueItem{ItemId: origin, Priority: 0, Predecessor: -1, Index: -1}
	forwardDijkstraItems[origin] = &originItem

	backwardDijkstraItems := make([]*queue.PriorityQueueItem, g.NodeCount(), g.NodeCount())
	destinationItem := queue.PriorityQueueItem{ItemId: destination, Priority: 0, Predecessor: -1, Index: -1}
	backwardDijkstraItems[destination] = &destinationItem

	forwardPq := make(queue.PriorityQueue, 0)
	heap.Init(&forwardPq)
	heap.Push(&forwardPq, forwardDijkstraItems[origin])

	backwardPq := make(queue.PriorityQueue, 0)
	heap.Init(&backwardPq)
	heap.Push(&backwardPq, backwardDijkstraItems[destination])

	distance := math.MaxInt
	connectionNode := 0

	for len(forwardPq) > 0 && len(backwardPq) > 0 {
		currentForwardPqItem := heap.Pop(&forwardPq).(*queue.PriorityQueueItem)
		currentForwardNodeId := currentForwardPqItem.ItemId
		currentBackwardPqItem := heap.Pop(&backwardPq).(*queue.PriorityQueueItem)
		currentBackwardNodeId := currentBackwardPqItem.ItemId

		if forwardDijkstraItems[currentForwardNodeId].Priority+backwardDijkstraItems[currentBackwardNodeId].Priority >= distance {
			break
		}

		for _, edge := range g.GetEdgesFrom(currentForwardNodeId) {
			successor := edge.To

			if forwardDijkstraItems[successor] == nil {
				newPriority := forwardDijkstraItems[currentForwardNodeId].Priority + edge.Distance
				pqItem := queue.PriorityQueueItem{ItemId: successor, Priority: newPriority, Predecessor: currentForwardNodeId, Index: -1}
				forwardDijkstraItems[successor] = &pqItem
				heap.Push(&forwardPq, &pqItem)
			} else {
				if updatedDistance := forwardDijkstraItems[currentForwardNodeId].Priority + edge.Distance; updatedDistance < forwardDijkstraItems[successor].Priority {
					forwardPq.Update(forwardDijkstraItems[successor], updatedDistance)
					forwardDijkstraItems[successor].Predecessor = currentForwardNodeId
				}
			}

			if connection := backwardDijkstraItems[successor]; connection != nil && forwardDijkstraItems[currentForwardNodeId].Priority+edge.Distance+connection.Priority < distance {
				distance = forwardDijkstraItems[currentForwardNodeId].Priority + edge.Distance + connection.Priority
				forwardDijkstraItems[successor].Predecessor = currentForwardNodeId
				connectionNode = successor
			}
		}

		if currentForwardNodeId == destination {
			break
		}

		for _, edge := range g.GetEdgesFrom(currentBackwardNodeId) {
			successor := edge.To

			if backwardDijkstraItems[successor] == nil {
				newPriority := backwardDijkstraItems[currentBackwardNodeId].Priority + edge.Distance
				pqItem := queue.PriorityQueueItem{ItemId: successor, Priority: newPriority, Predecessor: currentBackwardNodeId, Index: -1}
				backwardDijkstraItems[successor] = &pqItem
				heap.Push(&backwardPq, &pqItem)
			} else {
				if updatedDistance := backwardDijkstraItems[currentBackwardNodeId].Priority + edge.Distance; updatedDistance < backwardDijkstraItems[successor].Priority {
					backwardPq.Update(backwardDijkstraItems[successor], updatedDistance)
					backwardDijkstraItems[successor].Predecessor = currentBackwardNodeId
				}
			}

			if connection := forwardDijkstraItems[successor]; connection != nil && backwardDijkstraItems[currentBackwardNodeId].Priority+edge.Distance+connection.Priority < distance {
				distance = backwardDijkstraItems[currentBackwardNodeId].Priority + edge.Distance + connection.Priority
				backwardDijkstraItems[successor].Predecessor = currentBackwardNodeId
				connectionNode = successor
			}
		}

		if currentBackwardNodeId == origin {
			break
		}
	}

	length := -1
	path := make([]int, 0)

	if distance < math.MaxInt {
		length = distance
		for nodeId := connectionNode; nodeId != -1; nodeId = forwardDijkstraItems[nodeId].Predecessor {
			path = append([]int{nodeId}, path...)
		}
		// connectionNode may be doubled
		for nodeId := connectionNode; nodeId != -1; nodeId = backwardDijkstraItems[nodeId].Predecessor {
			path = append(path, nodeId)
		}
	}
	return path, length
}
