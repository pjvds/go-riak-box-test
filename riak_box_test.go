package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//"github.com/tpjg/goriakpbc"
	"os"
	"testing"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Riak Suite")
}

var _ = Describe("Riak", func() {
	Context("When using the riak box at wercker", func() {
		It("should export PBCONNECT", func() {
			Expect(os.Getenv("MJDSYS_RIAK_PBCONNECT")).ToNot(BeEmpty())
		})
	})
})
