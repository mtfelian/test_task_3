package storage_test

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPostgres(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	RegisterFailHandler(Fail)
	RunSpecs(t, "Postgres Suite")
}
