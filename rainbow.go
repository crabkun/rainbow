package rainbow

import (
	"runtime"
	"syscall"
	"fmt"
	"os"
)
var kernel32 *syscall.DLL
var setConsoleTextAttribute *syscall.Proc
var stdoutHandle syscall.Handle
var fontColor Color
var backgrondColor Color

//SetFontColor Set console font color
func SetFontColor(color Color){
	setColor(1,color)
}
//SetFontDarkColor Set console font dark color
func SetFontDarkColor(color Color){
	color=color&7
	setColor(1,color)
}
//SetBackgroundColor Set console background color
func SetBackgroundColor(color Color){
	setColor(2,color)
}
//SetBackgroundDarkColor Set console background dark color
func SetBackgroundDarkColor(color Color){
	color=color&7
	setColor(2,color)
}
//CleanFontColor set console font color to white
func CleanFontColor(){
	setColor(1,White)
}
//CleanBackgroundColor set console background color to black
func CleanBackgroundColor(){
	setColor(2,Black)
}
//CleanColor set console background color to black,font color to white
func CleanColor(){
	CleanFontColor()
	CleanBackgroundColor()
}
func setColor(command int,color Color){
	var f Color
	switch command{
	case 1:
		fontColor=color
	case 2:
		backgrondColor=color
	}
	f=backgrondColor*16+fontColor
	setConsoleTextAttribute.Call(uintptr(stdoutHandle),uintptr(f))
}
func init(){
	if runtime.GOOS=="windows" {
		kernel32, _ = syscall.LoadDLL("kernel32.dll")
		setConsoleTextAttribute, _ = kernel32.FindProc("SetConsoleTextAttribute")
		stdoutHandle, _ = syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	}else{
		fmt.Println("only in windows")
		os.Exit(-1)
	}
}
