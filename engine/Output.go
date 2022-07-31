package engine

import "fmt"

func Warn(msg string) {
	output("Warning", msg)
}

func Error(msg string) {
	output("Error", msg)
}

func output(level string, msg string) {
	fmt.Printf("%s: %s", level, msg)
}