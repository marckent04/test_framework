package reporters

import (
	"time"

	"github.com/cucumber/godog"
)

type Scenario struct {
	title     string
	steps     []Step
	err       string
	startDate time.Time
	duration  time.Duration
}

func (s *Scenario) AddStep(title string, status godog.StepResultStatus, duration time.Duration, err error) {
	if err != nil {
		s.err = err.Error()
	}

	s.steps = append(s.steps, Step{
		title:    title,
		status:   status,
		duration: duration,
	})
}

func (s *Scenario) SetTitle(title string) {
	s.title = title
}

func (s *Scenario) End() {
	s.duration = time.Since(s.startDate)
}

type Step struct {
	title    string
	status   godog.StepResultStatus
	duration time.Duration
}

func NewScenario() Scenario {
	return Scenario{
		startDate: time.Now(),
	}
}
