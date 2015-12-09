package printlog

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

var multi_logfile []io.Writer
var logfile *os.File
var logger *log.Logger
var err1 error

/*func Initlog() （*os.File） {
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
	return logpath
}
*/

func Cmdprintlog(c2 *exec.Cmd, logger *log.Logger) {
	stdout, err := c2.StdoutPipe()
	if err != nil {
		logger.Println(err.Error())
		os.Exit(1)
	}
	if err := c2.Start(); err != nil {
		logger.Println("Command  error:", err.Error())
		os.Exit(1)
	}
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		logger.Println(in.Text())
	}
	if err := in.Err(); err != nil {
		logger.Println("Err:", err.Error())
		os.Exit(1)
	}

}

//func Printlog(info string) {
//	logger.Println("123")

//}

func main() {
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
}
