package ovsdb

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/go-logr/stdr"
	"github.com/kubeovn/kube-ovn/pkg/ovsdb/ovnnb"
	ovsdbclient "github.com/ovn-org/libovsdb/client"
	"github.com/ovn-org/libovsdb/model"

	"k8s.io/klog/v2"
)

// NewNbClient creates a new OVN NB client
func NewNbClient(addr string, timeout int) (ovsdbclient.Client, error) {
	dbModel, err := ovnnb.FullDatabaseModel()
	if err != nil {
		return nil, err
	}

	dbModel.SetIndexes(map[string][]model.ClientIndex{
		"Logical_Switch_Port": {{Columns: []model.ColumnKey{{Column: "external_ids", Key: "security_groups"}}}},
	})

	logger := stdr.NewWithOptions(log.New(os.Stderr, "", log.LstdFlags), stdr.Options{LogCaller: stdr.All}).
		WithName("libovsdb").
		WithValues("database", dbModel.Name())
	options := []ovsdbclient.Option{
		ovsdbclient.WithReconnect(time.Duration(timeout)*time.Second, &backoff.ZeroBackOff{}),
		ovsdbclient.WithLeaderOnly(false),
		ovsdbclient.WithLogger(&logger),
	}

	var ssl bool
	for _, ep := range strings.Split(addr, ",") {
		if !ssl && strings.HasPrefix(ep, ovsdbclient.SSL) {
			ssl = true
		}
		options = append(options, ovsdbclient.WithEndpoint(ep))
	}

	if ssl {
		cert, err := tls.LoadX509KeyPair("/var/run/tls/cert", "/var/run/tls/key")
		if err != nil {
			return nil, fmt.Errorf("failed to load x509 cert key pair: %v", err)
		}
		caCert, err := os.ReadFile("/var/run/tls/cacert")
		if err != nil {
			return nil, fmt.Errorf("failed to read ca cert: %v", err)
		}

		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM(caCert)
		// #nosec
		tlsConfig := &tls.Config{
			Certificates:       []tls.Certificate{cert},
			RootCAs:            certPool,
			InsecureSkipVerify: true,
		}

		options = append(options, ovsdbclient.WithTLSConfig(tlsConfig))
	}

	c, err := ovsdbclient.NewOVSDBClient(dbModel, options...)
	if err != nil {
		return nil, err
	}

	if err = c.Connect(context.TODO()); err != nil {
		klog.Errorf("failed to connect to OVN NB server %s: %v", addr, err)
		return nil, err
	}

	monitorOpts := []ovsdbclient.MonitorOption{
		ovsdbclient.WithTable(&ovnnb.LogicalRouter{}),
		ovsdbclient.WithTable(&ovnnb.LogicalRouterPolicy{}),
		ovsdbclient.WithTable(&ovnnb.LogicalSwitchPort{}),
		ovsdbclient.WithTable(&ovnnb.PortGroup{}),
	}
	if _, err = c.Monitor(context.TODO(), c.NewMonitor(monitorOpts...)); err != nil {
		klog.Errorf("failed to monitor database on OVN NB server %s: %v", addr, err)
		return nil, err
	}

	return c, nil
}

func list(client ovsdbclient.Client) error {
	name := "sg-test-5bc575d858-vhllv.lee.attachnet1.mec-nets.ovn"

	lsp := &ovnnb.LogicalSwitchPort{Name: name}
	if err := client.Get(context.TODO(), lsp); err != nil {

		return fmt.Errorf("failed to get logical switch port %s: %v", name, err)
	}

	fmt.Println(lsp)

	return nil
}

func ListLogicalSwitchPorts(client ovsdbclient.Client, externalIDs map[string]string) ([]ovnnb.LogicalSwitchPort, error) {
	lspList := make([]ovnnb.LogicalSwitchPort, 0)

	if err := client.WhereCache(func(lsp *ovnnb.LogicalSwitchPort) bool {
		if lsp.Type != "" {
			return false
		}

		if len(lsp.ExternalIDs) < len(externalIDs) {
			return false
		}

		if len(lsp.ExternalIDs) != 0 {
			for k, v := range externalIDs {
				// if only key exist but not value in externalIDs, we should include this lsp,
				// it's equal to shell command `ovn-nbctl --columns=xx find logical_switch_port external_ids:key!=\"\"`
				if 0 == len(v) {
					if 0 == len(lsp.ExternalIDs[k]) {
						return false
					}
				} else {
					if lsp.ExternalIDs[k] != v {
						return false
					}
				}

			}
		}
		return true
	}).List(context.TODO(), &lspList); err != nil {
		klog.Errorf("failed to list logical switch ports: %v", err)
		return nil, err
	}

	return lspList, nil
}

func ListLogicalSwitchPortsWithIdx(client ovsdbclient.Client, externalIDs map[string]string) ([]ovnnb.LogicalSwitchPort, error) {
	lsp := &ovnnb.LogicalSwitchPort{
		// ExternalIDs: map[string]string{"security_groups": ""},
		ExternalIDs: map[string]string{"security_groups": externalIDs["security_groups"]},
	}

	lspList := make([]ovnnb.LogicalSwitchPort, 0)

	// quick indexed result
	client.Where(lsp).List(context.TODO(), &lspList)

	return lspList, nil
}

func set(client ovsdbclient.Client) error {
	nbGlobal := &ovnnb.NBGlobal{}
	nbGlobalList := make([]ovnnb.NBGlobal, 0)
	if err := client.List(context.TODO(), &nbGlobal); err != nil {
		return fmt.Errorf("failed to get nbGlobal : %v", err)
	}

	// nbGlobal := &ovnnb.NBGlobal{}

	// client.WhereCache(
	// 	func(ls *ovnnb.NBGlobal) bool {
	// 		return true
	// 		// return 0 != len(ls.Connections)
	// 	}).List(context.TODO(), &nbGlobalList)

	// fmt.Println(nbGlobalList)

	// client.Where(nbGlobal, model.Condition{
	// 	Field:    &nbGlobal.UUID,
	// 	Function: ovsdb.ConditionEqual,
	// 	Value:    "a10860a5-78bc-4e77-a00b-4521dd84ea71",
	// }).List(context.TODO(), &nbGlobalList)

	fmt.Println(nbGlobalList)
	fmt.Println(nbGlobal)
	return nil
}
