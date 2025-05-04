package logs

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

var logger slog.Logger = *slog.Default()

type Metadata struct {
	Name       string
	Version    string
	DateFormat string
}

const LOG_FILENAME = "changelog.md"

func WriteLog(logSlice []FLog, metadata Metadata) {

	isExist := checkLogFileExist()

	var logfile *os.File

	if !isExist {
		logger.Info("Creating", "logfile:", LOG_FILENAME)
		logfile = createLogFile()
	} else {

		openFile, err := os.OpenFile(LOG_FILENAME, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			logger.Error("Error opening logfile")
		}

		logfile = openFile
	}

	logger.Info("Ready to write logs to ", "logfile:", logfile.Name())

	writeLogHeader(logfile, metadata)

	for _, flog := range logSlice {
		bytesToWrite := []byte("\t" + flog.Date.Format(DATE_FORMAT) + " : " + flog.Msg + "\n")
		logfile.Write([]byte(bytesToWrite))
	}

}

// check if log file to write logs exist or not
func checkLogFileExist() bool {

	_, err := os.Stat(LOG_FILENAME)

	if err != nil {
		logger.Error("File ", LOG_FILENAME, "does not exit")
		return false
	}
	return true
}

// create new logfile
func createLogFile() *os.File {

	logfile, err := os.Create(LOG_FILENAME)

	if err != nil {
		logger.Error("Error creating", "logfile:", LOG_FILENAME)
	}

	return logfile
}

// writing header for new logs
func writeLogHeader(file *os.File, metadata Metadata) {

	logger.Info("Writing header to", "logfile:", file)

	today := time.Now().Format(DATE_FORMAT)

	header := "--------------------------------------\n %s-%s (%s)\n--------------------------------------\n"
	header = fmt.Sprintf(header, metadata.Name, metadata.Version, today)
	_, err := file.Write([]byte(header))
	if err != nil {
		logger.Error(
			"Error writing header to logfile %v",
			slog.Attr{Key: "error", Value: slog.AnyValue(err)},
		)
	}

}
