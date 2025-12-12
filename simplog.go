// Copyright © 2025 CToID <funk443@icloud.com>
//
// This file is part of Simplog.
//
// Simplog is free software. It comes without any warranty, to the
// extent permitted by applicable law. You can redistribute it and/or
// modify it under the terms of the Do What The Fuck You Want To Public
// License, Version 2, as published by Sam Hocevar. See the COPYING file
// for more details.

// Package simplog provides a simple logging functionality. I use this kind of
// logging a lot in my projects, so I think it's better to make it a separate
// package.
package simplog

import (
	"os"
	"fmt"
)

type Level int

// These are levels for logging.
const (
	Info Level = iota
	Error // Error is for continuable errors.
	Fatal // Fatal is for non-continuable errors.
)

var LevelNames = map[Level]string{
	Info: "INFO",
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

// I prefixes "INFO: " before the format string.
func I(format string, a ...any) {
	if CurrentLevel > Info {
		return
	}

	Printf("%s: ", Info.String())
	Printf(format, a...)
	Printf("\n")
}

// E prefixes "ERROR: " before the format string.
func E(format string, a ...any) {
	if CurrentLevel > Error {
		return
	}

	Printf("%s: ", Error.String())
	Printf(format, a...)
	Printf("\n")
}

// I prefixes "FATAL: " before the format string.
func F(format string, a ...any) {
	if CurrentLevel > Fatal {
		return
	}

	Printf("%s: ", Fatal.String())
	Printf(format, a...)
	Printf("\n")
}
