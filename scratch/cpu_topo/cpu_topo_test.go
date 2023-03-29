package cpu_topo

import (
	"fmt"
	"log"
	"testing"

	info "github.com/google/cadvisor/info/v1"
	"github.com/google/cadvisor/utils/sysfs"
	"github.com/google/cadvisor/utils/sysinfo"
)

// go test ./scratch/cpu_topo -c -o /tmp/test
func Test_getNodeCpuTopology(t *testing.T) {
	topology, err := getNodeCpuTopology(func() ([]info.Node, error) {
		nodeInfo, _, err := sysinfo.GetNodesInfo(sysfs.NewRealSysFs())
		if err != nil {
			return nil, fmt.Errorf("get node cpu info: %v", err)
		}
		return nodeInfo, nil
	})

	if err != nil {
		log.Fatalf("get node cpu info: %v", err)
	}

	for _, cell := range topology.NumaCells {
		fmt.Printf("cell id is %d\n", cell.Id)
		for _, cpu := range cell.Cpus {
			fmt.Printf("%d %v\n", cpu.Id, cpu.Siblings)
		}
	}

}
