package logger

import (
	"log/slog"
	"sync"
)

type slogger struct {
	requestAttrs *sync.Map
	Logger       *slog.Logger
	reqIDLogging bool
}

var _ Logger = &slogger{}

func (l *slogger) DebugWithID(reqID string, msg string, args ...any) {

}
func (l *slogger) Debug(msg string, args ...any) {
	slog.Debug(msg, makeAttrs(args))
}
func (l *slogger) InfoWithID(reqID string, msg string, args ...any) {

}
func (l *slogger) Info(msg string, args ...any) {
	slog.Info(msg, makeAttrs(args))
}
func (l *slogger) WarnWithID(reqID string, msg string, args ...any) {

}
func (l *slogger) Warn(msg string, args ...any) {
	slog.Warn(msg, makeAttrs(args))
}
func (l *slogger) ErrorWithID(reqID string, msg string, args ...any) {

}
func (l *slogger) Error(msg string, args ...any) {
	slog.Error(msg, makeAttrs(args))
}
func (l *slogger) FatalWithID(reqID string, msg string, args ...any) {

}
func (l *slogger) Fatal(msg string, args ...any) {
	//slog.Fatal(msg, makeAttrs(args))
	// There is no Fatal level in slog.  Is it needed? Should we added it?
}

// ============================================

func makeAttrs(args ...any) []slog.Attr {
	// if there are an odd number of
	if len(args)%2 != 0 {
		lastArg := args[len(args)-1]
		args[len(args)-1] = "BADKEY!"
		args = append(args, lastArg)
	}
	out := make([]slog.Attr, len(args))
	for i := 0; i < len(args); i += 2 {
		var k string
		switch args[i].(type) {
		case string:
			k = k
		default:
			k = "!BADKEY"
		}
		val := args[i+1]
		out = append(out, slog.Any(k, val))
	}
	return out
}
func (l *slogger) WriteFields(reqID string, fields ...slog.Attr) {
	if !l.reqIDLogging {
		return
	}
	res := append(l.getFields(reqID), fields...)
	l.requestAttrs.Store(reqID, res)
}

func (l *slogger) getFields(reqID string) []slog.Attr {
	res := []slog.Attr{}
	f, ok := l.requestAttrs.Load(reqID)
	if ok {
		r, ok := f.([]slog.Attr)
		if ok {
			res = r
		}
	}
	return res
}

func (l *slogger) getFieldsForLog(reqID string) []slog.Attr {
	fields := l.getFields(reqID)
	fields = append(fields, slog.Any("requestID", reqID))
	fields = append(fields, l.fields...)
	return fields
}

// ClearFields clears all stored fields for a given requestID, important for maintaining performance
func (l *slogger) ClearFields(reqID string) {
	if !l.reqIDLogging {
		return
	}
	l.requestAttrs.Delete(reqID)
}
