// mauLogger - A logger for Go programs
// Copyright (C) 2016 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package maulogger ...
package maulogger

import (
	"fmt"
)

// Sublogger ...
type Sublogger struct {
	Parent       *Logger
	Module       string
	DefaultLevel Level
}

// CreateSublogger creates a Sublogger
func (log *Logger) CreateSublogger(module string, DefaultLevel Level) *Sublogger {
	return &Sublogger{
		Parent:       log,
		Module:       module,
		DefaultLevel: DefaultLevel,
	}
}

// SetModule changes the module name of this Sublogger
func (log *Sublogger) SetModule(mod string) {
	log.Module = mod
}

// SetDefaultLevel changes the default logging level of this Sublogger
func (log *Sublogger) SetDefaultLevel(lvl Level) {
	log.DefaultLevel = lvl
}

// SetParent changes the parent of this Sublogger
func (log *Sublogger) SetParent(parent *Logger) {
	log.Parent = parent
}

//Write ...
func (log *Sublogger) Write(p []byte) (n int, err error) {
	log.Parent.Raw(log.DefaultLevel, log.Module, string(p))
	return len(p), nil
}

// Log formats the given parts with fmt.Sprint and log them with the Log level
func (log *Sublogger) Log(level Level, parts ...interface{}) {
	log.Parent.Raw(level, "", fmt.Sprint(parts...))
}

// Logln formats the given parts with fmt.Sprintln and log them with the Log level
func (log *Sublogger) Logln(level Level, parts ...interface{}) {
	log.Parent.Raw(level, "", fmt.Sprintln(parts...))
}

// Logf formats the given message and args with fmt.Sprintf and log them with the Log level
func (log *Sublogger) Logf(level Level, message string, args ...interface{}) {
	log.Parent.Raw(level, "", fmt.Sprintf(message, args...))
}

// Debug formats the given parts with fmt.Sprint and log them with the Debug level
func (log *Sublogger) Debug(parts ...interface{}) {
	log.Parent.Raw(LevelDebug, log.Module, fmt.Sprint(parts...))
}

// Debugln formats the given parts with fmt.Sprintln and log them with the Debug level
func (log *Sublogger) Debugln(parts ...interface{}) {
	log.Parent.Raw(LevelDebug, log.Module, fmt.Sprintln(parts...))
}

// Debugf formats the given message and args with fmt.Sprintf and log them with the Debug level
func (log *Sublogger) Debugf(message string, args ...interface{}) {
	log.Parent.Raw(LevelDebug, log.Module, fmt.Sprintf(message, args...))
}

// Info formats the given parts with fmt.Sprint and log them with the Info level
func (log *Sublogger) Info(parts ...interface{}) {
	log.Parent.Raw(LevelInfo, log.Module, fmt.Sprint(parts...))
}

// Infoln formats the given parts with fmt.Sprintln and log them with the Info level
func (log *Sublogger) Infoln(parts ...interface{}) {
	log.Parent.Raw(LevelInfo, log.Module, fmt.Sprintln(parts...))
}

// Infof formats the given message and args with fmt.Sprintf and log them with the Info level
func (log *Sublogger) Infof(message string, args ...interface{}) {
	log.Parent.Raw(LevelInfo, log.Module, fmt.Sprintf(message, args...))
}

// Warn formats the given parts with fmt.Sprint and log them with the Warn level
func (log *Sublogger) Warn(parts ...interface{}) {
	log.Parent.Raw(LevelWarn, log.Module, fmt.Sprint(parts...))
}

// Warnln formats the given parts with fmt.Sprintln and log them with the Warn level
func (log *Sublogger) Warnln(parts ...interface{}) {
	log.Parent.Raw(LevelWarn, log.Module, fmt.Sprintln(parts...))
}

// Warnf formats the given message and args with fmt.Sprintf and log them with the Warn level
func (log *Sublogger) Warnf(message string, args ...interface{}) {
	log.Parent.Raw(LevelWarn, log.Module, fmt.Sprintf(message, args...))
}

// Error formats the given parts with fmt.Sprint and log them with the Error level
func (log *Sublogger) Error(parts ...interface{}) {
	log.Parent.Raw(LevelError, log.Module, fmt.Sprint(parts...))
}

// Errorln formats the given parts with fmt.Sprintln and log them with the Error level
func (log *Sublogger) Errorln(parts ...interface{}) {
	log.Parent.Raw(LevelError, log.Module, fmt.Sprintln(parts...))
}

// Errorf formats the given message and args with fmt.Sprintf and log them with the Error level
func (log *Sublogger) Errorf(message string, args ...interface{}) {
	log.Parent.Raw(LevelError, log.Module, fmt.Sprintf(message, args...))
}

// Fatal formats the given parts with fmt.Sprint and log them with the Fatal level
func (log *Sublogger) Fatal(parts ...interface{}) {
	log.Parent.Raw(LevelFatal, log.Module, fmt.Sprint(parts...))
}

// Fatalln formats the given parts with fmt.Sprintln and log them with the Fatal level
func (log *Sublogger) Fatalln(parts ...interface{}) {
	log.Parent.Raw(LevelFatal, log.Module, fmt.Sprintln(parts...))
}

// Fatalf formats the given message and args with fmt.Sprintf and log them with the Fatal level
func (log *Sublogger) Fatalf(message string, args ...interface{}) {
	log.Parent.Raw(LevelFatal, log.Module, fmt.Sprintf(message, args...))
}
