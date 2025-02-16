package config

import (
	"etoolse/pkg/logger"
)

const defaultCliConfigPath = "cliConfig.yml"
const appDetailsYamlPath = "$.application"
const testingYamlPath = "$.configuration"
const reportingYamlPath = "$.reporting"

type cliConfig struct {
	reportingConfig
	AppName         string `yaml:"app_name"`
	AppDescription  string `yaml:"app_description,omitempty"`
	Timeout         string `yaml:"timeout"`
	SlowMotion      string `yaml:"slowMotion"`
	GherkinLocation string `yaml:"gherkin_location"`
}

func (c *cliConfig) init(content string) {
	appDetails := c.getAppDetailsConfig(content)
	testing := c.getTestingConfig(content)
	reporting := c.getReportingConfig(content)

	c.AppName = appDetails.AppName
	c.AppDescription = appDetails.AppDescription
	c.Timeout = testing.Timeout
	c.SlowMotion = testing.SlowMotion
	c.GherkinLocation = testing.GherkinLocation
	c.reportingConfig = reporting
}

func (c *cliConfig) getAppDetailsConfig(file string) appDetailsConfig {
	config := appDetailsConfig{}
	err := getConfig(file, appDetailsYamlPath, &config)
	if err != nil {
		return config
	}

	return config
}

func (c *cliConfig) getTestingConfig(file string) testingConfig {
	config := testingConfig{}
	err := getConfig(file, testingYamlPath, &config)
	if err != nil {
		logger.Fatal("testing config getting failed", err)
	}

	return config
}

func (c *cliConfig) getReportingConfig(file string) reportingConfig {
	config := reportingConfig{}
	err := getConfig(file, reportingYamlPath, &config)
	if err != nil {
		logger.Fatal("reporting config getting failed", err)
	}
	return config
}
