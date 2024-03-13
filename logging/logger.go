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
	VERBOSE: &level{Ordinal: 0, Name: "VERBOSE"},
	DEBUG:   &level{Ordinal: 1, Name: "DEBUG"},
	INFO:    &level{Ordinal: 2, Name: "INFO"},
	WARNING: &level{Ordinal: 3, Name: "WARNING"},
	ERROR:   &level{Ordinal: 4, Name: "ERROR"},
	FATAL:   &level{Ordinal: 5, Name: "FATAL"},
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
	Ordinal int
	Name    string
}
