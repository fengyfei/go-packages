package runner

import (
	"fmt"
	"runtime/debug"
)

func defaultRecoverHandler(err interface{}) {
	if str, ok := err.(string); ok {
		fmt.Printf("runner: default recover capture a panic [%s]\n", str)
	} else {
		fmt.Printf("runner: default recover capture a panic [%v]\n", err.(error))
	}

	debug.PrintStack()
}
