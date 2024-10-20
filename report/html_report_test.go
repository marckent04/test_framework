package report

import (
	"errors"
	"html"
	"strings"
	"testing"
	"time"

	"github.com/tdewolff/minify/v2/minify"
)

func TestGetTemplateByDelimiters(t *testing.T) {
	formatter := htmlReportFormatter{}

	scenarioTpl := `
          <!--STEP_TEMPLATE-->
			<p>hello step template</p>
          <!--STEP_TEMPLATE-->
	`

	expected := "<p>hello step template</p>"

	if received := formatter.getTemplateByDelimiters(scenarioTpl, "<!--STEP_TEMPLATE-->"); received != expected {
		t.Errorf("Step template getting error\nExpected:%s\nreceived:%s", expected, received)
	}
}

func TestFormatSuccessCaseScenarioTemplate(t *testing.T) {
	formatter := htmlReportFormatter{}

	first := newScenario("first", []Step{
		{
			title:  "je",
			status: 0,
		},
		{
			title:  "fais",
			status: 0,
		},
	}, 10*time.Second, nil)

	firstTpl := formatter.fillScenarioTemplate(first, template)

	if !strings.Contains(firstTpl, strings.TrimSpace(`
		<details>
		first
		SUCCEEDED
        10s
        -
		<details>
	`)) {
		t.Errorf("Template for scenario success case malformed\n generated template: %s", firstTpl)
	}
}

func TestFormatFailedCaseScenarioTemplate(t *testing.T) {
	formatter := htmlReportFormatter{}
	failed := newScenario("connect", make([]Step, 0), 3*time.Second, errors.New("error"))

	firstTpl := formatter.fillScenarioTemplate(failed, template)

	if !strings.Contains(firstTpl, strings.TrimSpace(`
		<details>
		connect
		FAILED
        3s
        error
		<details>
	`)) {
		t.Errorf("Template for scenario failed case malformed\n generated template: %s", firstTpl)
	}
}

func TestFormatStepTemplate(t *testing.T) {
	formatter := htmlReportFormatter{}
	step := Step{title: "etape", status: 0}
	tpl := "{{ STEP_TITLE }} / {{ STEP_STATUS_COLOR }} / {{ STEP_STATUS }}"
	expected := "etape / green / passed"

	if result := formatter.fillStepTemplate(step, tpl); result != expected {
		t.Errorf("Error when step formatting\n Expected: %s\nreceived:%s", expected, result)
	}
}

func TestFormatScenario(t *testing.T) {
	t.Skipf("wizard error")
	formatter := htmlReportFormatter{}

	steps := []Step{
		{
			title:  "je",
			status: 0,
		},
		{
			title:  "fais",
			status: 1,
		},
	}

	sc := newScenario("TESTING", steps, 2*time.Second, errors.New("error"))

	scenarioTpl := `
    <!--SCENARIO_TEMPLATE-->
			<h1>{{ SCENARIO_NAME }}</h1>
			<div>
          <!--STEP_TEMPLATE-->
			<p>{{ STEP_TITLE }}</p>
          <!--STEP_TEMPLATE-->
			</div>
    <!--SCENARIO_TEMPLATE-->
`

	expected, _ := minify.HTML(`
	<h1>TESTING</h1>
	<div>
		<p>je</p>
		<p>fais</p>
	</div>
`)

	result, _ := minify.HTML(formatter.fillScenarioTemplate(sc, scenarioTpl))

	if strings.TrimSpace(html.EscapeString(result)) != strings.TrimSpace(html.EscapeString(expected)) {
		t.Errorf("Expected: %s, Received: %s", strings.TrimSpace(expected), strings.TrimSpace(result))
	}
}

const template = `
	<doc>
    <!--SCENARIO_TEMPLATE-->
    <scenario>
        <details>
		{{ SCENARIO_NAME }}
		{{ SCENARIO_RESULT }}
        {{ SCENARIO_DURATION }}
        {{ SCENARIO_ERROR_MESSAGE }}
		<details>
        <steps>
          <!--STEP_TEMPLATE-->
          <li>
			{{ STEP_TITLE }}
			{{ STEP_STATUS }}
			text-{{ STEP_STATUS_COLOR }}-500
          </li>
          <!--STEP_TEMPLATE-->
        <steps>
	<scenario>
    <!--SCENARIO_TEMPLATE-->
	</doc>
`

func newScenario(title string, steps []Step, duration time.Duration, err error) Scenario {
	sc := Scenario{
		title:    title,
		steps:    steps,
		duration: duration,
	}

	if err != nil {
		sc.err = err.Error()
	}

	return sc
}
