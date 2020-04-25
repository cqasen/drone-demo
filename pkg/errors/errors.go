package errors

import (
	"fmt"
	"log"
)

func FatalError(msg string, err error) {
	if err != nil {
		defer coverPanic()
		//panic(errors.New(fmt.Sprintf("%s Error: %v\n", msg, err)))
		panic(fmt.Sprintf("%s Error: %v\n", msg, err))
	}

	log.Printf("%s Success\n", msg)
}
func coverPanic() {
	message := recover()
	log.Printf("coverPanic:%s", message)
}
