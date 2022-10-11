package mootex

import (
	"fmt"
	"time"
)

// yes i know about the log package

var (
	level = 3
	label = "[Mootex] "
)

const (
	Info int = 0
	Warn     = 1
	Error    = 2
	Debug    = 3
)

func SetLevel(l int) {
	level = l
}

func SetLabel(l string) {
	if l != "" {
		label = fmt.Sprintf("%s ", l)
	} else {
		label = ""
	}
}

func LogInfo(format string, a ...interface{}) {
	if level >= Info {
		fmt.Printf("%s%s INFO  %s\n", label, time.Now().Format("01-02-2006 15:04:05"), fmt.Sprintf(format, a))
	}
}

func LogInfoln(line string) {
	if level >= Info {
		fmt.Printf("%s%s INFO  %s\n", label, time.Now().Format("01-02-2006 15:04:05"), line)
	}
}

func LogWarn(format string, a ...interface{}) {
	if level >= Warn {
		fmt.Printf("%s%s WARN  %s\n", label, time.Now().Format("01-02-2006 15:04:05"), fmt.Sprintf(format, a))
	}
}

func LogWarnln(line string) {
	if level >= Warn {
		fmt.Printf("%s%s WARN  %s\n", label, time.Now().Format("01-02-2006 15:04:05"), line)
	}
}

func LogError(format string, a ...interface{}) {
	if level >= Error {
		fmt.Printf("%s%s ERROR %s\n", label, time.Now().Format("01-02-2006 15:04:05"), fmt.Sprintf(format, a))
	}
}

func LogErrorln(line string) {
	if level >= Error {
		fmt.Printf("%s%s ERROR %s\n", label, time.Now().Format("01-02-2006 15:04:05"), line)
	}
}

func LogDebug(format string, a ...interface{}) {
	if level >= Debug {
		fmt.Printf("%s%s DEBUG %s\n", label, time.Now().Format("01-02-2006 15:04:05"), fmt.Sprintf(format, a))
	}
}

func LogDebugln(line string) {
	if level >= Debug {
		fmt.Printf("%s%s DEBUG %s\n", label, time.Now().Format("01-02-2006 15:04:05"), line)
	}
}
