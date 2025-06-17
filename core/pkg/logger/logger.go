package logger

import (
	"go.uber.org/zap"
)

/*
This package wraps the zap logging API for use across services without passing a shared logger.
Fields can be added to a requestID using the WriteFields method, these will be added to any
subsequent XxxWithID log calls. To preserve performance ClearFields must be called when the
requestID's thread is closed as a sync map is used internally.
Child loggers can be created from a parent logger using the WithFields method, this child logger
will append the provided fields to all logs, whilst maintaining a reference to the top level
request fields pool.

Example:

WrappedLogger := NewLogger(myLogger)
WrappedLogger.WriteFields("my-id", zap.String("foo", "bar"))
WrappedLogger2 := WrappedLogger.WithFields(zap.String("ping", "pong"))

WrappedLogger.DebugWithID("myID", "my log line")
	=> {"level":"debug","requestID":"myID","foo":"bar","msg":"my log line""}

WrappedLogger2.DebugWithID("myID", "my log line")
	=> {"level":"debug","requestID":"myID","foo":"bar","ping":"pong","msg":"my log line""}

WrappedLogger2.WriteFields("myID", zap.String("food", "bars"))

WrappedLogger.DebugWithID("myID", "my log line")
	=> {"level":"debug","requestID":"myID","foo":"bar","food":"bars","msg":"my log line""}
*/

const RequestIDFieldName = "requestID"

type Logger interface {
	DebugWithID(reqID string, msg string, args ...any)
	Debug(msg string, fields ...zap.Field)
	InfoWithID(reqID string, msg string, args ...any)
	Info(msg string, fields ...zap.Field)
	WarnWithID(reqID string, msg string, args ...any)
	Warn(msg string, fields ...zap.Field)
	ErrorWithID(reqID string, msg string, args ...any)
	Error(msg string, fields ...zap.Field)
	FatalWithID(reqID string, msg string, args ...any)
	Fatal(msg string, args ...any)
	WriteFields(reqID string, args ...any)
	ClearFields(reqID string)
}
