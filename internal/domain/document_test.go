package domain_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/theobitoproject/gedo_mazo/internal/domain"
)

var _ = Describe("Document", func() {
	var (
		document *domain.Document
		err      error

		id string
	)

	BeforeEach(func() {
		id = "some id"
	})

	Describe("Creating a document", func() {
		JustBeforeEach(func() {
			document, err = domain.NewDocument(id)
		})

		Context("when id is empty", func() {
			BeforeEach(func() {
				id = ""
			})

			It("should throw an error", func() {
				Expect(document).To(BeNil())
				Expect(err).ToNot(BeNil())
			})
		})

		Context("when all parameters are defined", func() {
			It("should return a proper document", func() {
				Expect(document).ToNot(BeNil())
				Expect(err).To(BeNil())

				Expect(document.Id()).To(Equal(id))
			})
		})
	})
})
