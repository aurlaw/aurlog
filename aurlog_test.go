package aurlog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func getLogFileName(logFile string) string {
	const layout = "2006-01-02"
	t := time.Now()
	dir := filepath.Dir(logFile)
	fileName := filepath.Base(logFile)

	newFileName := fmt.Sprintf("%s_%s", t.Format(layout), fileName)
	return filepath.Join(dir, newFileName)

}

func TestLogWriter(t *testing.T) {

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	a := Configure(nil)
	msg := "Test Debug"
	a.Debugln(msg)

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	if !strings.Contains(out, msg) {
		// expect(t, msg, out)
		t.Errorf(out + "does not contain " + msg)
	}

}

func TestLogFile(t *testing.T) {
	logFile := "../../test.log"

	testLogName := getLogFileName(logFile)

	if _, err := os.Stat(testLogName); err == nil {
		fmt.Println(testLogName + " found. remove")
		os.Remove(testLogName)
	}
	lc := LogConfiguration{LogFile: logFile}
	a := Configure(&lc)

	a.Infoln("Test log file")
	if _, err := os.Stat(testLogName); os.IsNotExist(err) {
		t.Errorf("logfile does not exist. " + testLogName)
	}
}

// func main() {

// 	al := Configure(nil)

// 	al.Debugln("I have something to say for debug")
// 	al.Infoln("I have something to say for info")
// 	al.Warningln("I have something to say for warning")
// 	al.Errorln("I have something to say for error")

// 	lc := LogConfiguration{LogFile: "test.log", NoStdOut: true}
// 	lc.IsInfo = true
// 	lc.IsError = true

// 	al2 := Configure(&lc)

// 	al2.Debugln("I have something to say for debug again")
// 	al2.Infoln("I have something to say for info again")
// 	al2.Warningln("I have something to say for warning again")
// 	al2.Errorln("I have something to say for error again")

// }
