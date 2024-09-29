package common

type FrontStepDefinition = any
type FrontStepDefinitionClosure = func(*Context) FrontStepDefinition

type FrontStep struct {
	Sentences  []string
	Definition FrontStepDefinitionClosure
}
