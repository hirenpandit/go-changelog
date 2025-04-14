package main

import (
	"encoding/json"
	"fmt"
	"hirenpandit/go-changelog/pkg/logs"
)

func main() {

	logsSlice := logs.Logs()

	for i, l := range logsSlice {
		fLog, _ := json.MarshalIndent(l, "", "  ")
		fmt.Printf("%d log > %s \n", i, fLog)
	}

}
