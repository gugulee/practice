package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
)

func main() {
        opts := gophercloud.AuthOptions{
                IdentityEndpoint: "http://10.99.1.41:35357/v3",
                //  openstack token issue
                TokenID: "gAAAAABliqW0hjDIwQaa3aMsV4663UKpIwACfhQCcAY0OkjAnbnQUEIoJXTCulWrv6bSRXSr1h4n5vEVXmjUdeTpRCJwBEWgKJ1BUrroGb2dcpADSeqq86vplIoa2YzgSGcwowIgv3pIhExKIFe5yZVhIEPPLuFLmbYHn1U3c_5SsHlHUkPoBsE",
        }

        provider, err := openstack.AuthenticatedClient(opts)
        if err != nil {
                log.Fatalf("AuthenticatedClient: %v", err)
        }

        client, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
                Region: "RegionOne",
        })

        if err != nil {
                log.Fatalf("NewNetworkV2: %v", err)
        }

        router, err := routers.Get(client, "f3720aa5-2c4a-4215-ab26-9ce44c1567cd").Extract()
        if err != nil {
                log.Fatalf("Get: %v", err)
        }

        log.Printf("%#v\n", router)

        log.Println()
        log.Println()
        fmt.Fprintf(os.Stdout, "xxxx+++++++++++\n")

        pager, err := ports.List(client, &ports.ListOpts{
                DeviceOwner: "network:router_interface",
                DeviceID:    "f3720aa5-2c4a-4215-ab26-9ce44c1567cd",
        }).AllPages()
        if err != nil {
                log.Fatalf("List: %v", err)
        }

        p, err := ports.ExtractPorts(pager)
        if err != nil {
                log.Fatalf("ExtractPorts: %v", err)
        }

        log.Printf("%#v\n", p)
}