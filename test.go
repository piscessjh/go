package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("calc") //调用cmd程序，输入命令calc
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
