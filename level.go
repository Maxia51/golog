package golog

type Level int

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
	PANIC
)

func (level Level) ToString() string {

	switch level {
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	case PANIC:
		return "PANIC"
	}

	return "?"

}
