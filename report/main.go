package report

import "time"

type Report struct {
	scenarios []scenario
	startDate time.Time
	duration  time.Duration
}

func (r *Report) AddScenario(title string,
	steps []string,
	error StepError) {
	r.scenarios = append(r.scenarios, newScenario(title, steps, error))
}

func (r *Report) SetStartDate(date time.Time) {
	r.startDate = date
}

func New() Report {
	return Report{}
}
