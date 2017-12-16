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
	fp := "src/hello"
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
		fmt.Print("Not Load Dir")
	}
	for _, file := range files {
		fmt.Println("ReadDir:",file.Name())
	}
	//	createFile部分
	file, err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	fmt.Fprintln(file, "Hello")

	//	readFile部分
	content, err := ioutil.ReadFile("src/hello")
	if err != nil {
		log.Print(err)
	}
	fmt.Print("File contents: ", string(content), "\n")
}
func createFile(fp string)  {
	file, err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Fprintln(file, "Hello")
}

func touchFile()  {
	file, err := exec.Command("touch", "src/hello").CombinedOutput()
	if err != nil {
		log.Print("Error...Touch\n",string(file),"\n")
	}
}

func readDir(dn string)  {
	files, err := ioutil.ReadDir(dn)
	if err != nil {
		fmt.Print("Not Load Dir")
	}
	for _, file := range files {
		fmt.Println("ReadDir:",file.Name())
	}
}

func makeDir(dn string)  {
	mkdir, err := exec.Command("mkdir", dn).CombinedOutput()
	if err != nil {
		log.Print("Error...Mkdir\n",string(mkdir),"\n")
	}
}
func readFile()  {

	content, err := ioutil.ReadFile("src/hello")
	if err != nil {
		log.Print(err)
	}

	fmt.Print("File contents: ", string(content), "\n")

}
