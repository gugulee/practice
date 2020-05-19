module github.com/practice

go 1.13

require (
	cni-operator-mod v0.0.0-incompatible
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.1 // indirect
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.7.0
	github.com/stretchr/testify v1.4.0
	go.etcd.io/etcd v3.3.20+incompatible
	google.golang.org/genproto v0.0.0-20200430143042-b979b6f78d84 // indirect
	google.golang.org/grpc v1.29.1 // indirect
	mgr-operator v0.0.0-incompatible
)

replace mgr-operator => ../../mgr-operator

replace cni-operator-mod => ../../cni-operator-mod

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
