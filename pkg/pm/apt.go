package pm

import "os/exec"

type apt struct {
	basemanager
}

func (a *apt) Install(pkg ...string) *exec.Cmd {
	return a.execute(append([]string{"install", "-y"}, pkg...)...)
}

func (a *apt) Uninstall(pkg ...string) *exec.Cmd {
	return a.execute(append([]string{"remove", "-y"}, pkg...)...)
}

func (a *apt) Update() *exec.Cmd {
	return a.execute("update", "-y")
}
