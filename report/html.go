package report

import "fmt"

func (r *report) HTMLBuildContent() (content string) {
	const stepTemplate = `<p style="color: %s;">%s</p>`
	for _, scenario := range r.scenarios {
		title := r.htmlWrapIn("h2", scenario.title)

		var stepsContent string
		for _, step := range scenario.steps {
			color := "green"
			if step == scenario.error.step {
				color = "red"
			}
			stepsContent += fmt.Sprintf(stepTemplate, color, step)
		}

		stepsContent = r.htmlWrapIn("div", stepsContent)
		content += r.htmlWrapIn("div", fmt.Sprintf("%s%s", title, stepsContent))
	}

	return
}

func (r *report) htmlWrapIn(selector, content string) string {
	return fmt.Sprintf("<%s>%s</%s>", selector, content, selector)
}
