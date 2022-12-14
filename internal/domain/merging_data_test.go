package domain_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/theobitoproject/gedo_mazo/internal/domain"
)

var _ = Describe("MergingData", func() {
	var (
		mergingData *domain.MergingData
		err         error

		key   string
		value string
	)

	Describe("Creating a merging data", func() {
		JustBeforeEach(func() {
			mergingData, err = domain.NewMergingData()
		})

		Context("when all parameters are defined", func() {
			It("should return a proper merging data", func() {
				Expect(mergingData).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Setting a key", func() {
		JustBeforeEach(func() {
			err = mergingData.SetKey(key, value)
		})

		BeforeEach(func() {
			mergingData, err = domain.NewMergingData()
			Expect(err).To(BeNil())

			key = "some key"
			value = "some value"
		})

		Context("when key is empty", func() {
			BeforeEach(func() {
				key = ""
			})

			It("should throw an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})

		Context("when key is NOT empty", func() {
			Context("when value is empty", func() {
				BeforeEach(func() {
					value = ""
				})

				It("should NOT throw an error", func() {
					Expect(err).To(BeNil())
				})
			})

			Context("when value is NOT empty", func() {
				It("should NOT throw an error", func() {
					Expect(err).To(BeNil())
				})
			})
		})
	})
})
