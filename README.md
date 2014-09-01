# aurlog


Simple file-based wrapper for the standard Go log


###Installation
```
$ go get github.com/aurlaw/aurlog
```


## Getting Started
```go
package main

import "github.com/aurlaw/aurlog"

func main() {
	lc := aurlog.LogConfiguration{LogFile: "server.log"}

	alog := aurlog.Configure(&lc)

	alog.Debugln("Here is a debug log")
	alog.Infoln("Here is a info log")
	alog.Warningln("Here is a warning log")
	alog.Errorln("Here is a error log")

}
```

Executing this program will create a 'server.log' file with the following entries

```
DEBUG: 2014/08/31 00:47:22 [Here is a debug log]
INFO: 2014/08/31 00:47:22 [Here is a info log]
WARNING: 2014/08/31 00:47:22 [Here is a warning log]
ERROR: 2014/08/31 00:47:22 [Here is a error log]
```


### Configuration

The following log types are supported:

* Debug
* Info
* Warning
* Error
* Fatal

Debug, Info, Warning & Error support the Print, Printf & Println from the [Standard Go Log](http://golang.org/pkg/log/#Print)

i.e. Debug, Debugf, Debugln, etc.

Fatal support the Fatal, Fatalf & Fatalln from the [Standard Go Log](http://golang.org/pkg/log/#Fatal)


#### Log Levels
The following log levels can be enabled

* IsDebug - true or false. Defaults to true if no log level is configured.
* IsInfo - true or false. Defaults to true if no log level is configured.
* IsWarning - true or false. Defaults to true if no log level is configured.
* IsError - true or false. Defaults to true if no log level is configured.

```go
	lc := aurlog.LogConfiguration{}
	lc.IsInfo = true
	lc.IsError = true
```


#### Additional Configuration
* LogFile - defines the log file. All logfiles are prepended with the current date in yyyy-mm-dd format. If not supplied, uses the StdOut and StdErr defined by the [Standard Go Log](http://golang.org/pkg/log/). LogFile can contain an absolute or relative path
* NoStdOut - if set to true, disables logging to the StdOut or StdErr. False by default


```go
	lc := aurlog.LogConfiguration{LogFile: "test.log", NoStdOut: true}
	lc.IsInfo = true
	lc.IsError = true
```

Full Sample
```go
package main

import "github.com/aurlaw/aurlog"

func main() {
	lc := aurlog.LogConfiguration{LogFile: "server.log"}
	lc.IsInfo = true
	lc.IsError = true

	alog := aurlog.Configure(&lc)

	alog.Debugln("Here is a debug log")
	alog.Infoln("Here is a info log")
	alog.Warningln("Here is a warning log")
	alog.Errorln("Here is a error log")

}
```

Executing this program will create a 'server.log' file with the following entries

```
INFO: 2014/08/31 00:47:22 [Here is a info log]
ERROR: 2014/08/31 00:47:22 [Here is a error log]
```





