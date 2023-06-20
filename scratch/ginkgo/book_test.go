package ginkgo

import (
	. "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega"
)

var _ = Describe("some JSON decoding edge cases", func() {
	var json string
	JustBeforeEach(func() {
		_ = json
	})
	When("the JSON fails to parse", Label("json test"), func() {
		BeforeEach(func() {
			json = "abc"
		})

		It("errors", func() {

		})
	})

	When("the JSON is incomplete", func() {
		BeforeEach(func() {
			json = "123"
		})

		It("errors", func() {

		})
	})
})
