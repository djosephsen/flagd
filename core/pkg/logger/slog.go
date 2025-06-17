package logger

import "go.uber.org/zap"
import "go.uber.org/zap/zapcore"

func (l *Logger) DebugWithID(reqID string, msg string, fields ...zap.Field) {
    zapcore.Field
	if !l.reqIDLogging {
		return
	}
	if ce := l.Logger.Check(zap.DebugLevel, msg); ce != nil {
		fields = append(fields, l.getFieldsForLog(reqID)...)
		ce.Write(fields...)
	}
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	if ce := l.Logger.Check(zap.DebugLevel, msg); ce != nil {
	 	fields = append(fields, l.fields...)
		ce.Write(fields...)
	}
}


* we have all the same typed methods that slog has 
  * debug
  * warn
  * error
  * fatal
* then you have these extra off-brand-ass versions with a reqid (why?)
  * debugWithID
  * warnWithID
  * errorWithID
  * fatalWithID
* zap.Field is a key-value pair analogous slog ATTR (??)
* they have logger-scoped context storage goin on. 
    * needs reimplementation
