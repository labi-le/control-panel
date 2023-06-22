package limiter

import (
	"github.com/korovkin/limiter"
)

var ErrClosed = limiter.ErrorClosed

// ConcurrencyLimiter is a wrapper for limiter.ConcurrencyLimiter
// with WaitAndClose which no panic on closed channel
type ConcurrencyLimiter struct {
	*limiter.ConcurrencyLimiter
}

func NewConcurrencyLimiter(limit int) *ConcurrencyLimiter {
	return &ConcurrencyLimiter{ConcurrencyLimiter: limiter.NewConcurrencyLimiter(limit)}
}

func (l *ConcurrencyLimiter) WaitAndClose() (err error) {
	// recover closed channel
	defer func() {
		if p := recover(); p != nil {
			err = ErrClosed
		}
	}()

	return l.ConcurrencyLimiter.WaitAndClose()
}
