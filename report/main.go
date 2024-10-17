package report

type report struct {
	scenarios []Scenario
}

func (r *report) AddScenario(scenario Scenario) {
	r.scenarios = append(r.scenarios, scenario)
}
