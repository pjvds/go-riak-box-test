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
		pbconnect := os.Getenv("MJDSYS_RIAK_PBCONNECT")

		It("should export PBCONNECT", func() {
			Expect(pbconnect).ToNot(BeEmpty())
		})
		It("should accept connections via pbconnect", func() {
			client := riak.New(pbconnect)
			err := client.Connect()
			Expect(err).To(BeNil())

			client.Close()
		})
	})
})
