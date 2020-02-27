package formatter

import "github.com/zippunov/logging/def"

func Level() def.Formatter {
	return &levelFormatter{}
}

type levelFormatter struct {
}

// GetPrefix returns ""
func (f *levelFormatter) GetPrefix(lvl def.Level) string {
	return lvl.Name()
}

// GetSuffix returns ""
func (f *levelFormatter) GetSuffix(lvl def.Level) string {
	return ""
}

// Format adds filename and line number before the log message
func (f *levelFormatter) Format(lvl def.Level, format string, values ...interface{}) (string, []interface{}) {
	return format, values
}
