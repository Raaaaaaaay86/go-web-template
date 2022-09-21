package check

import (
	"errors"

	"github.com/google/wire"
)

type IChecker interface {
	String(strings []string, predicate func(str string) bool) error
}

type Checker struct{}

var CheckerWireSet = wire.NewSet(
	wire.Bind(new(IChecker), new(Checker)),
	CheckerProvider,
)

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
