package logging

type LogFunction func(levelName string, message string, err error, args ...any)
type AbstractLogger struct {
	Level       *level
	LogFunction LogFunction
}

// Verbose .
func (this *AbstractLogger) Verbose(message string, args ...any) {
	this.logByLevel(Level.VERBOSE, message, nil, args...)
}

// Debug .
func (this *AbstractLogger) Debug(message string, args ...any) {
	this.logByLevel(Level.DEBUG, message, nil, args...)
}

// Info .
func (this *AbstractLogger) Info(message string, args ...any) {
	this.logByLevel(Level.INFO, message, nil, args...)
}

// Warning .
func (this *AbstractLogger) Warning(message string, args ...any) {
	this.logByLevel(Level.WARNING, message, nil, args...)
}

// Error .
func (this *AbstractLogger) Error(message string, err error, args ...any) {
	this.logByLevel(Level.ERROR, message, err, args...)
}

// Fatal .
func (this *AbstractLogger) Fatal(message string, args ...any) {
	this.logByLevel(Level.FATAL, message, nil, args...)
}

// LogByLevel .
func (this *AbstractLogger) logByLevel(level *level, message string, err error, args ...any) {
	if level.Ordinal >= this.Level.Ordinal {
		this.log(level.Name, message, err, args...)
	}
}

// Log .
func (this *AbstractLogger) log(levelName string, message string, err error, args ...any) {
	this.LogFunction(levelName, message, err, args...)
}

// Destroy .
func (this *AbstractLogger) Destroy() {
}
