package common

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"os"
)

func CreateInstance() *rod.Browser {
	path, _ := launcher.LookPath()
	if os.Getenv("DEBUG") == "true" {
		u := launcher.New().Bin(path).
			Headless(false).
			MustLaunch()
		return rod.New().ControlURL(u).MustConnect()
	}
	return rod.New().ControlURL(launcher.New().Bin(path).MustLaunch()).MustConnect()
}
