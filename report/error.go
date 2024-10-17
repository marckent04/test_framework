package report

type StepError struct {
	step, msg string
}

func (e StepError) Error() string {
	return e.msg
}

func NewStepError(step, errorMsg string) error {
	return StepError{
		step: step,
		msg:  errorMsg,
	}
}
