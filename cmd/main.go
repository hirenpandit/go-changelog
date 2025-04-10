package main

import (
	"fmt"
	"hirenpandit/go-changelog/pkg/logs"
)

func main() {

	logsSlice := logs.Logs()

	for i, l := range logsSlice {
		fmt.Printf("%d log > %s \n", i, l)
	}

}
