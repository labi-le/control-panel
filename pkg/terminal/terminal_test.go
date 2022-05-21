package terminal

import (
	"os"
	"testing"
)

const TestFile = "https://file-examples.com/storage/fe79bebc316288cc495169e/2017/04/file_example_MP4_480_1_5MG.mp4"
const TestFileName = "/tmp/testqqq.mp4"

func TestTerm(t *testing.T) {
	term := NewTerminal("wget test")
	err := term.Exec(
		"wget", []string{TestFile, "-O", TestFileName},
		func(o *Output) {},
	)

	defer os.Remove(TestFileName)

	if err != nil {
		t.Error(err)
	}
}
