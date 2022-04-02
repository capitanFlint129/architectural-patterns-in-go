package command

import (
	"os"

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
