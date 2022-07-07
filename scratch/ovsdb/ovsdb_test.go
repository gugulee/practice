package ovsdb

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const addr = "172.20.149.60"

func Test_list(t *testing.T) {
	nbClient, err := NewNbClient(fmt.Sprintf("tcp:[%s]:6641", addr), 60)
	require.NoError(t, err)

	err = list(nbClient)
	require.NoError(t, err)
}

func Test_ListLogicalSwitchPorts(t *testing.T) {
	nbClient, err := NewNbClient(fmt.Sprintf("tcp:[%s]:6641", addr), 60)
	require.NoError(t, err)

	ports, err := ListLogicalSwitchPorts(nbClient, map[string]string{"security_groups": "sg"})
	require.NoError(t, err)
	fmt.Printf("%v\n", ports)
}

func Test_ListLogicalSwitchPortsWithIdx(t *testing.T) {
	nbClient, err := NewNbClient(fmt.Sprintf("tcp:[%s]:6641", addr), 60)
	require.NoError(t, err)

	ports, err := ListLogicalSwitchPortsWithIdx(nbClient, map[string]string{"security_groups": "sg"})
	require.NoError(t, err)
	fmt.Printf("%#v\n", ports)
}

func Test_set(t *testing.T) {
	nbClient, err := NewNbClient(fmt.Sprintf("tcp:[%s]:6641", addr), 60)
	require.NoError(t, err)

	err = set(nbClient)
	require.NoError(t, err)
}
