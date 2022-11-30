package virt

import (
	"fmt"
	"log"

	"libvirt.org/go/libvirt"
)

const memoryCellTemplate = `<memory model='dimm'>
  <target>
    <size unit='%s'>%d</size>
    <node>%d</node>
  </target>
</memory>`

func main() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		log.Fatalf("new connection: %v", err)
	}
	defer conn.Close()

	dom, err := conn.LookupDomainByName("mec-images_test-lee")
	if err != nil {
		log.Fatalf("look up domain name: %v", err)
	}

	var targetMemory uint64 = 5368709120

	info, err := dom.GetInfo()
	if err != nil {
		log.Fatalf("get current memory: %v", err)
	}

	currentMemory := info.Memory * 1024

	if currentMemory >= targetMemory {
		log.Printf("memory %d already equal to desired memory %d", currentMemory, targetMemory)
		return
	}

	incrementalMemory := (targetMemory - currentMemory) / 1024 / 1024
	memoryDevice := fmt.Sprintf(memoryCellTemplate, "MiB", incrementalMemory, 0)

	log.Printf("current memory is %d, target memory is %d", currentMemory, targetMemory)

	if err = dom.AttachDeviceFlags(memoryDevice, libvirt.DOMAIN_DEVICE_MODIFY_CONFIG|libvirt.DOMAIN_DEVICE_MODIFY_LIVE); err != nil {
		log.Fatalf("add memory to guest on fly: %s", err)
	}
}