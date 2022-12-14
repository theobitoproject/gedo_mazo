package domain_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/theobitoproject/gedo_mazo/internal/domain"
)

var _ = Describe("Folder", func() {
	var (
		folder *domain.Folder
		err    error

		id string
	)

	BeforeEach(func() {
		id = "some id"
	})

	Describe("Creating a folder", func() {
		JustBeforeEach(func() {
			folder, err = domain.NewFolder(id)
		})

		Context("when id is empty", func() {
			BeforeEach(func() {
				id = ""
			})

			It("should throw an error", func() {
				Expect(folder).To(BeNil())
				Expect(err).ToNot(BeNil())
			})
		})

		Context("when all parameters are defined", func() {
			It("should return a proper folder", func() {
				Expect(folder).ToNot(BeNil())
				Expect(err).To(BeNil())

				Expect(folder.Id()).To(Equal(id))
			})
		})
	})
})
