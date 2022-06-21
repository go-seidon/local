/**
 * The idea with this package is whenever i want to testing I need to mock
 * time.Now() to return static time. So in the application code, I will use
 * clock.Now() instead and in unit testing I will override it with a mock.
 *
 * To achieve that the clock package should accepts an object that contains
 * real implementation
 */
package clock

import "time"

type Clock interface {
	Now() time.Time
}

/**
 * Set your own clock implementation by change this variable to your object
 */
var Implementation Clock = &defaultImplementation{}

/**
 * Now() function
 */
func Now() time.Time {
	return Implementation.Now()
}
