package main

import (
	//"bufio"
	//"flag"
	//"io"
	//"io/ioutil"
	//"log"
	//"os"
	//"os/exec"
	//"regexp"
	//"strconv"
	//"strings"
	//"time"
	//"printlog"

	"io"
	"log"
	"os"
	//"os/exec"
	"tbsync"
	"time"
)

var multi_logfile []io.Writer
var logfile *os.File
var logger *log.Logger
var err1 error

func main() {
	//printlog.Initlog()

	timestamp := time.Now().Format("20060102150405")
	logpath := "./log.log." + timestamp
	logfile, err1 = os.OpenFile(logpath, os.O_RDWR|os.O_CREATE, 0666)
	defer logfile.Close()
	if err1 != nil {
		logger.Println(err1.Error())
		os.Exit(1)
	}
	multi_logfile = []io.Writer{
		logfile,
		os.Stdout,
	}
	logfiles := io.MultiWriter(multi_logfile...)
	logger = log.New(logfiles, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("=====job start.=====")
	tbsync.Tbsync(logger)

}
