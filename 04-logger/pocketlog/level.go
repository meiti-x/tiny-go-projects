package pocketlog

const (
	// LevelDebug represents the lowest level of log, mostly used for debugging
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information deemed
	LevelInfo
	// LevelError represents the highest logging level, only to be used to t
	LevelError
)
