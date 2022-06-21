package tests

import (
	"time"

	"github.com/go-seidon/local/internal/clock"
)

/**
 * ----------------------------------------------------------------------------
 * Call this function will mock clock with stub implementation. You just need to
 * call this function in testing.
 *
 * Clock is a package that contains several functions that wrap golang time package
 * functionality such as clock.Now() that wraps time.Now(). The idea is instead of
 * application calls time package directly it should via clock then we can mock it
 * in testing.
 *
 * ----------------------------------------------------------------------------
 */
func MockClock() {
	clock.Implementation = &StubClock{}
}

/**
 * ----------------------------------------------------------------------------
 * Stub clock
 * ----------------------------------------------------------------------------
 */
type StubClock struct{}

var fixedTime = time.Now()

func (o *StubClock) Now() time.Time {
	return fixedTime
}
