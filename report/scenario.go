package report

type Scenario struct {
	title string
	steps []string
	error stepError
}
