package main_test

import (
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/mtfelian/test_task_3"
	"github.com/mtfelian/test_task_3/config"
	"github.com/mtfelian/test_task_3/service"
	"github.com/mtfelian/test_task_3/storage"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

var server *httptest.Server

func doRequest(method, addr string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, addr, body)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(request)
}

var _ = Describe("register suite hooks", func() {
	BeforeSuite(func() {
		// set DSN
		viper.Set(config.DSN, storage.PostgresDSNTests())
		viper.Set(config.DBDriver, storage.DriverNamePostgres)

		service.New()
		s := service.Get()
		httpServer := s.HTTPServer
		main.RegisterHTTPAPIHandlers(httpServer)
		server = httptest.NewServer(httpServer)

		// set server port if needed
		URL, err := url.Parse(server.URL)
		Expect(err).NotTo(HaveOccurred())
		port, err := strconv.Atoi(URL.Port())
		Expect(err).NotTo(HaveOccurred())
		viper.Set(config.Port, uint(port))

		Expect(s.Storage).NotTo(BeNil())
		Expect(s.Storage.DeleteAll()).To(Succeed())
		Expect(s.Storage.GenerateTestEntries(1000, 200)).To(Succeed())
	})

	AfterSuite(func() { server.Close() })
})

func TestAll(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	RegisterFailHandler(Fail)
	RunSpecs(t, "Main Suite")
}
