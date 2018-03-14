package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	e "github.com/mtfelian/error"
	"github.com/mtfelian/test_task_3/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing with Ginkgo", func() {
	Context("api.GetParam request", func() {
		validInputData := func() api.GetParamBody { return api.GetParamBody{Data: "d", Type: "t"} }
		invalidInputData := func() api.GetParamBody { return api.GetParamBody{Data: "", Type: "t"} }

		It("checks getting param if OK", func() {
			bInput, err := json.Marshal(validInputData())
			Expect(err).NotTo(HaveOccurred())

			addr := fmt.Sprintf("%s/param", server.URL)
			resp, err := doRequest(http.MethodPost, addr, bytes.NewReader(bInput))
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			b, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			_ = b
		})

		It("checks error if input body has not valid data structure", func() {
			addr := fmt.Sprintf("%s/param", server.URL)
			resp, err := doRequest(http.MethodPost, addr, bytes.NewReader([]byte(`"fake"`)))
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusUnprocessableEntity))

			b, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			var r e.StandardError
			Expect(json.Unmarshal(b, &r)).To(Succeed())
			Expect(r.Code()).To(BeEquivalentTo(api.ErrorInvalidInput))
		})

		It("checks error if input body invalid", func() {
			bInput, err := json.Marshal(invalidInputData())
			Expect(err).NotTo(HaveOccurred())

			addr := fmt.Sprintf("%s/param", server.URL)
			resp, err := doRequest(http.MethodPost, addr, bytes.NewReader(bInput))
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusUnprocessableEntity))

			b, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			var r e.StandardError
			Expect(json.Unmarshal(b, &r)).To(Succeed())
			Expect(r.Code()).To(BeEquivalentTo(api.ErrorValidation))
		})
	})
})
