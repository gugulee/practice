package pretest

// package main

// import (
// 	"fmt"
// 	"github.com/vishvananda/netlink"
// 	"net"
// )

// func main() {
// 	name := "eth1"
// 	tableID := 1
// 	eth, err := netlink.LinkByName(name)
// 	if err != nil {
// 		fmt.Printf("failed to lookup device %q: %v\n", name, err)
// 		return
// 	}
// 	ipa, err := netlink.AddrList(eth, netlink.FAMILY_V4)
// 	if err != nil {
// 		fmt.Printf("failed to lookup ip %q: %v\n", name, err)
// 		return
// 	}
// 	_, netID, _ := net.ParseCIDR(ipa[0].IPNet.String())
// 	fmt.Println(netID)

// 	//rule := netlink.NewRule()
// 	//rule.Src, rule.IifName, rule.Table = netID, "lo", tableID

// 	rule := &netlink.Rule{
// 		Src:               netID,
// 		IifName:           "lo",
// 		Table:             tableID,
// 		SuppressIfgroup:   -1,
// 		SuppressPrefixlen: -1,
// 		Priority:          -1,
// 		Mark:              -1,
// 		Mask:              -1,
// 		Goto:              -1,
// 		Flow:              -1,
// 	}

// 	err = netlink.RuleAdd(rule)
// 	if err != nil {
// 		fmt.Printf("ip rule add %q failed: %v\n", name, err)
// 		return
// 	}
// 	_, defNet, _ := net.ParseCIDR("0.0.0.0/0")
// 	err = netlink.RouteAdd(&netlink.Route{
// 		LinkIndex: eth.Attrs().Index,
// 		Scope:     netlink.SCOPE_UNIVERSE,
// 		Dst:       defNet,
// 		Gw:        net.ParseIP("10.20.20.1"),
// 		Table:     tableID,
// 	})
// 	if nil != err {
// 		fmt.Printf("ip route add %q failed: %v\n", name, err)
// 	}
// }
