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
})
