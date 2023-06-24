package pm

import (
	"github.com/rs/zerolog/log"
	"os/exec"
)

type PackageManager interface {
	Install(pkg ...string) *exec.Cmd
	Uninstall(pkg ...string) *exec.Cmd
	Update() *exec.Cmd
}

type basemanager string

func (b basemanager) execute(args ...string) *exec.Cmd {
	return exec.Command(string(b), args...)
}

func MustManager() PackageManager {
	for _, p := range []basemanager{"apt", "pacman"} {
		switch p {
		case "apt":
			if _, err := exec.LookPath(string(p)); err == nil {
				return &apt{p}
			}

		case "pacman":
			if _, err := exec.LookPath(string(p)); err == nil {
				return &pacman{p}
			}

		default:
			log.Fatal().Msg("No package manager found")
		}
	}

	return nil
}
