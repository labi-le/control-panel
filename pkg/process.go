package pkg

import (
	logger "github.com/ShinyTrinkets/meta-logger"
	"github.com/ShinyTrinkets/overseer"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

type Process struct {
	*logrus.Logger
}

func NewProcess(l *logrus.Logger) *Process {
	overseer.SetupLogBuilder(func(name string) logger.Logger {
		return NewLogger(l)
	})

	return &Process{l}
}

func (p *Process) ManageProc(cmd *overseer.Cmd, over *overseer.Overseer, w io.Writer) error {
	go over.SuperviseAll()

	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case line := <-cmd.Stdout:
			_, err := w.Write([]byte(line))
			if err != nil {
				return err
			}
		case line := <-cmd.Stderr:
			_, err := w.Write([]byte(line))
			if err != nil {
				return err
			}
		case <-ticker.C:
			if !over.IsRunning() {
				break //terminate go routine
			}
		}
	}
}

func (p *Process) MonitorState(over *overseer.Overseer, fn func(state *overseer.ProcessJSON) string) {
	status := make(chan *overseer.ProcessJSON)
	over.Watch(status)

	go func() {
		for state := range status {
			p.Infof("%v\n", state)
			fn(state)
		}
	}()
}
