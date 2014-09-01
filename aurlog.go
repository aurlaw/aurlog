// A custom log library that extends the Standard Go log

// Defines a Debug, Info, Warning and Error log level

// Allows for redirecting log output to a file and/or the Standard Out & Err

package aurlog

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

// AurLog represents the top-level struct
type AurLog struct {
}

// LogConfiguration defines configuration settings
// NoStdOut: log will not redirect to std out or std err
// LogFile: path to write output. File name will be prepended with date format yyyy-mm-dd
type LogConfiguration struct {
	LogLevel
	NoStdOut      bool
	LogFileRotate string
	LogFile       string
}

// LogLevel defines the logging levels: Debug, Info, Warning and Error
type LogLevel struct {
	IsDebug   bool
	IsInfo    bool
	IsWarning bool
	IsError   bool
}

var (
	debugLvl   *log.Logger
	infoLvl    *log.Logger
	warningLvl *log.Logger
	errorLvl   *log.Logger
	fatalLvl   *log.Logger
	logFile    *string
)

// Configures the Log using the LogConfiguration. Passing nil will set all levels of logging to true and will only output to StdOut & StdErr
func Configure(config *LogConfiguration) *AurLog {
	if config == nil {
		config = &LogConfiguration{}
		config.IsDebug = true
		config.IsInfo = true
		config.IsWarning = true
		config.IsError = true
	}
	file := ioutil.Discard
	std := ioutil.Discard
	stdErr := ioutil.Discard

	if config.LogFile != "" {
		const layout = "2006-01-02"
		t := time.Now()
		dir := filepath.Dir(config.LogFile)
		fileName := filepath.Base(config.LogFile)

		newFileName := fmt.Sprintf("%s_%s", t.Format(layout), fileName)
		absFile := filepath.Join(dir, newFileName)
		lfile, err := os.OpenFile(absFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		file = lfile
	}
	if !config.NoStdOut {
		std = os.Stdout
		stdErr = os.Stderr
	}

	debugHandle := ioutil.Discard
	infoHandle := ioutil.Discard
	warningHandle := ioutil.Discard
	errorHandle := ioutil.Discard

	if config.IsDebug {
		debugHandle = io.MultiWriter(file, std)
	}
	if config.IsInfo {
		infoHandle = io.MultiWriter(file, std)
	}
	if config.IsWarning {
		warningHandle = io.MultiWriter(file, std)
	}
	if config.IsError {
		errorHandle = io.MultiWriter(file, stdErr)
	}

	fatalHandle := io.MultiWriter(file, stdErr)

	debugLvl = log.New(debugHandle,
		"DEBUG: ",
		log.Ldate|log.Ltime)

	infoLvl = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime)

	warningLvl = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime)

	errorLvl = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime)

	fatalLvl = log.New(fatalHandle,
		"FATAL:",
		log.Ldate|log.Ltime)

	return &AurLog{}
}

// Debug message using log.Print
func (l *AurLog) Debug(v ...interface{}) {
	debugLvl.Print(v...)
}

// Debug message using log.PrintLn
func (l *AurLog) Debugln(v ...interface{}) {
	debugLvl.Println(v...)
}

// Debug message using log.Printf
func (l *AurLog) Debugf(format string, v ...interface{}) {
	debugLvl.Printf(format, v...)
}

// Info message using log.Print
func (l *AurLog) Info(v ...interface{}) {
	infoLvl.Print(v...)
}

// Info message using log.PrintLn
func (l *AurLog) Infoln(v ...interface{}) {
	infoLvl.Println(v...)
}

// Info message using log.Printf
func (l *AurLog) Infof(format string, v ...interface{}) {
	infoLvl.Printf(format, v...)
}

// Warning message using log.Print
func (l *AurLog) Warning(v ...interface{}) {
	warningLvl.Print(v...)
}

// Warning message using log.PrintLn
func (l *AurLog) Warningln(v ...interface{}) {
	warningLvl.Println(v...)
}

// Warning message using log.Printf
func (l *AurLog) Warningf(format string, v ...interface{}) {
	warningLvl.Printf(format, v...)
}

// Error message using log.Print
func (l *AurLog) Error(v ...interface{}) {
	errorLvl.Print(v...)
}

// Error message using log.PrintLn
func (l *AurLog) Errorln(v ...interface{}) {
	errorLvl.Println(v...)
}

// Error message using log.Printf
func (l *AurLog) Errorf(format string, v ...interface{}) {
	errorLvl.Printf(format, v...)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func (l *AurLog) Fatal(v ...interface{}) {
	fatalLvl.Fatal(v...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func (l *AurLog) Fatalf(format string, v ...interface{}) {
	fatalLvl.Fatalf(format, v...)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func (l *AurLog) Fatalln(v ...interface{}) {
	fatalLvl.Fatalln(v...)
}
