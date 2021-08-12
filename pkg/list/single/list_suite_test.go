package single

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLinklist(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "list Suite")
}
