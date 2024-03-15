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

var Level = levelEnum{
	VERBOSE: &level{ordinal: 0, name: "VERBOSE"},
	DEBUG:   &level{ordinal: 1, name: "DEBUG"},
	INFO:    &level{ordinal: 2, name: "INFO"},
	WARNING: &level{ordinal: 3, name: "WARNING"},
	ERROR:   &level{ordinal: 4, name: "ERROR"},
	FATAL:   &level{ordinal: 5, name: "FATAL"},
}

type levelEnum struct {
	VERBOSE *level
	DEBUG   *level
	INFO    *level
	WARNING *level
	ERROR   *level
	FATAL   *level
}

type level struct {
	ordinal int
	name    string
}
