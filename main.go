package main

import (
	"os"
	"fmt"
	"time"
	"bufio"
)
const df = "2006/01/02 15:04"
func main(){
	tm := time.Now()
	dir, err := os.Getwd()
	if err == nil {
		fmt.Print(dir, "\n")
	}
	fmt.Print("DateTime:", tm.Format(df), "\n")
//	Scan Stdin
	scr := bufio.NewScanner(os.Stdin)
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nScan Text > ")
	for scr.Scan(){
			fmt.Println(scr.Text())
			if scr.Text() == string("exit") {
				os.Exit(0)
			}
	}
	if err := scr.Err();err != nil {
		fmt.Fprintln(os.Stderr, "readinf standerd input", err)
	}
}
