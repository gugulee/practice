package pretest

// package main

// import (
// 	"fmt"
// 	"go.etcd.io/etcd/clientv3"
// 	"mgr-operator/pkg/util/etcd"
// 	"os"
// )

// func main() {
// 	var cli *clientv3.Client

// 	clis, err := etcd.NewTLSClient(&etcd.ETCD{
// 		Endpoints: "https://172.16.7.81:2379,https://172.16.7.82:2379,https://172.16.7.83:2379",
// 		CaFile:    `etcd/ca.crt`,
// 		CertFile:  `etcd/peer.crt`,
// 		KeyFile:   `etcd/peer.key`,
// 	})

// 	if nil != err {
// 		fmt.Printf("check etcd connection failed, error: %s\n", err)
// 		os.Exit(1)
// 	}

// 	if cli = etcd.CheckConn(clis); cli == nil {
// 		fmt.Printf("no endpoint connect normally\n")
// 		os.Exit(1)
// 	}
// 	defer func() {
// 		for _, cli := range clis {
// 			_ = cli.Close()
// 		}
// 	}()

// 	key := fmt.Sprintf("/%s/%s/%s/%s/", "ip-manager", "mgr", "mgr", "vip")

// 	values, err := etcd.GetMultiValue(cli, key, "")
// 	if nil != err {
// 		fmt.Printf("get, error: %s\n", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println(values)
// }
