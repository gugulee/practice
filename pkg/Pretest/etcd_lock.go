package pretest
// package main

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"time"

// 	"mgr-operator/pkg/util/etcd"

// 	"go.etcd.io/etcd/clientv3"
// )

// // EtcdMutex is the lock of etcd
// type EtcdMutex struct {
// 	// TTL is the advisory time-to-live in seconds. 
// 	TTL     int64

// 	// the etcd client
// 	cli     *clientv3.Client

// 	// the key of the lock
// 	Key     string

// 	// the function when unlock
// 	cancel  context.CancelFunc 
// 	lease   clientv3.Lease
// 	leaseID clientv3.LeaseID
// 	txn     clientv3.Txn
// }

// // init initialise the lock
// func (em *EtcdMutex) init() error {
// 	var err error
// 	var ctx context.Context

// 	em.txn = clientv3.NewKV(em.cli).Txn(context.TODO())
// 	em.lease = clientv3.NewLease(em.cli)
// 	leaseResp, err := em.lease.Grant(context.TODO(), em.TTL)
// 	if err != nil {
// 		return err
// 	}
// 	ctx, em.cancel = context.WithCancel(context.TODO())
// 	em.leaseID = leaseResp.ID
// 	_, err = em.lease.KeepAlive(ctx, em.leaseID)
// 	return err
// }

// // Lock acquire the lock
// func (em *EtcdMutex) Lock() error {
// 	err := em.init()
// 	if err != nil {
// 		return err
// 	}
// 	//LOCK:
// 	em.txn.If(clientv3.Compare(clientv3.CreateRevision(em.Key), "=", 0)).
// 		Then(clientv3.OpPut(em.Key, "", clientv3.WithLease(em.leaseID))).
// 		Else()
// 	txnResp, err := em.txn.Commit()
// 	if err != nil {
// 		return err
// 	}
// 	if !txnResp.Succeeded { //判断txn.if条件是否成立
// 		return fmt.Errorf("acquire lock failed")
// 	}
// 	return nil
// }

// // UnLock release the lock
// func (em *EtcdMutex) UnLock() {
// 	em.cancel()
// 	em.lease.Revoke(context.TODO(), em.leaseID)
// 	fmt.Println("unlock")
// }

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

// 	eMutex := &EtcdMutex{
// 		Key: "mgr.huacloud.com/lock/cluster/mgr",
// 		TTL: 5,
// 		cli: cli,
// 	}
// 	//groutine1
// 	go func() {
// 		err := eMutex.Lock()
// 		if err != nil {
// 			fmt.Println("groutine抢锁失败")
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Println("groutine抢锁成功")
// 		time.Sleep(5 * time.Second)
// 		defer eMutex.UnLock()
// 	}()

// 	time.Sleep(30 * time.Second)
// }
