package testingIntro
import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestingIntro", func() {

	Context("first func", func(){
		It("Given the name Josh", func(){
			Expect(firstTestFunc("Josh")).To(Equal("Hello Josh"))
		})
		It("Given the name Jacob", func(){
			Expect(firstTestFunc("Jacob")).To(Equal("Hello Jacob"))
		})
	})
})
