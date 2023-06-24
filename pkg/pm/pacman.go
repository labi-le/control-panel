package pm

import "os/exec"

type pacman struct {
	basemanager
}

func (p *pacman) Install(pkg ...string) *exec.Cmd {
	return p.execute(append([]string{"-S", "--noconfirm"}, pkg...)...)
}

func (p *pacman) Uninstall(pkg ...string) *exec.Cmd {
	return p.execute(append([]string{"-Rs", "--noconfirm"}, pkg...)...)
}

func (p *pacman) Update() *exec.Cmd {
	return p.execute("-Syu", "--noconfirm")
}
