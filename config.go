package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
)

var (
	config Config // holds the global config

	c struct {
		logLevel string
		harmony  struct {
			api       string
			verifyssl bool
		}
	}

	configFile        = ""
	defaultConfigFile = "config.toml"
	printVersion      bool
)

func init() {
	flag.StringVar(&configFile, "configFile", "", "the config file")
	flag.BoolVar(&printVersion, "version", false, "print version and exit")

	flag.StringVar(&c.logLevel, "logLevel", "", "the level of messages to log")

	flag.StringVar(&c.harmony.api, "harmony.api", "http://harmony.dev:4774", "the url to the Harmony API")
	flag.BoolVar(&c.harmony.verifyssl, "harmony.verifyssl", true, "verify ssl connections to the harmony api")
}

// Config is the main config type
type Config struct {
	// LogLevel main application loggin level
	LogLevel string `toml:"LogLevel"`

	// Harmony is the main Harmony config
	Harmony HarmonyConfig `toml:"Harmony"`
}

// HarmonyConfig is the main Harmony config
type HarmonyConfig struct {
	// API url to the Harmony API
	API string `toml:"API"`

	// VerifySSL is wether ot not we are to verify the secure Harmony API connections
	VerifySSL bool `toml:"VerifySSL"`
}

func initConfig() error {
	if configFile == "" {
		if _, err := os.Stat(defaultConfigFile); !os.IsNotExist(err) {
			configFile = defaultConfigFile
		}
	}

	// Set defaults.
	config = Config{
		LogLevel: "info",
		Harmony: HarmonyConfig{
			API:       "http://harmony.dev:4774",
			VerifySSL: true,
		},
	}

	// Update config from the TOML configuration file.
	if configFile == "" {
		log.Info("Skipping config file parsing")
	} else {
		log.WithField("file", configFile).Info("Loading config")

		configBytes, err := ioutil.ReadFile(configFile)
		if err != nil {
			return err
		}
		_, err = toml.Decode(string(configBytes), &config)
		if err != nil {
			return err
		}
	}
	// Update config from commandline flags.
	processFlags()

	if config.LogLevel != "" {
		LogSetLevel(config.LogLevel)
	}

	return nil
}

func processFlags() {
	flag.Visit(setConfigFromFlag)
}

func setConfigFromFlag(f *flag.Flag) {
	switch f.Name {
	case "logLevel":
		config.LogLevel = c.logLevel
	case "harmony.api":
		config.Harmony.API = c.harmony.api
	case "harmony.verifyssl":
		config.Harmony.VerifySSL = c.harmony.verifyssl
	}
}
