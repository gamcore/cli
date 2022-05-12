package log

var (
	defaultLogger Logger
)

func SetDefaultLogger(logger Logger) {
	defaultLogger = logger
}

func Close() error {
	return defaultLogger.Close()
}

func DebugF(message string, f ...any) {
	defaultLogger.printf(LevelDebug, message, f...)
}

func Debug(message string) {
	defaultLogger.println(LevelDebug, message)
}

func InfoF(message string, f ...any) {
	defaultLogger.printf(LevelInfo, message, f...)
}

func Info(message string) {
	defaultLogger.println(LevelInfo, message)
}

func WarnF(message string, f ...any) {
	defaultLogger.printf(LevelWarn, message, f...)
}

func Warn(message string) {
	defaultLogger.println(LevelWarn, message)
}

func ErrorF(message string, f ...any) {
	defaultLogger.printf(LevelError, message, f...)
}

func Error(message string) {
	defaultLogger.println(LevelError, message)
}

func CriticalF(message string, f ...any) {
	defaultLogger.printf(LevelCritical, message, f...)
}

func Critical(message string) {
	defaultLogger.println(LevelCritical, message)
}

func FatalF(message string, f ...any) {
	defaultLogger.printf(LevelFatal, message, f...)
}

func Fatal(message string) {
	defaultLogger.println(LevelFatal, message)
}
