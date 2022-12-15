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

		id   string
		name string
	)

	BeforeEach(func() {
		id = "some id"
		name = "some name"
	})

	Describe("Creating a document", func() {
		JustBeforeEach(func() {
			document, err = domain.NewDocument(id, name)
		})

		Context("when id is empty", func() {
			BeforeEach(func() {
				id = ""
			})

			It("should return a proper document", func() {
				Expect(document).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			It("should have proper attributes", func() {
				Expect(document.Id()).To(Equal(id))
				Expect(document.Name()).To(Equal(name))
			})

			It("should NOT be saved", func() {
				Expect(document.IsSaved()).To(Equal(false))
			})
		})

		Context("when name is empty", func() {
			BeforeEach(func() {
				name = ""
			})

			It("should return a proper document", func() {
				Expect(document).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			It("should have proper attributes", func() {
				Expect(document.Id()).To(Equal(id))
				Expect(document.Name()).To(Equal(domain.DefaultFileName))
			})

			It("should be saved", func() {
				Expect(document.IsSaved()).To(Equal(true))
			})
		})

		Context("when all parameters are defined", func() {
			It("should return a proper document", func() {
				Expect(document).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			It("should have proper attributes", func() {
				Expect(document.Id()).To(Equal(id))
				Expect(document.Name()).To(Equal(name))
			})

			It("should be saved", func() {
				Expect(document.IsSaved()).To(Equal(true))
			})
		})
	})

	Describe("Creating an unsaved document", func() {
		JustBeforeEach(func() {
			document, err = domain.NewUnsavedDocument(name)
		})

		Context("when name is empty", func() {
			BeforeEach(func() {
				name = ""
			})

			It("should return a proper document", func() {
				Expect(document).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			It("should have proper attributes", func() {
				Expect(document.Id()).To(Equal(""))
				Expect(document.Name()).To(Equal(domain.DefaultFileName))
			})

			It("should NOT be saved", func() {
				Expect(document.IsSaved()).To(Equal(false))
			})
		})

		Context("when all parameters are defined", func() {
			It("should return a proper document", func() {
				Expect(document).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			It("should have proper attributes", func() {
				Expect(document.Id()).To(Equal(""))
				Expect(document.Name()).To(Equal(name))
			})

			It("should NOT be saved", func() {
				Expect(document.IsSaved()).To(Equal(false))
			})
		})
	})
})
