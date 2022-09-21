package check

import "errors"

type IChecker interface {
	String(strings []string, predicate func(str string) bool) error
}

type Checker struct{}

func (c Checker) String(strings []string, predicate func(str string) bool) error {
	for _, str := range strings {
		if !predicate(str) {
			return errors.New("input value is not as expectation")
		}
	}

	return nil
}
