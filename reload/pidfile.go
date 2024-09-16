package reload

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"syscall"
)

type processFile struct {
	*os.File
}

// Close and Remove pid file
func (pf processFile) Unlock() error {
	path := pf.Name()
	if err := pf.File.Close(); err != nil {
		return err
	}
	return os.Remove(path)
}

// Open or Create specified pid file.
func LockProcess(pid_path string) (interface{ Unlock() error }, error) {
	if _, err := os.Stat(pid_path); !os.IsNotExist(err) {
		raw, err := os.ReadFile(pid_path)
		if err != nil {
			return nil, err
		}
		pid, err := strconv.Atoi(string(raw))
		if err != nil {
			return nil, err
		}

		if proc, err := os.FindProcess(pid); err == nil && !errors.Is(proc.Signal(syscall.Signal(0)), os.ErrProcessDone) {
			return nil, fmt.Errorf("process is running, pid=%v", pid)
		} else if err = os.Remove(pid_path); err != nil {
			return nil, err
		}
	}

	f, err := os.OpenFile(pid_path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	if _, err := f.Write([]byte(fmt.Sprint(os.Getpid()))); err != nil {
		return nil, err
	}
	return processFile{File: f}, nil
}

// Send signal @sig to Process @pidFile refers
func SendSignal(pidFile string, sig os.Signal) error {
	bd, err := os.ReadFile(pidFile)
	if err != nil {
		return err
	}
	ipid, err := strconv.Atoi(string(bd))
	if err != nil {
		return err
	}
	proc, err := os.FindProcess(ipid)
	if err != nil {
		return err
	}
	return proc.Signal(sig)
}
