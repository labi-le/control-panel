package internal

import (
	"os/exec"
)

// TerminalParam struct for terminal
type TerminalParam struct {
	Utility   string   `json:"utility"`
	Arguments []string `json:"arguments"`
}

// RunTerminal command in terminal
func RunTerminal(param TerminalParam) (string, error) {
	cmd := exec.Command(param.Utility, param.Arguments...) //nolint:gosec
	stdout, err := cmd.Output()

	if err != nil {
		return string(stdout), err
	}

	return string(stdout), nil
}
