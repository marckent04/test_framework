package report

type stepError struct {
	step, msg string
}

func (e stepError) Error() string {
	return e.msg
}

func NewStepError(step, errorMsg string) error {
	return stepError{
		step: step,
		msg:  errorMsg,
	}
}
