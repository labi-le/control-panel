package utils

import (
	"github.com/ShinyTrinkets/overseer"
	"io"
	"time"
)

func init() {
	overseer.SetupLogBuilder(func(name string) overseer.Logger {
		return Log()
	})
}

func ManageProc(cmd *overseer.Cmd, over *overseer.Overseer, w io.Writer) error {
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

func MonitorState(over *overseer.Overseer, fn func(state *overseer.ProcessJSON) string) {
	status := make(chan *overseer.ProcessJSON)
	over.Watch(status)

	go func() {
		for state := range status {
			Log().Infof("%v\n", state)
			fn(state)
		}
	}()
}
