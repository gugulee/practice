module github.com/practice

go 1.13

require (
	github.com/coreos/etcd v3.3.20+incompatible
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000
	github.com/deckarep/golang-set v1.7.1
	github.com/golang/protobuf v1.4.1 // indirect
	github.com/google/uuid v1.1.1
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.7.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/vishvananda/netlink v1.1.0
	go.etcd.io/etcd v3.3.20+incompatible
	golang.org/x/text v0.3.2
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135
	google.golang.org/genproto v0.0.0-20200430143042-b979b6f78d84 // indirect
	google.golang.org/grpc v1.29.1 // indirect
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v0.18.2
	k8s.io/gengo v0.0.0-20190128074634-0689ccc1d7d6
	k8s.io/klog v1.0.0
	sigs.k8s.io/kustomize v2.0.3+incompatible
)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
