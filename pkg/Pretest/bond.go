package pretest

// package main

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/vishvananda/netlink"
// )

// func main() {

// 	js := []byte(`{"mode":1,"miimon":100}`)
// 	bondName := "bond0"
// 	vlanName := "bond0.100"

// 	linkAttrs := netlink.NewLinkAttrs()
// 	bondDev := netlink.NewLinkBond(linkAttrs)
// 	bondDev.Name = bondName

// 	if err := json.Unmarshal(js, bondDev); nil != err {
// 		fmt.Println(err)
// 		return
// 	}

// 	if err := netlink.LinkAdd(bondDev); err != nil {
// 		fmt.Printf("failed to add device %q: %v", bondName, err)
// 		return
// 	}

// 	if err := netlink.LinkSetUp(bondDev); err != nil {
// 		fmt.Printf("failed to setup device %q: %v", bondName, err)
// 		return
// 	}

// 	linkAttrs = netlink.NewLinkAttrs()
// 	a := &netlink.Vlan{LinkAttrs: linkAttrs, VlanId: 100}
// 	a.Name = vlanName
// 	a.ParentIndex = bondDev.Index

// 	if err := netlink.LinkAdd(a); err != nil {
// 		fmt.Printf("failed to add device %q: %v", vlanName, err)
// 		return
// 	}

// 	if err := netlink.LinkSetUp(a); err != nil {
// 		fmt.Printf("failed to setup device %q: %v", vlanName, err)
// 		return
// 	}

// 	name := "A"
// 	test0, err := netlink.LinkByName(name)
// 	if err != nil {
// 		fmt.Printf("failed to lookup vf device %q: %v", name, err)
// 		return
// 	}

// 	if err := netlink.LinkSetBondSlave(test0, bondDev); err != nil {
// 		fmt.Printf("failed to set device %q to master %q: %v", name, bondName, err)
// 		return
// 	}

// 	name = "C"
// 	test2, err := netlink.LinkByName(name)

// 	if err != nil {
// 		fmt.Printf("failed to lookup vf device %q: %v", name, err)
// 		return
// 	}

// 	if err := netlink.LinkSetBondSlave(test2, bondDev); err != nil {
// 		fmt.Printf("failed to set device %q to master %q: %v", name, bondName, err)
// 		return
// 	}
// }
