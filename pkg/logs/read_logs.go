package logs

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type FLog struct {
	Msg  string    `json:"Message"`
	Date time.Time `json:"Date"`
}

// yyyy-MM-dd format
const DATE_FORMAT string = "2006-01-02"

/*
`ReadLog` executes git command to get the logs on current branch.
it returns slice of string each item in slice contains single line log
*/
func ReadLog() []FLog {

	cmd := exec.Command("git", "log", "--pretty=reference")

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Printf("error while runnng the command %v error: %v", cmd, err)
	}

	logs := strings.TrimSpace(stdout.String())
	arr := strings.Split(logs, "\n")

	formattedLogs := make([]FLog, 0)

	for _, item := range arr {
		s := strings.SplitAfter(item, "(")
		if len(s) < 1 {
			fmt.Println("not expected format of log >", s)
		}
		log := strings.TrimSuffix(s[1], ")")

		log, after, _ := strings.Cut(log, ", ")

		logDate, err := time.Parse(DATE_FORMAT, after)
		if err != nil {
			fmt.Println("error parsing date", after, "in format", "yyyy-mm-dd", "error", err)
		}

		formattedLogs = append(formattedLogs, FLog{Msg: log, Date: logDate})
	}

	return formattedLogs
}
