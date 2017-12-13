package main

import (
	"os"
	"os/exec"
	"log"
	"fmt"
	"time"
)
const df = "2006/01/02 15:04"
func main(){
	tm := time.Now()
	dir, err := os.Getwd()
	if err == nil {
		fmt.Print(dir, "\n")
	}
	fmt.Print("DateTime:", tm.Format(df), "\n")
	boot, err := exec.Command("clear").CombinedOutput()
	if err != nil {
		log.Print("Error NotClearScreen\n",string(boot))
	}
}
