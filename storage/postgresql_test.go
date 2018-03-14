package storage_test

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mtfelian/test_task_3/config"
	"github.com/mtfelian/test_task_3/models"
	"github.com/mtfelian/test_task_3/storage"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

// model1Model returns a pointer to a valid test model
var model1Model = func() *models.Param {
	return &models.Param{}
}

// initDB initializes DB for tests with given DSN string
func initDB(DSN string) *gorm.DB {
	db, err := gorm.Open(viper.GetString(config.DBDriver), DSN)
	Expect(err).NotTo(HaveOccurred())
	return db
}

func expectModelsAreEqual(model1, model2 models.Param) {
	model1.ID, model2.ID = 0, 0
	Expect(model1).To(Equal(model2))
}

var _ = Describe("testing Model1 Model via Postgres storage", func() {
	var keeper storage.Keeper
	BeforeEach(func() {
		viper.Set(config.DSN, storage.PostgresDSNTests())
		viper.Set(config.DBDriver, storage.DriverNamePostgres)
		keeper = storage.NewPostgresWithDB(initDB(viper.GetString(config.DSN)))
		Expect(keeper.DeleteAll()).To(Succeed())
	})
	AfterEach(func() { Expect(keeper.(storage.Postgres).DB().Close()).To(Succeed()) })

	Context("CRUD", func() {
		It("checks adding, loading log entry model", func() {
			model := model1Model()
			Expect(keeper.Add(model)).To(Succeed())
			loadedModel, err := keeper.Get(model.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(loadedModel).NotTo(BeNil())
			expectModelsAreEqual(*loadedModel, *model)
		})
	})

	Context("with mass generating", func() {
		// something
	})
})
