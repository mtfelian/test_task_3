package api_test

import (
	"fmt"

	"github.com/mtfelian/test_task_3/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("validation tests", func() {
	It("tests validation of GetParamsBody", func() {
		testCases := []struct {
			in      api.GetParamBody
			isValid bool
		}{
			{api.GetParamBody{Type: "t", Data: "d"}, true},
			{api.GetParamBody{Type: "t", Data: ""}, false},
			{api.GetParamBody{Type: "", Data: "d"}, false},
			{api.GetParamBody{Type: "", Data: ""}, false},
		}

		for i, tc := range testCases {
			By(fmt.Sprintf("testing case %d: %v...", i, testCases))

			hasErrors, details := tc.in.Validate()
			if tc.isValid {
				Expect(details).To(Equal(""))
			} else {
				Expect(details).NotTo(Equal(""))
			}

			Expect(tc.isValid).ToNot(Equal(hasErrors))
		}
	})
})
