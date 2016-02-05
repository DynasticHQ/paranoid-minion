/*
 * Copyright (C) 2016 Miguel Moll
 *
 * This file is part of the Paranoid Minion
 *
 * The Paranoid Minion is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The Paranoid Minion is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with The Paranoid Minion.  If not, see <http://www.gnu.org/licenses/>.
 */

package minion

import "github.com/Sirupsen/logrus"

type Logger struct {
	log *logrus.Logger
}

type LogFields map[string]interface{}

var Log *Logger

// InitLogging initializes logging system used by the minion.
func InitLogging(mc MinionConfig) {
	Log = &Logger{logrus.New()}
	Log.SetLogFormatter(mc.Log.Format)
	Log.SetLogLevel(mc.Log.Level)
	// TODO(miguelmoll): Implement a type for logger. e.g. to file or stdout
	// Set to stdout because of defaults.

}

// SetLogFormatter sets the format (JSON, text, etc) used by the logger.
func (l *Logger) SetLogFormatter(format string) {
	switch format {
	case LogFormatJSON:
		l.log.Formatter = new(logrus.JSONFormatter)
	default:
		l.log.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	}
}

// SetLogLevel will set the log level (Info, Debug, etc).
func (l *Logger) SetLogLevel(level string) {
	switch level {
	case LogLevelDebug:
		l.log.Level = logrus.DebugLevel
	case LogLevelWarning:
		l.log.Level = logrus.WarnLevel
	case LogLevelError:
		l.log.Level = logrus.ErrorLevel
	default:
		l.log.Level = logrus.InfoLevel
	}
}

// Debug logs a message at level Debug.
func (l *Logger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

// Debugf logs a message at level Debug.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
}

// Info logs a message at level Info.
func (l *Logger) Info(args ...interface{}) {
	l.log.Info(args...)
}

// Infof logs a message at level Info.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

// Warn logs a message at level Warn.
func (l *Logger) Warn(args ...interface{}) {
	l.log.Warn(args...)
}

// Warnf logs a message at level Warn.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

// Error logs a message at level Error.
func (l *Logger) Error(args ...interface{}) {
	l.log.Error(args...)
}

// Errorf logs a message at level Error.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}
