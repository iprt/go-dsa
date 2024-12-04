// Package topology
// @Description: 拓扑排序
package topology

import "go-dsa/pkg/graph"

type topology struct {
	g graph.Graph
}

func CreateTopologyAlgo(g graph.Graph) *topology {
	return &topology{
		g: g,
	}
}
