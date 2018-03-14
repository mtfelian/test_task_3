package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	e "github.com/mtfelian/error"
	"github.com/mtfelian/test_task_3"
	"github.com/mtfelian/test_task_3/service"
	"github.com/mtfelian/test_task_3/storage"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing with Ginkgo", func() {
	Context("main.getParam request", func() {
		var (
			okInputData        = main.GetParamBody{Type: "Develop.mr_robot", Data: "Database.processing"}
			notExistsInputData = main.GetParamBody{Type: "t", Data: "d"}
			invalidInputData   = main.GetParamBody{Type: "t", Data: ""}
		)

		It("checks getting param if OK and param exists", func() {
			bInput, err := json.Marshal(okInputData)
			Expect(err).NotTo(HaveOccurred())

			addr := fmt.Sprintf("%s/param", server.URL)
			resp, err := doRequest(http.MethodPost, addr, bytes.NewReader(bInput))
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			b, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			Expect(string(b)).To(Equal(`{"host":"localhost","port":"5432","database":"devdb",` +
				`"user":"mr_robot","password":"secret","schema":"public"}`))
		})

		It("checks getting param if DB link broken", func() {
			Expect(service.Get().Storage.Close()).To(Succeed())
			defer func() { service.Get().Storage = storage.NewPostgres() }()

			bInput, err := json.Marshal(okInputData)
			Expect(err).NotTo(HaveOccurred())

			addr := fmt.Sprintf("%s/param", server.URL)
			resp, err := doRequest(http.MethodPost, addr, bytes.NewReader(bInput))
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))

			b, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			var r e.StandardError
			Expect(json.Unmarshal(b, &r)).To(Succeed())
			Expect(r.Code()).To(BeEquivalentTo(main.ErrorStorage))
		})

		It("checks getting param if OK but such param is not exists", func() {
			bInput, err := json.Marshal(notExistsInputData)
			Expect(err).NotTo(HaveOccurred())

			addr := fmt.Sprintf("%s/param", server.URL)
			resp, err := doRequest(http.MethodPost, addr, bytes.NewReader(bInput))
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			b, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			Expect(string(b)).To(Equal(""))
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
			Expect(r.Code()).To(BeEquivalentTo(main.ErrorInvalidInput))
		})

		It("checks error if input body invalid", func() {
			bInput, err := json.Marshal(invalidInputData)
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
			Expect(r.Code()).To(BeEquivalentTo(main.ErrorValidation))
		})
	})
})
