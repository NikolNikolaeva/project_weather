package logs

import (
	"log/slog"
	"strings"

	"github.com/NikolNikolaeva/project_weather/it/testbed/internal/gomisc/lang/array"
)

var (
	LevelDebug = &_Level{raw: slog.LevelDebug, keys: []string{"dbg", "debug"}}
	LevelInfo  = &_Level{raw: slog.LevelInfo, keys: []string{"info", "information"}}
	LevelWarn  = &_Level{raw: slog.LevelWarn, keys: []string{"warn", "warning"}}
	LevelError = &_Level{raw: slog.LevelError, keys: []string{"err", "error"}}
)

type Level interface {
	level() slog.Level
	matches(value string) bool
}

func ParseLevel(value string) Level {
	for _, level := range []Level{LevelDebug, LevelInfo, LevelWarn, LevelError} {
		if level.matches(value) {
			return level
		}
	}

	return LevelInfo // info by default
}

type _Level struct {
	keys []string
	raw  slog.Level
}

func (self *_Level) matches(value string) bool {
	return arrays.OneOf(strings.ToLower(value), self.keys...)
}

func (self *_Level) level() slog.Level {
	return self.raw
}
