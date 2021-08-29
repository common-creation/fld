package utils

import (
	"bufio"
	"github.com/common-creation/fld/i18n"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

type (
	Command struct {
		Cmd  string
		Args []string
	}

	wrapExec struct {
		cmd *exec.Cmd
		logPrefix string
	}
)

func HasCommands(commands []Command) bool {
	ch := make(chan bool)
	for _, v := range commands {
		go func(c Command) {
			defer func() {
				err := recover()
				if err != nil {
					ch <- false
				}
			}()
			code, err := NewCommand("", c.Cmd, c.Args...).Silent().Exec()
			ok := code == 0 && err == nil
			if !ok {
				Errorln(i18n.T("command.error", map[string]interface{}{
					"command": c.Cmd,
				}))
			}
			ch <- ok
		}(v)
	}
	result := true
	for i := 0; i < len(commands); i++ {
		if !<-ch {
			result = false
		}
	}
	return result
}

func NewCommand(logPrefix string, command string, args ...string) *wrapExec {
	this := &wrapExec{
		cmd: exec.Command(command, args...),
		logPrefix: logPrefix,
	}
	this.cmd.Dir, _ = filepath.Split(command)
	this.cmd.Env = os.Environ()
	this.cmd.Stdin = os.Stdin
	this.cmd.Stdout = os.Stdout
	this.cmd.Stderr = os.Stderr

	return this
}

func (this *wrapExec) Cwd(path string) *wrapExec {
	dir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	this.cmd.Dir = dir
	Debugln("change cwd:", this.cmd.Dir)
	return this
}

func (this *wrapExec) Silent() *wrapExec {
	this.cmd.Stdout = nil
	this.cmd.Stderr = nil
	this.cmd.Stdin = nil
	Debugln("disconnect IO", "|", "command:", this.cmd.Path, "args:", this.cmd.Args, "cwd:", this.cmd.Dir)
	return this
}

func (this *wrapExec) Exec() (int, error) {
	Debugln("execute", "|", "command:", this.cmd.Path, "args:", this.cmd.Args, "cwd:", this.cmd.Dir)
	defer Debugln("execute done", "|", "command:", this.cmd.Path, "args:", this.cmd.Args, "cwd:", this.cmd.Dir)
	return wrapResult(this.cmd.Run())
}

func (this *wrapExec) ExecAsync() (int, error) {
	Debugln("execute", "|", "command:", this.cmd.Path, "args:", this.cmd.Args, "cwd:", this.cmd.Dir)
	defer Debugln("execute done", "|", "command:", this.cmd.Path, "args:", this.cmd.Args, "cwd:", this.cmd.Dir)

	this.Silent()

	this.cmd.Stdin = os.Stdin
	stdout, _ := this.cmd.StdoutPipe()
	stderr, _ := this.cmd.StderrPipe()
	os := bufio.NewScanner(stdout)
	es := bufio.NewScanner(stderr)
	go func() {
		l := NewSyncLogger(this.logPrefix, false)
		for os.Scan() {
			l.Write(os.Bytes())
		}
	}()
	go func() {
		l := NewSyncLogger(this.logPrefix, true)
		for es.Scan() {
			l.Write(es.Bytes())
		}
	}()

	this.cmd.Start()
	return wrapResult(this.cmd.Wait())
}

func (this *wrapExec) MustExec() int {
	code, err := this.Exec()
	if err != nil {
		panic(err)
	}
	return code
}

func (this *wrapExec) ExtendArgs(args ...string) {
	this.cmd.Args = append(this.cmd.Args, args...)
}

func wrapResult(err error) (int, error) {
	var code int
	if err != nil {
		if err2, ok := err.(*exec.ExitError); ok {
			if s, ok := err2.Sys().(syscall.WaitStatus); ok {
				err = nil
				code = s.ExitStatus()
			}
		}
	}
	return code, err
}
