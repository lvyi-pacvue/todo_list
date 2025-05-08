package logger

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
}

type Field struct {
	Key   string
	Value any
}

func Error(err error) Field {
	return Field{
		Key:   "error",
		Value: err,
	}
}
