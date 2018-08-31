package cart_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail) // makes the connection between ginkgo and gomega; registers ginkgo.Fail with gomega
	RunSpecs(t, "Shopping Cart Suite")
}

var _ = Describe("Shopping Cart", func() {
	Context("initially", func() {
		It("has 0 items", func() {})
		It("has 0 units", func() {})
		Specify("the total amount is 0.00", func() {})
	})
	Context("when a new item is added", func() {
		Context("the shopping cart", func() {
			It("has 1 more unique item than it had earlier", func() {})
			It("has 1 more unit than it had earlier", func() {})
			Specify("the total amount increases by item price", func() {})
		})
	})
	Context("when an existing item is added", func() {
		Context("the shopping cart", func() {
			It("has the same number of unique items as earlier", func() {})
			It("has 1 more of the item than earlier", func() {})
			Specify("the total amount increases by the item price", func() {})
		})
	})
	Context("that has 0 units of item A", func() {
		Context("remove item A", func() {
			It("should not change the number of unique items", func() {})
			It("should not change the number of units", func() {})
			It("should not change the total amount", func() {})
		})
	})
	Context("that has 1 unit of item A", func() {
		Context("remove 1 unit of item A", func() {
			It("should reduce the number of unique items by 1", func() {})
			It("should reduce the number of units by 1", func() {})
			It("should reduce the total amount by the item price", func() {})
		})
	})
	Context("that has 2 unit of item A", func() {
		Context("remove 1 unit of item A", func() {
			It("should have the same number of unique items as earlier", func() {})
			It("should reduce the number of units by 1", func() {})
			It("should reduce the total amount by the item price", func() {})
		})
		Context("remove 2 units of item A", func() {
			It("should reduce the number of unique items by 1", func() {})
			It("should reduce the number of units by 2", func() {})
			It("should reduce the total amount by twice the item price", func() {})
		})
	})
})
