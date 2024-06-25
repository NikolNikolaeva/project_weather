package logs

import "log/slog"

type Processor interface {
	Process(record slog.Record) slog.Record
}

func NewProcessor(processor func(slog.Record) slog.Record) Processor {
	return &_Processor{
		processor: processor,
	}
}

type _Processor struct {
	processor func(slog.Record) slog.Record
}

func (self *_Processor) Process(record slog.Record) slog.Record {
	return self.processor(record)
}
