// ゴルーチンを利用してgrepを早くする
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

func main() {
	grepString := os.Args[1]
	filePath := os.Args[2]
	regex := regexp.MustCompile(grepString)

	file, err := os.Open(filePath)

	if err != nil {
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if regex.Match([]byte(text)) {
			fmt.Println(text)
		}
	}
}
