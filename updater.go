package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	update     = ""
	asar       = ""
	executable = ""
)

func main() {
	update = os.Args[1]
	asar = os.Args[2]
	executable = os.Args[3]
	time.Sleep(time.Duration(5) * time.Second)
	renameFile(update, asar)
	openExe(executable)
}

func renameFile(update string, asar string) {
	info, fileErr := os.Stat(update)
	if os.IsNotExist(fileErr) {
		return
	}
	fmt.Println(info)
	err := os.Rename(update, asar)
	if err != nil {
		fmt.Println(err)
	}
}

func openExe(executable string) {
	cmd := exec.Command("cmd.exe", "/C", "start", "/b", executable)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
	return
}
