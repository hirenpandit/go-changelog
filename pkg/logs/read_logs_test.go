package logs

import "testing"

func TestLogs(t *testing.T) {
	formattedLogs := Logs()
	have := len(formattedLogs)
	if have < 0 {
		t.Errorf("expected length of logs greater than 0 found %d", have)
	}
}
