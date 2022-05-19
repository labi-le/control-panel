package pkg

import (
	"os"
	"os/signal"
	"testing"
)

func TestTerm(t *testing.T) {
	term := NewTerminal("wget test")
	err := term.Exec(
		"wget", []string{"https://speed.hetzner.de/100MB.bin", "-O", "/tmp/100MB.bin"},
		func(b []byte) {},
	)

	defer os.Remove("/tmp/100MB.bin")

	if err != nil {
		t.Error(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	println("Interrupt signal received. Cleanup...")
	os.Remove("/tmp/100MB.bin")
}
