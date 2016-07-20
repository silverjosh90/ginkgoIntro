package testingIntro
import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Shopping Cart", func() {

	Context("add Item", func(){
		var(
			newCart ShoppingCart
		)
		It("Given an item, should add it to the shopping cart list", func(){
			newCart.addItem("potatoes")
			Expect(newCart.inventory[0]).To(Equal("potatoes"))
		})

		It("Can have multiple items in inventory", func(){
			newCart.addItem("Spam")
			Expect(newCart.inventory[1]).To(Equal("Spam"))
		})

	})
	Context("remove item", func() {
		var(
			newCart ShoppingCart
		)
		It("Can remove items", func(){
			newCart.addItem("potatoes")
			newCart.addItem("taco")
			newCart.removeItem("potatoes")
			Expect(newCart.inventory[0]).To(Equal("taco"))
		})

		It("can remove from the middle of the array", func(){
			newCart.addItem("chimi")
			newCart.addItem("egg")
			newCart.addItem("borche")
			newCart.removeItem("borche")
			Expect(len(newCart.inventory)).To(Equal(3))
		})

	})
})
