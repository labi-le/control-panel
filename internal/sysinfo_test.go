package internal

import (
	"testing"
)

func TestGetCPUInfo(t *testing.T) {
	_, err := GetCPUInfo()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestGetCPUAvg(t *testing.T) {
	_, err := GetAvg()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestGetCPUTimes(t *testing.T) {
	_, err := GetCPUTimes()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
