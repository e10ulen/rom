package main

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	"log"
)

func main() {
	dir, err := os.Getwd()
	if err == nil {
		fmt.Print(dir, "\n")
	}
	dn := "src"
	fp := dn+"/hello"
	loadFile(dn,fp)
	// readDir(dn)
	// makeDir(dn)
	// readDir(dn)
	// readFile()
  //
	// touchFile()
	// createFile(fp)
	// readFile()
}
// htmlソースを読み込んで、htmlの中に規定の記述があればロードしてある別のファイルを挿入する
func loadFile(dn string,fp string)  {
	//	readDir部分
	files, err := ioutil.ReadDir(dn)
	if err != nil {
		//	dirがなければmakeDirに飛んでdirを作り、帰ってきて続行する。
		makeDir(dn)
	}
	for _, file := range files {
		fmt.Println("ReadDir:",file.Name())
	}
	//	createFile部分
	cfile, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Print(err)
	}
	defer cfile.Close()

	fmt.Fprintln(cfile, "Hello")

	//	readFile部分
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Print(err)
	}
	fmt.Print("File contents: ", string(content), "\n")
}
func makeDir(dn string)  {
	mkdir, err := exec.Command("mkdir", dn).CombinedOutput()
	if err != nil {
		log.Print("Error...Mkdir\n",string(mkdir),"\n")
	}
}
