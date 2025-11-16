package main

import "syscall"

var (
	user32               = syscall.MustLoadDLL("user32.dll")
	procShowWindow       = user32.MustFindProc("ShowWindow")
	kernel32             = syscall.MustLoadDLL("kernel32.dll")
	procGetConsoleWindow = kernel32.MustFindProc("GetConsoleWindow")
)

const (
	SW_HIDE = 0
	SW_SHOW = 5
)

func HideConsole() {
	hwnd, _, _ := procGetConsoleWindow.Call()
	if hwnd != 0 {
		procShowWindow.Call(hwnd, SW_HIDE)
	}
}

func ShowConsole() {
	hwnd, _, _ := procGetConsoleWindow.Call()
	if hwnd != 0 {
		procShowWindow.Call(hwnd, SW_SHOW)
	}
}
