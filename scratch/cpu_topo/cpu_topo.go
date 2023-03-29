package cpu_topo

import (
	"fmt"
	"sort"

	cmdv1 "ecs/pkg/kubevirt/handler-launcher-com/cmd/v1"

	info "github.com/google/cadvisor/info/v1"
)

// getNodeCpuTopology get cpu topology on node
func getNodeCpuTopology(getNodesInfo func() ([]info.Node, error)) (*cmdv1.Topology, error) {
	nodes, err := getNodesInfo()
	if err != nil {
		return nil, fmt.Errorf("get node cpu info: %v", err)
	}

	topology := &cmdv1.Topology{
		NumaCells: make([]*cmdv1.Cell, 0, len(nodes)),
	}

	for _, node := range nodes {
		cell := &cmdv1.Cell{
			Id: uint32(node.Id),
		}

		for _, core := range node.Cores {
			siblings := []uint32{}
			sort.Ints(core.Threads) // keep threads increasing order

			for _, thread := range core.Threads {
				siblings = append(siblings, uint32(thread))
			}

			for _, thread := range core.Threads {
				cell.Cpus = append(cell.Cpus, &cmdv1.CPU{
					Id:       uint32(thread),
					Siblings: siblings,
				})
			}
		}

		topology.NumaCells = append(topology.NumaCells, cell)
	}

	return topology, nil
}
