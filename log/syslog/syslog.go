package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.Dial("", "localhost:1234", syslog.LOG_WARNING|syslog.LOG_DAEMON, "demotag")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(sysLog, "this is a daemon warning with demotag.")
	sysLog.Emerg("And this is a daemon emergency with demotag.")
}
