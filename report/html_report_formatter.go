package report

import (
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/cucumber/godog"
)

const reportTemplateKey, scenarioTemplateKey, stepTemplateKey = "report", "scenario", "step"

var htmlTemplates = map[string]string{
	reportTemplateKey:   "",
	scenarioTemplateKey: "",
	stepTemplateKey:     "",
}

type htmlReportFormatter struct {
}

func (r htmlReportFormatter) fillReport(startDate time.Time, scenarios []Scenario, reportTemplate, scenarioTemplate, stepTemplate string) string {
	year, month, day := startDate.Date()
	dateTime := fmt.Sprintf("%d-%d-%d at %d:%d", month, day, year, startDate.Hour(), startDate.Minute())

	total := len(scenarios)
	var content string
	var succeed, failed int
	for _, sc := range scenarios {
		if len(sc.err) > 0 {
			failed++
		} else {
			succeed++
		}

		content += fmt.Sprintln(r.fillScenarioTemplate(sc, scenarioTemplate, stepTemplate))
	}

	testSuiteReport := r.setTemplateVar(reportTemplate, "SCENARIOS", content)
	testSuiteReport = r.setTemplateVar(testSuiteReport, "EXECUTION_DATE", dateTime)
	testSuiteReport = r.setTemplateVar(testSuiteReport, "TOTAL_TESTS", strconv.Itoa(total))
	testSuiteReport = r.setTemplateVar(testSuiteReport, "SUCCEEDED_TESTS", strconv.Itoa(succeed))
	testSuiteReport = r.setTemplateVar(testSuiteReport, "FAILED_TESTS", strconv.Itoa(failed))

	const totalRate = 100
	successRate := succeed * totalRate / total
	testSuiteReport = r.setTemplateVar(testSuiteReport, "SUCCESS_RATE", strconv.Itoa(successRate))

	return testSuiteReport
}

func (r htmlReportFormatter) fillScenarioTemplate(sc Scenario, scenarioTemplate, stepTemplate string) string {
	formattedTemplate := scenarioTemplate
	formattedTemplate = r.setTemplateVar(formattedTemplate, "SCENARIO_NAME", sc.title)

	duration := sc.duration.Seconds()
	duration = math.Max(math.Ceil(duration), duration)
	formattedTemplate = r.setTemplateVar(formattedTemplate, "SCENARIO_DURATION", fmt.Sprintf("%vs", duration))
	result, color := "FAILED", "red"
	if sc.err == "" {
		result, color, sc.err = "SUCCEEDED", "green", "-"
	}
	formattedTemplate = r.setTemplateVar(formattedTemplate, "SCENARIO_RESULT", result)
	formattedTemplate = r.setTemplateVar(formattedTemplate, "SCENARIO_STATUS_COLOR", color)
	formattedTemplate = r.setTemplateVar(formattedTemplate, "SCENARIO_ERROR_MESSAGE", sc.err)

	var filledSteps string
	for _, step := range sc.steps {
		filledSteps += fmt.Sprintln(r.fillStepTemplate(step, stepTemplate))
	}

	formattedTemplate = r.setTemplateVar(formattedTemplate, "STEPS", filledSteps)
	return formattedTemplate
}

func (r htmlReportFormatter) fillStepTemplate(step Step, template string) string {
	getColor := func(status godog.StepResultStatus) string {
		if status == 0 {
			return "green"
		}

		if status == 1 {
			return "red"
		}

		return "gray"
	}

	tpl := r.setTemplateVar(template, "STEP_TITLE", step.title)
	tpl = r.setTemplateVar(tpl, "STEP_STATUS", step.status.String())
	tpl = r.setTemplateVar(tpl, "STEP_STATUS_COLOR", getColor(step.status))
	return strings.TrimSpace(tpl)
}

func (r htmlReportFormatter) setTemplateVar(template, variableName, value string) string {
	variable := fmt.Sprintf("{{ %s }}", variableName)
	return strings.ReplaceAll(template, variable, value)
}

func (r htmlReportFormatter) WriteReport(startDate time.Time, scenarios []Scenario) {
	report, scenario, step := r.getTemplates()
	content := r.fillReport(startDate, scenarios, report, scenario, step)
	file, err := os.Create("report.html")
	if err != nil {
		log.Panicf("cannot create report file in this folder ( %s )\n", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Panicf("error when report filling ( %s )", err)
	}
}

func (r htmlReportFormatter) getTemplates() (report, scenario, step string) {
	for name, content := range htmlTemplates {
		if len(content) > 0 {
			continue
		}

		templatePath := path.Join("report", "templates", fmt.Sprintf("%s.template.html", name))
		file, err := os.ReadFile(templatePath)
		if err != nil {
			log.Printf("%s template not found\n", name)
			panic(err)
		}
		htmlTemplates[name] = string(file)
	}

	return htmlTemplates[reportTemplateKey], htmlTemplates[scenarioTemplateKey], htmlTemplates[stepTemplateKey]
}
