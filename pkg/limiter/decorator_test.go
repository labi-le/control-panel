package limiter

import (
	"testing"
	"time"
)

func TestConcurrencyLimiter_WaitAndClose(t *testing.T) {
	cl := NewConcurrencyLimiter(2)

	go func() {
		for i := 0; i < 10; i++ {
			_, _ = cl.Execute(func() {
				dur := time.Duration(i * 100)
				time.Sleep(dur * time.Second) //nolint:durationcheck //dn
			})
		}
	}()

	if err := cl.WaitAndClose(); err != nil {
		t.Error(err)
	}

	if err := cl.WaitAndClose(); err == nil {
		t.Error("expect error")
	}
}
