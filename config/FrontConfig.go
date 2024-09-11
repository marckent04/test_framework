package config

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"log"
	"os"
	"strings"
)

var yamlContent string

func InitializeFrontConfig() {
	file, err := os.ReadFile("frontend.yml")
	if err != nil {
		log.Fatal("frontend config file not found")
	}
	yamlContent = string(file)
}

func GetPageUrl(page string) (pageUrl string, err error) {
	page = wildcardToKey(page)

	path, err := yaml.PathString(fmt.Sprintf("$.global.pages.%s", page))
	if err != nil {
		return
	}

	err = path.Read(strings.NewReader(yamlContent), &pageUrl)
	return
}

func GetElementSelectors(label string) (selectors []string) {
	label = wildcardToKey(label)

	path, err := yaml.PathString(fmt.Sprintf("$.global.elements.%s", label))
	if err != nil {
		return
	}

	err = path.Read(strings.NewReader(yamlContent), &selectors)
	return
}

func wildcardToKey(label string) string {
	return strings.ToLower(strings.ReplaceAll(label, " ", "_"))
}
