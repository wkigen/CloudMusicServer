package log

import(
	"fmt"
)

const ( 
	Debug = iota
	Info
	Warn
	Error
	Fatel
)

var levelString = [...]string{
	"[DEBUG]",
	"[INFO]",
	"[WARN]",
	"[ERROR]",
	"[FATAL]",
}

func Log(level int, format string, v ...interface{}){
	var text string
	if format == "" {
		text = fmt.Sprintln(v...)
	} else {
		text = fmt.Sprintf(format, v...)
	}
	fmt.Println(levelString[level]+text)
}