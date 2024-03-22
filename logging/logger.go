package logging

type Logger interface {

	// Verbose .
	Verbose(message string, args ...any)

	// Debug .
	Debug(message string, args ...any)

	// Info .
	Info(message string, args ...any)

	// Warning .
	Warning(message string, args ...any)

	// Error .
	Error(message string, err error, args ...any)

	// Fatal .
	Fatal(message string, args ...any)

	// Destroy .
	Destroy()
}

var LevelConstants = levelEnum{
	VERBOSE: &Level{ordinal: 0, name: "VERBOSE"},
	DEBUG:   &Level{ordinal: 1, name: "DEBUG"},
	INFO:    &Level{ordinal: 2, name: "INFO"},
	WARNING: &Level{ordinal: 3, name: "WARNING"},
	ERROR:   &Level{ordinal: 4, name: "ERROR"},
	FATAL:   &Level{ordinal: 5, name: "FATAL"},
}

type levelEnum struct {
	VERBOSE *Level
	DEBUG   *Level
	INFO    *Level
	WARNING *Level
	ERROR   *Level
	FATAL   *Level
}

type Level struct {
	ordinal int
	name    string
}
