package logs

import (
	"context"
	arrays "github.com/NikolNikolaeva/project_weather/it/testbed/internal/gomisc/lang/array"
	"io"
	"log/slog"
)

type Handler interface {
	slog.Handler

	Level() Level
	Writer() io.Writer
	Processors() []Processor
}

func NewHandler(level Level, writer io.Writer, processors ...Processor) Handler {
	return &_Handler{
		level:      level,
		writer:     writer,
		processors: processors,
		inner: slog.NewTextHandler(
			writer,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     level.level(),
			},
		),
	}
}

type _Handler struct {
	level      Level
	writer     io.Writer
	processors []Processor
	inner      slog.Handler
}

func (self *_Handler) Level() Level {
	return self.level
}

func (self *_Handler) Writer() io.Writer {
	return self.writer
}

func (self *_Handler) Processors() []Processor {
	return append([]Processor{}, self.processors...)
}

func (self *_Handler) WithGroup(name string) slog.Handler {
	return &_Handler{
		level:      self.level,
		inner:      self.inner.WithGroup(name),
		processors: append(make([]Processor, 0), self.processors...),
	}
}

func (self *_Handler) WithAttrs(attributes []slog.Attr) slog.Handler {
	return &_Handler{
		level:      self.level,
		inner:      self.inner.WithAttrs(attributes),
		processors: append(make([]Processor, 0), self.processors...),
	}
}

func (self *_Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return self.inner.Enabled(ctx, level)
}

func (self *_Handler) Handle(ctx context.Context, record slog.Record) error {
	if !self.Enabled(ctx, record.Level) {
		return nil
	}

	return self.inner.Handle(
		ctx,
		arrays.Reduce(self.processors, record, func(current slog.Record, processor Processor) slog.Record {
			return processor.Process(current)
		}),
	)
}
