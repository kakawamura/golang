package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	target := "./"
	// 引数があれば上書き
	if len(os.Args) == 2 {
		target = os.Args[1]
	}

	files, _ := ioutil.ReadDir(target)

	for _, f := range files {
		if f.IsDir() {
			fmt.Printf("%s/\n", f.Name())
		} else {
			fmt.Println(f.Name())
		}
	}
}
