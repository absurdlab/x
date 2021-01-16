package tracer

import (
	"errors"
	"regexp"
	"runtime"
)

// Normalize converts any error to a tracer error. If the error is already a tracer error,
// the error is returned as is with type asserted.
func Normalize(err error) *tracerErr {
	if err == nil {
		return nil
	}

	te, ok := err.(*tracerErr)
	if ok {
		return te
	}

	return &tracerErr{
		err:    err,
		frames: []Frame{},
	}
}

// Touch records the file and line number of the direct caller's call site of this method. It
// adds a trace information to the current error.
//
//	// For example:
//	if err != nil {
//		return trace.Touch(err)
//	}
//
// The recorded file information is prefix trimmed of the common project and module path to only
// leave the top level directories of this project.
func Touch(err error) *tracerErr {
	te := Normalize(err)

	_, f, l, ok := runtime.Caller(1)
	if ok {
		fm := Frame{File: f, Line: l}
		fm.File = trimPath.ReplaceAllString(fm.File, "/")
		te.frames = append(te.frames, fm)
	}

	return te
}

// Trace renders the error traces in a slice of Frame. It first converts the
// error to trace error by Normalize it, and then invoke Traces method on the
// error.
//
// Its return signature is designed to be interface{} to be compatible with
// zerolog.ErrorStackMarshaler so that this function can be used directly.
//
//	zerolog.ErrorStackMarshaler = tracer.Trace
//	log.Ctx(ctx).Debug().Stack().Err(err).Msg("trace will be rendered")
//
// If normal return signature of []Frame is required, invoke Normalize and Traces
// directly:
//
//	Normalize(err).Traces()
func Trace(err error) interface{} {
	return Normalize(err).Traces()
}

type tracerErr struct {
	err error

	// cause is the wrapped error, which forms
	// a traceable chain when multiple tracerErr
	// is linked together with cause.
	cause error

	// frames is the stack of Frame that the caller
	// would like to trace for the current traceErr.
	frames []Frame
}

// Error returns the error string of this error.
func (e *tracerErr) Error() string {
	return e.err.Error()
}

// Wrap sets the given error as this error's cause, thus forming
// a chain. Wrap has no effect when this error already has a cause.
func (e *tracerErr) Wrap(w error) *tracerErr {
	if e.cause == nil {
		e.cause = w
	}
	return e
}

// WrapStr calls Wrap with a new error of w.
func (e *tracerErr) WrapStr(w string) *tracerErr {
	return e.Wrap(errors.New(w))
}

// Traces collects all trace Frame by following the cause chain. The returned frames
// are ordered in reversed chronological order of invoking Touch and Wrap.
func (e *tracerErr) Traces() []Frame {
	var frames []Frame

	for c := e; c != nil; c = Normalize(c.cause) {
		for i := len(c.frames) - 1; i >= 0; i-- {
			frames = append(frames, c.frames[i])
		}
	}

	if frames == nil {
		return []Frame{}
	}

	return frames
}

// Unwrap returns the cause of the error, so it could work
// well with the native errors wrapping utilities.
func (e *tracerErr) Unwrap() error {
	return e.cause
}

// Is returns true if the provided error is this trace error or
// the error contained in this trace error.
func (e *tracerErr) Is(err error) bool {
	if e == nil {
		return false
	}
	return e == err || e.err == err
}

// Frame is one chunk of trace data.
type Frame struct {
	// File is the file name of the traced location. It
	// could be trimmed for better visual experience.
	File string `json:"file"`
	// Line is the line number in the File.
	Line int `json:"line"`
}

var trimPath = regexp.MustCompile(`^(.*)/absurdlab/tiga/`)
