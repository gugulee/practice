package main

import (
	"fmt"

	"gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvpci"
)

func getBARsMaxAddressableMemory() (uint64, uint64) {

	pci := nvpci.New()
	devs, _ := pci.GetAllDevices()

	// Since we do not know which devices are going to be hotplugged,
	// we're going to use the GPU with the biggest BARs to initialize the
	// root port, this should work for all other devices as well.
	// defaults are 2MB for both, if no suitable devices found
	max32bit := uint64(2 * 1024 * 1024)
	max64bit := uint64(2 * 1024 * 1024)

	for _, dev := range devs {
		if !dev.IsGPU() {
			continue
		}
		memSize32bit, memSize64bit := dev.Resources.GetTotalAddressableMemory(true)
		if max32bit < memSize32bit {
			max32bit = memSize32bit
		}
		if max64bit < memSize64bit {
			max64bit = memSize64bit
		}
	}
	// The actual 32bit is most of the time a power of 2 but we need some
	// buffer so double that to leave space for other IO functions.
	// The 64bit size is not a power of 2 and hence is already rounded up
	// to the higher value.
	return max32bit * 2, max64bit
}

func main() {
	memSize32bit, memSize64bit := getBARsMaxAddressableMemory()
	fmt.Println(memSize32bit)
	fmt.Println(memSize64bit)
}
