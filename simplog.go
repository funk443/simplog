// Copyright (C) 2026 by CToID <funk443@icloud.com>
// SPDX-License-Identifier: 0BSD
//
// This file is part of simplog.
//
// simplog is free software. It comes without any warranty, to the
// extent permitted by applicable law. You can use, copy, modify, and/or
// distribute it under the terms of the BSD Zero Clause License. See the
// COPYING file for more details.

// Package simplog provides a simple logging functionality. I use this kind of
// logging a lot in my projects, so I think it's better to make it a separate
// package. Convinience functions to log a certain log level is provided, such
// as [D], [I], [E], [F].
package simplog

import (
	"fmt"
	"os"
)

type Level int

// These are levels for logging.
const (
	Debug = Level(iota)
	Info
	Error // Error is for continuable errors.
	Fatal // Fatal is for non-continuable errors.
)

var LevelNames = map[Level]string{
	Debug: "DEBUG",
	Info:  "INFO",
	Error: "ERROR",
	Fatal: "FATAL",
}

func (s Level) String() string {
	return LevelNames[s]
}

var CurrentLevel = Info

// Printf is almost identical to [fmt.Printf], except that it prints to
// [os.Stderr].
func Printf(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format, a...)
}

// Log checks if level is smaller than [CurrentLevel], if so, the message is
// printed with level's [LevelNames] as prefix.
func Log(level Level, format string, a ...any) {
	if CurrentLevel > level {
		return
	}

	Printf("%s: ", level.String())
	Printf(format, a...)
	Printf("\n")
}

// D prefixes "DEBUG: " before the format string.
func D(format string, a ...any) {
	Log(Debug, format, a...)
}

// I prefixes "INFO: " before the format string.
func I(format string, a ...any) {
	Log(Info, format, a...)
}

// E prefixes "ERROR: " before the format string.
func E(format string, a ...any) {
	Log(Error, format, a...)
}

// F prefixes "FATAL: " before the format string.
func F(format string, a ...any) {
	Log(Fatal, format, a...)
}
