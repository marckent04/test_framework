package report

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cucumber/godog"
)

const htmlScenarioTemplateDelimiter = "<!--SCENARIO_TEMPLATE-->"
const htmlStepTemplateDelimiter = "<!--STEP_TEMPLATE-->"

var htmlTemplate string

type htmlReportFormatter struct {
}

func (r htmlReportFormatter) fillReport(template string, startDate time.Time, scenarios []Scenario) string {
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

		content += fmt.Sprintln(r.fillScenarioTemplate(sc, template))
	}

	scenarioTpl := r.getTemplateByDelimiters(template, htmlScenarioTemplateDelimiter)
	testSuiteReport := strings.Replace(template, scenarioTpl, content, 1)
	testSuiteReport = r.setTemplateVar(testSuiteReport, "EXECUTION_DATE", dateTime)
	testSuiteReport = r.setTemplateVar(testSuiteReport, "TOTAL_TESTS", strconv.Itoa(total))
	testSuiteReport = r.setTemplateVar(testSuiteReport, "SUCCEEDED_TESTS", strconv.Itoa(succeed))
	testSuiteReport = r.setTemplateVar(testSuiteReport, "FAILED_TESTS", strconv.Itoa(failed))

	const totalRate = 100
	successRate := succeed * totalRate / total
	testSuiteReport = r.setTemplateVar(testSuiteReport, "SUCCESS_RATE", strconv.Itoa(successRate))

	return testSuiteReport
}

func (r htmlReportFormatter) getTemplateByDelimiters(template, delimiter string) string {
	_, stepTemplate, _ := strings.Cut(template, delimiter)
	stepTemplate, _, _ = strings.Cut(stepTemplate, delimiter)
	return strings.TrimSpace(stepTemplate)
}

func (r htmlReportFormatter) fillScenarioTemplate(sc Scenario, reportTemplate string) string {
	formattedTemplate := r.getTemplateByDelimiters(reportTemplate, htmlScenarioTemplateDelimiter)
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

	stepTemplate := r.getTemplateByDelimiters(reportTemplate, htmlStepTemplateDelimiter)
	var filledSteps string
	for _, step := range sc.steps {
		filledSteps += fmt.Sprintln(r.fillStepTemplate(step, stepTemplate))
	}

	formattedTemplate = strings.Replace(formattedTemplate, stepTemplate, filledSteps, 1)
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
	template := r.getTemplate()
	content := r.fillReport(template, startDate, scenarios)
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

func (r htmlReportFormatter) getTemplate() string {
	if len(htmlTemplate) > 0 {
		return htmlTemplate
	}

	file, err := os.ReadFile("report.template.html")
	if err != nil {
		panic(err)
	}

	htmlTemplate = string(file)

	return htmlTemplate
}
