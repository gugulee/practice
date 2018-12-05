package etcd_test

import (
	"crypto/tls"
	"crypto/x509"
	"time"
	"io/ioutil"
	"strings"

	"go.etcd.io/etcd/clientv3"
	"context"
	"fmt"
)

// ETCD represents the certificate information with etcd
type ETCD struct {
	Endpoints string `json:"endpoints"`
	CaFile    string `json:"caFile"`
	CertFile  string `json:"certFile"`
	KeyFile   string `json:"keyFile"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}

func NewTLSClient(conf *ETCD) ([]*clientv3.Client, error) {
	machines := strings.Split(conf.Endpoints, ",")
	clis := make([]*clientv3.Client, len(machines))
	// load cert
	cert, err := tls.LoadX509KeyPair(conf.CertFile, conf.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("tls.LoadX509KeyPair failed, error: %s", err)
	}
	// load root ca
	caData, err := ioutil.ReadFile(conf.CaFile)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile failed, error: %s", err)
	}

	pool := x509.NewCertPool()
	if ok := pool.AppendCertsFromPEM(caData); !ok {
		return nil, fmt.Errorf("pool.AppendCertsFromPEM failed")
	}

	_tlsConfig := &tls.Config{
		Certificates: [] tls.Certificate{cert},
		RootCAs:      pool,
	}
	for i, ep := range machines {
		clis[i], err = clientv3.New(clientv3.Config{
			Endpoints: []string{ep},
			TLS:       _tlsConfig,
			Username:  conf.Username,
			Password:  conf.Password,
		})

		if err != nil {
			return nil, fmt.Errorf("NewTLSClient failed,ep: %s error: %s", ep, err)
		}
	}
	return clis, nil
}

func CheckConn(clis []*clientv3.Client) *clientv3.Client {
	for _, cli := range clis {
		endpoint := cli.Endpoints()[0]
		ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
		if _, err := cli.Status(ctx, endpoint); err == nil {
			return cli
		}
	}
	return nil
}

func main() {
	clis, err := NewTLSClient(&ETCD{
		Endpoints: "https://172.16.18.10:2379",
		CaFile:    "ca.crt",
		CertFile:  "peer.crt",
		KeyFile:   "peer.key",
	})

	if nil != err {
		fmt.Println(err)
		return
	}
	cli := CheckConn(clis)
	if cli == nil {
		fmt.Println("no endpoint connect normally")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.Get(ctx, "/ip-manager/mgr/mgr/mgr/", clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Println(err)
		return
	}
	value := map[string]string{}
	for _, ev := range resp.Kvs {
		if strings.Contains(string(ev.Key), "eth0") {
			value[string(ev.Key)] = string(ev.Value)
			fmt.Printf("%s : %s\n", ev.Key, ev.Value)
		}
	}
	fmt.Println(value)
}
