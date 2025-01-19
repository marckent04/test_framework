package navigation

import (
	"etoolse/frontend/common"
)

type navigation struct {
}

func GetSteps() []common.FrontStep {
	handlers := navigation{}

	return []common.FrontStep{
		handlers.iShouldBeNavigatedToPage(),
		handlers.iNavigateToPage(),
		handlers.iOpenNewBrowserTab(),
		handlers.iOpenNewPrivateBrowserTab(),
	}
}
