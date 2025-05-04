package main

import (
	// "encoding/json"
	"fmt"
	"hirenpandit/go-changelog/pkg/logs"
	"os"
	// "log/slog"
)

func main() {

	args := os.Args[1:]

	fmt.Println(">>>", args)

	var metadata logs.Metadata = logs.Metadata{
		Name:       args[0],
		Version:    args[1],
		DateFormat: args[2],
	}

	logsSlice := logs.ReadLog()
	// logger := slog.Default()
	//
	// for i, flogItem := range logsSlice {
	// 	fLog, _ := json.MarshalIndent(flogItem, "", "  ")
	// 	logger.Info("log >>",
	// 		slog.Attr{Key: "Index", Value: slog.IntValue(i)},
	// 		slog.Attr{Key: "Object", Value: slog.AnyValue(fLog)})
	//
	// }

	logs.WriteLog(logsSlice, metadata)

}
