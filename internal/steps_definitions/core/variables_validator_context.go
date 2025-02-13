package core

import (
	"etoolse/internal/config/testsconfig"
	"fmt"
	"slices"
	"strings"
	"sync"
)

var addLock sync.Mutex

type ValidatorContext struct {
	missingPageErrors    []string
	missingElementErrors []string
}

func (vc *ValidatorContext) AddMissingPage(label string) {
	key := testsconfig.GetLabelKey(label)

	if slices.Contains(vc.missingPageErrors, key) {
		return
	}
	vc.missingPageErrors = append(vc.missingPageErrors, key)
}

func (vc *ValidatorContext) AddMissingElement(label string) {
	key := testsconfig.GetLabelKey(label)
	if slices.Contains(vc.missingElementErrors, key) {
		return
	}
	vc.missingElementErrors = append(vc.missingElementErrors, key)
}

func (vc *ValidatorContext) HasErrors() bool {
	return len(vc.missingPageErrors) > 0 || len(vc.missingElementErrors) > 0
}

func (vc *ValidatorContext) GetErrors() []string {
	return append(vc.missingPageErrors, vc.missingElementErrors...)
}

func (vc *ValidatorContext) GetMissingPages() []string {
	return vc.missingPageErrors
}

func (vc *ValidatorContext) GetMissingElements() []string {
	return vc.missingElementErrors
}

func (vc *ValidatorContext) GetElementsErrorsFormatted() string {
	lines := []string{
		"global:",
		"\telements:",
	}

	const elementFormat = "\t\t%s:\n-<missing-selector>"
	
	for _, element := range vc.missingElementErrors {
		lines = append(lines, fmt.Sprintf(elementFormat, element))
	}
	return strings.Join(lines, "\n")
}

func (vc *ValidatorContext) AddValidationErrors(errors ValidationErrors) {
	addLock.Lock()
	defer addLock.Unlock()

	vc.missingPageErrors = append(vc.missingPageErrors, errors.missingPages...)
	vc.missingElementErrors = append(vc.missingElementErrors, errors.missingElements...)
}
