package navigation

import (
	"testflowkit/internal/steps_definitions/core"
)

type navigation struct {
}

func GetSteps() []core.TestStep {
	handlers := navigation{}

	return []core.TestStep{
		handlers.iShouldBeNavigatedToPage(),
		handlers.iNavigateToPage(),
		handlers.iOpenNewBrowserTab(),
		handlers.iOpenNewPrivateBrowserTab(),
	}
}
