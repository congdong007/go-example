package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	pl := fmt.Println
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}

	pl("> date")
	pl(string(dateOut))

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()

	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	pl("> grep hello")
	pl(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	pl("> ls -a -l -h")
	pl(string(lsOut))
}
