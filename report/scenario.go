package report

type scenario struct {
	title string
	steps []string
	error error
}

func newScenario(title string, steps []string, err StepError) scenario {
	return scenario{
		title: title,
		steps: steps,
		error: err,
	}
}
