package config

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	Port     = "port"
	FileName = "config"
	DSN      = "db"
	DBDriver = "dbdriver"
	Debug    = "debug"
)

// Parse the configuration data from different sources
func Parse() error {
	if err := parseFlags(); err != nil {
		log.Fatalln("Error parsing flags:", err)
	}
	if err := parseConfigFile(); err != nil {
		log.Fatalln("Error parsing config file:", err)
	}
	return nil
}

// parseFlags from command line
func parseFlags() error {
	pflag.UintVar(&params.Port, Port, 0, "application HTTP port")

	pflag.StringVar(&params.ConfigFileName, FileName, "config", "config file name")
	params.ConfigFileName = strings.TrimSuffix(params.ConfigFileName, filepath.Ext(params.ConfigFileName))

	pflag.StringVar(&params.DSN, DSN, "", "DSN data for DB access")
	pflag.StringVar(&params.DBDriver, DBDriver, "postgres", "DB driver name")
	pflag.BoolVar(&params.Debug, Debug, false, "enable debug mode")

	pflag.Parse()
	return viper.BindPFlags(pflag.CommandLine)
}

// parseConfigFile
func parseConfigFile(paths ...string) error {
	viper.AddConfigPath(".")
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	viper.SetConfigName(viper.GetString(FileName))
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to load config file: %v", err)
	}
	return nil
}

// Options describes command line options
type Options struct {
	sync.Mutex

	// Port to listen
	Port uint
	// DSN is a database access DSN string
	DSN string
	// DBDriver is a name of a DB driver
	DBDriver string
	// ConfigFileName is a configuration file name
	ConfigFileName string

	// Debug - set it true for debug
	Debug bool
}

// params is an application command line parameters
var params Options
