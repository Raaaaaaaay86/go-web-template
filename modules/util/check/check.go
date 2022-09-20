package check

import "errors"

type IChecker interface {
	String(values []string, predicate func(str string))
}

type Checker struct{}

func CheckerProvider() Checker {
	return Checker{}
}

func (c Checker) String(strings []string, predicate func(str string) bool) error {
	for _, str := range strings {
		if !predicate(str) {
			return errors.New("input value is not as expectation")
		}
	}

	return nil
}
