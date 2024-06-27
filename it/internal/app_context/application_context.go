package app_context

import (
	"context"
	"time"
)

type Context interface {
	context.Context
	Cancel(cause error)
	Child(timeout time.Duration) Context
}

type _Context struct {
	context context.Context
	cancel  context.CancelCauseFunc
}

func New() Context {
	context, cancel := context.WithCancelCause(context.Background())

	return &_Context{
		cancel:  cancel,
		context: context,
	}
}

func (self *_Context) Deadline() (deadline time.Time, ok bool) {
	return self.context.Deadline()
}

func (self *_Context) Done() <-chan struct{} {
	return self.context.Done()
}

func (self *_Context) Err() error {
	return context.Cause(self.context)
}

func (self *_Context) Value(key any) any {
	return self.context.Value(key)
}

func (self *_Context) Cancel(cause error) {
	self.cancel(cause)
}

func (self *_Context) Child(timeout time.Duration) Context {
	timed, cancelTimed := context.WithTimeout(self.context, timeout)
	withCause, cancelWithCause := context.WithCancelCause(timed)

	return &_Context{
		context: withCause,
		cancel: func(cause error) {
			cancelWithCause(cause)
			cancelTimed()
		},
	}
}
