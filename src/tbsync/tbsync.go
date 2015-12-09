package tbsync

import (
	"log"
	"os/exec"
	//"printlog"
	"bufio"
	"os"
)

func Tbsync(logger *log.Logger) {
	path := "/tmp/1.txt"
	argv2 := []string{"-avz", path, "/tmp/adfas/3.txt"}
	c2 := exec.Command("rsync", argv2...)
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

	//printlog.Cmdprintlog(c2, logger)

	logger.Println("123")
}
