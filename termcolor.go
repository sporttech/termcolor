package termcolor

import (
	"runtime"
)

type Color string

const (
	// https://github.com/git/git/blob/master/color.h
	NORMAL       = Color("")
	RESET        = Color("\033[m")
	BOLD         = Color("\033[1m")
	RED          = Color("\033[31m")
	GREEN        = Color("\033[32m")
	YELLOW       = Color("\033[33m")
	BLUE         = Color("\033[34m")
	MAGENTA      = Color("\033[35m")
	CYAN         = Color("\033[36m")
	BOLD_RED     = Color("\033[1;31m")
	BOLD_GREEN   = Color("\033[1;32m")
	BOLD_YELLOW  = Color("\033[1;33m")
	BOLD_BLUE    = Color("\033[1;34m")
	BOLD_MAGENTA = Color("\033[1;35m")
	BOLD_CYAN    = Color("\033[1;36m")
	BG_RED       = Color("\033[41m")
	BG_GREEN     = Color("\033[42m")
	BG_YELLOW    = Color("\033[43m")
	BG_BLUE      = Color("\033[44m")
	BG_MAGENTA   = Color("\033[45m")
	BG_CYAN      = Color("\033[46m")
)

func Colorized(s string, c Color) string {
	if runtime.GOOS == "windows" {
		return s
	}
	return string(c) + s + string(RESET)
}
