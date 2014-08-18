package aurlog

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

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
