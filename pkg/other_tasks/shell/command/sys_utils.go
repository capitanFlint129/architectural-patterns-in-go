package command

import (
	"os"
	"syscall"

	gp_ps "github.com/mitchellh/go-ps"
)

var getwd = func() (string, error) {
	return os.Getwd()
}

var chdir = func(path string) error {
	return os.Chdir(path)
}

type process struct {
	pid        int
	executable string
}

var ps = func() ([]process, error) {
	processes, err := gp_ps.Processes()
	if err != nil {
		return nil, err
	}
	resultProcesses := make([]process, len(processes))
	for i, p := range processes {
		resultProcesses[i] = process{
			pid:        p.Pid(),
			executable: p.Executable(),
		}
	}
	return resultProcesses, nil
}

var kill = func(pid int) error {
	return syscall.Kill(pid, syscall.SIGKILL)
}

var fork = func(executable string, args []string) (int, error) {
	return syscall.ForkExec(executable, args, nil)
}

var exec = func(executable string, args []string) error {
	return syscall.Exec(executable, args, nil)
}
