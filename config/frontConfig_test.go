package config

import (
	"log"
	"os"
	"slices"
	"testing"
)

func setup(t *testing.T) {
	bytes, fileErr := os.ReadFile("frontend.test.yml")
	if fileErr != nil {
		t.Error(fileErr)
	}
	yamlContent = string(bytes)
}

func TestGetPageUrl(t *testing.T) {
	setup(t)

	url, err := GetPageUrl("home")
	if err != nil {
		t.Error(err)
	}

	expected := "https://google.com"
	if url != expected {
		t.Errorf("url got %s / url expected: %s", url, expected)
	}
}

func TestGetElementSelectors(t *testing.T) {
	setup(t)

	log.Println(GetElementSelectors("login"))
	log.Println(slices.Equal([]string{"#login", "l"}, []string{"l", "#login"}))
	expected := []string{"#login", ".login", "button .login"}
	if selectors := GetElementSelectors("login"); !slices.Equal(selectors, expected) {
		t.Errorf("error ")
	}
}
