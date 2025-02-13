package core

type TestStep interface {
	GetSentences() []string
	GetDefinition(*TestSuiteContext) any
	Validate(*ValidatorContext) any
}
