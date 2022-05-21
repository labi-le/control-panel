package terminal

import (
	"log"
	"os"
	"os/exec"
)

type Terminal struct {
	PID     int
	Title   string
	Started bool
	Ended   bool
}

func NewTerminal(title string) *Terminal {
	return &Terminal{
		PID:     os.Getpid(),
		Title:   title,
		Started: false,
		Ended:   false,
	}
}

func (t *Terminal) Exec(name string, arg []string, progress func(o *Output)) error {
	cmd := exec.Command(name, arg...)
	log.Println(cmd.String())
	//  both the error output and standard output of the command are connected to the same pipe
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	t.Started = true

	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)

		progress(NewOutput(tmp))

		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		return err
	}

	t.Ended = true
	return nil
}
