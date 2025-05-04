package main

import (
	// "encoding/json"
	"hirenpandit/go-changelog/pkg/logs"
	// "log/slog"
)

func main() {

	logsSlice := logs.Logs()
	// logger := slog.Default()
	//
	// for i, flogItem := range logsSlice {
	// 	fLog, _ := json.MarshalIndent(flogItem, "", "  ")
	// 	logger.Info("log >>",
	// 		slog.Attr{Key: "Index", Value: slog.IntValue(i)},
	// 		slog.Attr{Key: "Object", Value: slog.AnyValue(fLog)})
	//
	// }

	logs.WriteLog(logsSlice)

}
