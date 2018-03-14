package main_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing with Ginkgo", func() {
	pingAddr := func() string { return server.URL }

	Context("api.Ping request", func() {
		It("checks ping", func() {
			addr := fmt.Sprintf("%s/ping", pingAddr())
			resp, err := doRequest(http.MethodGet, addr, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			b, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			var r string
			Expect(json.Unmarshal(b, &r)).To(Succeed())
			Expect(r).To(Equal("pong"))
		})
	})
})
