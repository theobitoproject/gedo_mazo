package tools_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/theobitoproject/gedo_mazo/tools"
)

var _ = Describe("Pointers", func() {
	Describe("Checking a value is nil", func() {
		var (
			value interface{}
			isNil bool
		)

		JustBeforeEach(func() {
			isNil = tools.IsNil(value)
		})

		Context("when value is NOT a pointer", func() {
			BeforeEach(func() {
				value = "some text"
			})

			It("should return a false", func() {
				Expect(isNil).To(Equal(false))
			})
		})

		Context("when value is a pointer", func() {
			Context("when value is NOT nil", func() {
				BeforeEach(func() {
					var txt string = "some text"
					value = &txt
				})

				It("should return a false", func() {
					Expect(isNil).To(Equal(false))
				})
			})

			Context("when value is nil", func() {
				BeforeEach(func() {
					var txt *string = nil
					value = txt
				})

				It("should return a true", func() {
					Expect(isNil).To(Equal(true))
				})
			})
		})
	})
})
