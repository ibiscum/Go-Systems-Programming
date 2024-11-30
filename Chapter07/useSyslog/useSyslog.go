package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {

	programName := filepath.Base(os.Args[0])
	sysLog, e := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if e != nil {
		log.Fatal(e)
	}
	err := sysLog.Crit("Crit: Logging in Go!")
	if err != nil {
		log.Fatal(err)
	}

	sysLog, e = syslog.New(syslog.LOG_ALERT|syslog.LOG_LOCAL7, "Some program!")
	if e != nil {
		log.Fatal(sysLog)
	}
	err = sysLog.Emerg("Emerg: Logging in Go!")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(sysLog, "log.Print: Logging in Go!")
}
