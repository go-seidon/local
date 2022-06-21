package clock

import "time"

type defaultImplementation struct{}

func (o *defaultImplementation) Now() time.Time {
	return time.Now()
}
