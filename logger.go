package golog

var (
	level      Level  = TRACE
	colorLevel bool   = false
	timeFormat string = "02/01/2006 - 15:04:05"
)

func T(m ...string) {

	lvl := TRACE

	if lvl >= level {
		print(YELLOW, lvl, m...)
	}
}
func D(m ...string) {
	lvl := DEBUG

	if lvl >= level {
		print(GREEN, lvl, m...)
	}
}
func I(m ...string) {
	lvl := INFO

	if lvl >= level {
		print(BLUE, lvl, m...)
	}
}
func W(m ...string) {
	lvl := WARNING

	if lvl >= level {
		print(ORANGE, lvl, m...)
	}
}
func E(m ...string) {
	lvl := ERROR

	if lvl >= level {
		print(RED, lvl, m...)
	}
}
func F(m ...string) {
	lvl := FATAL

	if lvl >= level {
		print(RED, lvl, m...)
	}
}
func P(m ...string) {
	lvl := PANIC

	if lvl >= level {
		print(RED, lvl, m...)
	}
}
