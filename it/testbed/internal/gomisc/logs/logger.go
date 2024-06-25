package logs

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/NikolNikolaeva/project_weather/it/testbed/internal/gomisc/lang"
)

type ContextKey string

const LOG_CONTEXT = ContextKey("log_context")
const CORRELATION_ID = ContextKey("correlation_id")

var NopLogger = NewLogger(NewHandler(LevelDebug, io.Discard))

//go:generate mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/logs/mock_logger.go . Logger
type Logger interface {
	Handler() Handler
	Writer(level Level) io.Writer

	WithError(err error) Logger
	WithGroup(name string) Logger
	WithField(name string, value any) Logger
	WithContext(ctx context.Context) Logger

	Warn(template string, args ...any)
	Info(template string, args ...any)
	Debug(template string, args ...any)
	Error(template string, args ...any)
	Fatal(template string, args ...any)

	Log(level Level, message string, args ...any)
}

func NewLogger(handler Handler) Logger {
	return &_Logger{
		handler: handler,
	}
}

type _Logger struct {
	handler Handler
}

func (self *_Logger) Handler() Handler {
	return self.handler
}

func (self *_Logger) Writer(level Level) io.Writer {
	return slog.NewLogLogger(self.handler, level.level()).Writer()
}

func (self *_Logger) WithGroup(name string) Logger {
	return NewLogger(
		self.handler.
			WithGroup(name).(Handler),
	)
}

func (self *_Logger) WithError(err error) Logger {
	return NewLogger(
		self.handler.
			WithAttrs([]slog.Attr{slog.Any("error", err)}).(Handler),
	)
}

func (self *_Logger) WithField(name string, value any) Logger {
	return NewLogger(
		self.handler.
			WithAttrs([]slog.Attr{slog.Any(name, value)}).(Handler),
	)
}

func (self *_Logger) WithContext(ctx context.Context) Logger {
	result := self.WithField(string(CORRELATION_ID), ctx.Value(CORRELATION_ID))

	context, ok := ctx.Value(LOG_CONTEXT).(map[string]any)
	if ok {
		for key, value := range context {
			result = result.WithField(key, value)
		}
	}

	return result
}

func (self *_Logger) Warn(template string, args ...any) {
	self.handle(slog.LevelWarn, template, args...)
}

func (self *_Logger) Info(template string, args ...any) {
	self.handle(slog.LevelInfo, template, args...)
}

func (self *_Logger) Debug(template string, args ...any) {
	self.handle(slog.LevelDebug, template, args...)
}

func (self *_Logger) Error(template string, args ...any) {
	self.handle(slog.LevelError, template, args...)
}

func (self *_Logger) Log(level Level, template string, args ...any) {
	self.handle(level.level(), template, args...)
}

func (self *_Logger) Fatal(template string, args ...any) {
	defer os.Exit(1)

	self.handle(slog.LevelError, template, args...)
}

func (self *_Logger) handle(level slog.Level, template string, args ...any) {
	_ = self.handler.
		Handle(
			context.Background(),
			slog.NewRecord(
				time.Now(),
				level,
				fmt.Sprintf(template, args...),
				lang.First(runtime.Caller(2)), // magic!
			),
		)
}
