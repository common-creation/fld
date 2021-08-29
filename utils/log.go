package utils

import (
	"fmt"
	"github.com/ttacon/chalk"
	"os"
	"strings"
	"sync"
)

func Errorln(a ...interface{}) {
	println(chalk.Red, "error", a...)
}

func Infoln(a ...interface{}) {
	println(chalk.Green, "info ", a...)
}

func Debugln(a ...interface{}) {
	if IsDebug() {
		println(chalk.Cyan, "debug", a...)
	}
}

func IsDebug() bool {
	return os.Getenv("FLD_DEBUG") != ""
}

var mutex sync.Mutex

func Println(a ...interface{}) (n int, err error) {
	args := []interface{}{chalk.Reset}
	args = append(args, a...)

	mutex.Lock()
	defer mutex.Unlock()
	return fmt.Println(args...)
}

func PrintErrorln(a ...interface{}) (n int, err error) {
	args := []interface{}{chalk.Red}
	args = append(args, a...)
	args = append(args, chalk.Reset)

	mutex.Lock()
	defer mutex.Unlock()
	return fmt.Println(args...)
}

func println(color chalk.Color, level string, a ...interface{}) (n int, err error) {
	base := []interface{}{color, strings.ToUpper(level), chalk.Reset, "|"}
	args := append(base, a...)
	mutex.Lock()
	defer mutex.Unlock()
	return fmt.Println(args...)
}

type (
	SyncLogger struct {
		prefix   string
		isStdErr bool
	}
)

func NewSyncLogger(prefix string, isStdErr bool) *SyncLogger {
	return &SyncLogger{
		prefix:   prefix,
		isStdErr: isStdErr,
	}
}

func (sl *SyncLogger) Write(p []byte) (n int, err error) {
	str := string(p)
	slice := strings.Split(str, "\n")
	totalN := 0
	for i, _ := range slice {
		line := strings.TrimSpace(sl.prefix + " | " + slice[i])
		if sl.isStdErr {
			n, _ := PrintErrorln(line)
			totalN += n
		} else {
			n, _ := Println(line)
			totalN += n
		}
	}
	return totalN, nil
}

func (sl *SyncLogger) Close() error {
	return nil
}

func (sl *SyncLogger) Flush() error {
	return nil
}
