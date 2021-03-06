package wercker

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tpjg/goriakpbc"
	"os"
	"testing"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Riak Suite")
}

var _ = Describe("Riak", func() {
	Context("When using the riak box at wercker", func() {
		host := os.Getenv("MJDSYS_RIAK_HOST")

		It("should export MJDSYS_RIAK_HOST", func() {
			Expect(host).ToNot(BeEmpty())
		})
		It("should accept connections", func() {
			client := riak.New(host + ":8087")
			err := client.Connect()
			Expect(err).To(BeNil())

			client.Close()
		})
		It("should not accept connections on localhost", func() {
			client := riak.New("localhost:8087")
			err := client.Connect()
			Expect(err).ToNot(BeNil())

			client.Close()
		})
	})
})
