package main

import (
	"encoding/json"
	"hirenpandit/go-changelog/pkg/logs"
	"log/slog"
)

func main() {

	logsSlice := logs.Logs()
	logger := slog.Default()

	for i, l := range logsSlice {
		fLog, _ := json.MarshalIndent(l, "", "  ")
		logger.Info("log >>", slog.Attr{Key: "Index", Value: slog.IntValue(i)}, slog.Attr{Key: "item", Value: slog.AnyValue(fLog)})
	}

}
