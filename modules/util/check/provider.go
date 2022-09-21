package check

import "github.com/google/wire"

var CheckerWireSet = wire.NewSet(
	wire.Bind(new(IChecker), new(Checker)),
	CheckerProvider,
)

func CheckerProvider() Checker {
	return Checker{}
}
