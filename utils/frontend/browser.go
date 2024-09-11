package Browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"log"
	"os"
)

var browser *rod.Browser

func GetInstance() *rod.Browser {
	if browser == nil {
		CreateInstance()
	}
	return browser
}

func CreateInstance() {
	path, _ := launcher.LookPath()
	if os.Getenv("DEBUG") != "true" {
		u := launcher.New().Bin(path).
			Headless(false).
			MustLaunch()
		browser = rod.New().ControlURL(u).MustConnect()
	} else {
		browser = rod.New().ControlURL(launcher.New().Bin(path).MustLaunch()).MustConnect()
	}
	log.Println("browser instance created")
}
