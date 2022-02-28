package interpreter

import (
	"os"
	"os/exec"
	"testing"
)

const GOLOX_TEST_RUNFILE = "GOLOX_TEST_RUNFILE"

func Test_RunFile_FileNotFound(t *testing.T) {
	if os.Getenv(GOLOX_TEST_RUNFILE) == "true" {
		RunFile("i_do_not_exist.lox")
		return
	}

	command := exec.Command(os.Args[0], "-test.run=Test_RunFile_FileNotFound")
	command.Env = append(os.Environ(), GOLOX_TEST_RUNFILE+"=true")

	err, ok := command.Run().(*exec.ExitError)

	if !ok {
		t.Fatalf("error running test")
	}

	if err.ExitCode() != 74 {
		t.Fatalf("expect exit code 66")
	}
}
