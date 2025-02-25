package form

import (
	"etoolse/internal/config/testsconfig"
	"etoolse/internal/steps_definitions/core"
	"etoolse/internal/utils"
	"fmt"
)

func (s steps) dropdownHaveValuesSelected() core.TestStep {
	formatVar := func(label string) string {
		return fmt.Sprintf("%s_dropdown", label)
	}
	return core.NewStepWithTwoVariables(
		[]string{`^the {string} dropdown should have "{string}" selected$`},
		func(ctx *core.TestSuiteContext) func(string, string) error {
			return func(dropdownId, optionLabels string) error {
				selector, err := testsconfig.GetHTMLElementSelectors(formatVar(dropdownId))
				if err != nil {
					return err
				}

				labels := utils.String{}.SplitAndTrim(optionLabels, ",")

				result := ctx.GetCurrentPage().ExecuteJS(`(selector, labels) => {
					const selectedOpts = Array.from(document.querySelector(selector).selectedOptions).map(opt => opt.label)
					return labels.every(label => selectedOpts.includes(label))
				}`, selector, labels)

				if result == "true" {
					return nil
				}
				return fmt.Errorf("%s value is not selected in %s dropdown", optionLabels, dropdownId)
			}
		},
		func(dropdownId, _ string) core.ValidationErrors {
			vErr := core.ValidationErrors{}
			label := formatVar(dropdownId)
			if !testsconfig.IsElementDefined(label) {
				vErr.AddMissingElement(label)
			}

			return vErr
		},
		core.StepDocumenation{
			Description: "Check if the dropdown has the specified values selected",
			Variables: []core.StepVariable{
				{Name: "dropdown", Type: "string"},
				{Name: "values", Type: "string"},
			},
			Example: `the "country" dropdown should have "USA,Canada" selected`,
		},
	)
}
