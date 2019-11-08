package golog

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func getCaller() string {
	pc, _, _, ok := runtime.Caller(3)
	details := runtime.FuncForPC(pc)

	if ok && details != nil {

		split := strings.Split(details.Name(), "/")

		leng := len(split)

		return split[leng-1]
	}

	return "Uknown"
}

func print(color Color, lvl Level, m ...string) {

	if colorLevel {
		fmt.Println(color, "[", time.Now().Format(timeFormat), "]", "[", lvl.ToString(), "]", "[", getCaller(), "]", m, "\033[0m")
	} else {
		fmt.Println("[", time.Now().Format(timeFormat), "]", color, "[", lvl.ToString(), "]", "\033[0m", "[", getCaller(), "]", m)
	}

}
